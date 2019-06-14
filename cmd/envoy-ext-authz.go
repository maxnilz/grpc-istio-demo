package main

import (
	"context"
	"log"
	"net"
	"time"

	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Check(ctx context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {
	log.Printf("Getting ext-auth-check request: %s", time.Now().Format(time.RFC3339))
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
