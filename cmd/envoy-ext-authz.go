// Envoy’s External Authorization Service.
// In order to provide session management, token acquisition and RCToken functionality
// as well as being extensible in the future to meet other SSO technologies
// TODO: The Service should be a pipe-and-filter component.
// Filters will be able to fail or progressively augment the response returned from the TS to the Envoy gateway.
package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	oidc "github.com/coreos/go-oidc"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	"github.com/gogo/googleapis/google/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v2"
)

type server struct{}

type AC struct {
	Name  string `json:"name"`
	Metas string `json:"metas"`
	Desc  string `json:"desc"`
}

type ACMeta struct {
	Resource     string `json:"resource"`
	Action       string `json:"action"`
	SupportedACs []AC   `json:"supported-acs" yaml:"supported-acs"`
	EnabledACs   []AC   `json:"enabled-acs" yaml:"enabled-acs"`
}

type APIACMeta struct {
	Service string   `json:"service"`
	ACMetas []ACMeta `json:"ac-metas" yaml:"ac-metas"`
}

type ServicesAPIACMeta struct {
	Services []APIACMeta `json:"services"`
}

var serviceACMetas map[string][]AC

// Hard code here for demo.
// TODO: should call Role-based-access-control service
func checkRBAC(ctx context.Context, req *auth.CheckRequest) bool {
	_ = ctx
	_ = req

	return true
}

// parsePath extract service name, resource and action name
func parsePath(req *auth.CheckRequest) (service string, resource string, action string) {
	httpRequest := req.Attributes.Request.GetHttp()

	parts := strings.Split(httpRequest.GetPath(), "/")
	if len(parts) > 1 {
		service = parts[1]
	}
	resource = strings.TrimPrefix(httpRequest.GetPath(), "/"+service)
	action = httpRequest.GetMethod()
	return
}

// acIDFromReq combine the unique id for a specific request
func acIDFromReq(req *auth.CheckRequest) (id string) {
	service, resource, action := parsePath(req)
	id = strings.Join([]string{service, resource, action}, ":")
	return
}

// acIDFromConf ...
func acIDFromConf(service, resource, action string) (id string) {
	id = strings.Join([]string{service, resource, action}, ":")
	return
}

func (s *server) Check(ctx context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {
	// Extract the http request
	httpRequest := req.Attributes.Request.GetHttp()

	// Extract the original path
	path := httpRequest.GetPath()
	log.Printf("Getting ext-auth-check request %s at %s", path, time.Now().Format(time.RFC3339))

	byPassPaths := []string{
		"/dex",
	}
	for _, v := range byPassPaths {
		if strings.HasPrefix(path, v) {
			return &auth.CheckResponse{}, nil
		}
	}

	// Check if the authorization header if valid or not
	// Extract the id_token from authorization
	hdrs := httpRequest.GetHeaders()
	log.Println(hdrs)
	authorization := ""
	if v, ok := hdrs["authorization"]; ok {
		authorization = v
	}
	parts := strings.Fields(authorization)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
		deniedHttpResponse := &auth.DeniedHttpResponse{
			Status: &_type.HttpStatus{
				Code: _type.StatusCode_Unauthorized,
			},
		}
		return &auth.CheckResponse{
			Status: &rpc.Status{
				Code: int32(rpc.ABORTED),
			},
			HttpResponse: &auth.CheckResponse_DeniedResponse{
				DeniedResponse: deniedHttpResponse,
			},
		}, nil
	}
	token := parts[1]
	// Create oidc provider, should not be here, just for demo
	oidcProvider, err := oidc.NewProvider(ctx, "http://192.168.39.224:31284/dex")
	if err != nil {
		deniedHttpResponse := &auth.DeniedHttpResponse{
			Status: &_type.HttpStatus{
				Code: _type.StatusCode_InternalServerError,
			},
			Body: err.Error(),
		}
		return &auth.CheckResponse{
			Status: &rpc.Status{
				Code: int32(rpc.ABORTED),
			},
			HttpResponse: &auth.CheckResponse_DeniedResponse{
				DeniedResponse: deniedHttpResponse,
			},
		}, nil
	}
	oidcConfig := &oidc.Config{SkipClientIDCheck: true}
	verifier := oidcProvider.Verifier(oidcConfig)
	IDToken, err := verifier.Verify(ctx, token)
	if err != nil {
		deniedHttpResponse := &auth.DeniedHttpResponse{
			Status: &_type.HttpStatus{
				Code: _type.StatusCode_Unauthorized,
			},
			Body: err.Error(),
		}
		return &auth.CheckResponse{
			Status: &rpc.Status{
				Code: int32(rpc.UNAUTHENTICATED),
			},
			HttpResponse: &auth.CheckResponse_DeniedResponse{
				DeniedResponse: deniedHttpResponse,
			},
		}, nil
	}
	_ = IDToken

	// Verify RBAC
	if !checkRBAC(ctx, req) {
		deniedHttpResponse := &auth.DeniedHttpResponse{
			Status: &_type.HttpStatus{
				Code: _type.StatusCode_Forbidden,
			},
		}
		return &auth.CheckResponse{
			Status: &rpc.Status{
				Code: int32(rpc.PERMISSION_DENIED),
			},
			HttpResponse: &auth.CheckResponse_DeniedResponse{
				DeniedResponse: deniedHttpResponse,
			},
		}, nil
	}
	// Check if need ABAC-verify
	acID := acIDFromReq(req)
	log.Println(acID)
	okHttpResponse := &auth.OkHttpResponse{}
	if enabledACs, ok := serviceACMetas[acID]; ok {
		// Forward the enable attribute condition to downstream services for the ABAC verification
		log.Println("Setting request context ", acID)
		for _, enabledAC := range enabledACs {
			log.Printf("Setting request context for %v with %v ", acID, enabledAC)
			okHttpResponse.Headers = append(okHttpResponse.Headers, &core.HeaderValueOption{
				Header: &core.HeaderValue{
					Key: enabledAC.Name,
					// the configured metadata comes from control panel
					Value: enabledAC.Metas,
				},
			})
		}
	}

	return &auth.CheckResponse{
		HttpResponse: &auth.CheckResponse_OkResponse{
			OkResponse: okHttpResponse,
		},
	}, nil
}

func main() {
	// Parse config file
	wd, _ := os.Getwd()
	filename := filepath.Join(wd, "cmd/apis-ac-meta-enabled.yaml")
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	servicesAPIACMeta := ServicesAPIACMeta{}
	if err := yaml.Unmarshal(b, &servicesAPIACMeta); err != nil {
		log.Fatalf("Failed to get config details: %v", err)
	}
	if serviceACMetas == nil {
		serviceACMetas = make(map[string][]AC)
	}
	for _, servicesAPIACMeta := range servicesAPIACMeta.Services {
		for _, acMeta := range servicesAPIACMeta.ACMetas {
			id := acIDFromConf(servicesAPIACMeta.Service, acMeta.Resource, acMeta.Action)
			if _, ok := serviceACMetas[id]; !ok {
				serviceACMetas[id] = acMeta.EnabledACs
			}
		}
	}

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Listening on %s", lis.Addr())

	s := grpc.NewServer()
	auth.RegisterAuthorizationServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
