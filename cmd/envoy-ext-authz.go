// Envoy’s External Authorization Service.
// In order to provide session management, token acquisition and RCToken functionality
// as well as being extensible in the future to meet other SSO technologies
// the Service better be a pipe-and-filter component.
// Filters will be able to fail or progressively augment the response returned from the TS to the Envoy gateway.
package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	oidc "github.com/coreos/go-oidc"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	"github.com/gogo/googleapis/google/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

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
				Code: int32(rpc.PERMISSION_DENIED),
			},
			HttpResponse: &auth.CheckResponse_DeniedResponse{
				DeniedResponse: deniedHttpResponse,
			},
		}, nil
	}
	token := parts[1]
	// Create oidc provider
	oidcProvider, err := oidc.NewProvider(ctx, "http://192.168.39.224:31380/dex")
	if err != nil {
		deniedHttpResponse := &auth.DeniedHttpResponse{
			Status: &_type.HttpStatus{
				Code: _type.StatusCode_InternalServerError,
			},
			Body: err.Error(),
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
	oidcConfig := &oidc.Config{SkipClientIDCheck: true}
	verifier := oidcProvider.Verifier(oidcConfig)
	IDToken, err := verifier.Verify(ctx, token)
	if err != nil {
		deniedHttpResponse := &auth.DeniedHttpResponse{
			Status: &_type.HttpStatus{
				Code: _type.StatusCode_InternalServerError,
			},
			Body: err.Error(),
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
	_ = IDToken

	return &auth.CheckResponse{}, nil
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
