// Envoy’s External Authorization Service.
// In order to provide session management, token acquisition and RCToken functionality
// as well as being extensible in the future to meet other SSO technologies
// TODO: The Service should be a pipe-and-filter component.
// Filters will be able to fail or progressively augment the response returned from the TS to the Envoy gateway.
package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	oidc "github.com/coreos/go-oidc"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	v2 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2alpha"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	"github.com/gogo/googleapis/google/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

type AC struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type ACMeta struct {
	Resource     string `json:"resource"`
	Action       string `json:"action"`
	SupportedACs []AC   `json:"supported-acs"`
	EnabledACs   []AC   `json:"enabled-acs"`
}

type APIACmeta struct {
	Service string   `json:"service"`
	ACMetas []ACMeta `json:"ac-metas"`
}

// Hard code here for demo.
// TODO: Later this will comes from comes from control panel
var APIAcMetas = map[string][]string{
	"/api/say":                         {"c1", "c2", "c3"},
	"/api/emoji":                       {"c1", "c2"},
	"/proto.EmojiService/InsertEmojis": {"c1", "c2"},
}

// Hard code here for demo.
// TODO: should call Role-based-access-control service
func checkRBAC(ctx context.Context, req *v2.CheckRequest) bool {
	_ = ctx
	_ = req

	return true
}

func (s *server) Check(ctx context.Context, req *v2.CheckRequest) (*v2.CheckResponse, error) {
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
			return &v2.CheckResponse{}, nil
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
		deniedHttpResponse := &v2.DeniedHttpResponse{
			Status: &_type.HttpStatus{
				Code: _type.StatusCode_Unauthorized,
			},
		}
		return &v2.CheckResponse{
			Status: &rpc.Status{
				Code: int32(rpc.ABORTED),
			},
			HttpResponse: &v2.CheckResponse_DeniedResponse{
				DeniedResponse: deniedHttpResponse,
			},
		}, nil
	}
	token := parts[1]
	// Create oidc provider, should not be here, just for demo
	oidcProvider, err := oidc.NewProvider(ctx, "http://192.168.39.224:31380/dex")
	if err != nil {
		deniedHttpResponse := &v2.DeniedHttpResponse{
			Status: &_type.HttpStatus{
				Code: _type.StatusCode_InternalServerError,
			},
			Body: err.Error(),
		}
		return &v2.CheckResponse{
			Status: &rpc.Status{
				Code: int32(rpc.ABORTED),
			},
			HttpResponse: &v2.CheckResponse_DeniedResponse{
				DeniedResponse: deniedHttpResponse,
			},
		}, nil
	}
	oidcConfig := &oidc.Config{SkipClientIDCheck: true}
	verifier := oidcProvider.Verifier(oidcConfig)
	IDToken, err := verifier.Verify(ctx, token)
	if err != nil {
		deniedHttpResponse := &v2.DeniedHttpResponse{
			Status: &_type.HttpStatus{
				Code: _type.StatusCode_Unauthorized,
			},
			Body: err.Error(),
		}
		return &v2.CheckResponse{
			Status: &rpc.Status{
				Code: int32(rpc.UNAUTHENTICATED),
			},
			HttpResponse: &v2.CheckResponse_DeniedResponse{
				DeniedResponse: deniedHttpResponse,
			},
		}, nil
	}
	_ = IDToken

	okHttpResponse := &v2.OkHttpResponse{}
	// Check if need ac-verify
	if enabledACs, ok := APIAcMetas[path]; ok {
		// Verify RBAC
		if !checkRBAC(ctx, req) {
			deniedHttpResponse := &v2.DeniedHttpResponse{
				Status: &_type.HttpStatus{
					Code: _type.StatusCode_Forbidden,
				},
			}
			return &v2.CheckResponse{
				Status: &rpc.Status{
					Code: int32(rpc.PERMISSION_DENIED),
				},
				HttpResponse: &v2.CheckResponse_DeniedResponse{
					DeniedResponse: deniedHttpResponse,
				},
			}, nil
		}
		// Forward the enable attribute condition to downstream services for the ABAC verification
		for _, enabledAC := range enabledACs {
			okHttpResponse.Headers = append(okHttpResponse.Headers, &core.HeaderValueOption{
				Header: &core.HeaderValue{
					Key: enabledAC,
					// TODO: Later we should and addition info here,
					// e.g: the configured metadata comes from control panel
					Value: "",
				},
			})
		}
	}

	return &v2.CheckResponse{
		HttpResponse: &v2.CheckResponse_OkResponse{
			OkResponse: okHttpResponse,
		},
	}, nil
}

func main() {
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
