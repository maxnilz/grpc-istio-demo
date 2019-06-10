// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"context"
	"log"
	"net"

	"github.com/maxnilz/grpc-istio-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	emoji "gopkg.in/kyokomi/emoji.v1"
)

type server struct{}

func (s *server) InsertEmojis(ctx context.Context, req *proto.EmojiRequest) (*proto.EmojiResponse, error) {
	log.Printf("Client says: %s", req.InputText)
	outputText := emoji.Sprint(req.InputText)
	log.Printf("Response: %s", outputText)
	return &proto.EmojiResponse{OutputText: outputText}, nil
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error)  {
	log.Printf("Client: %s greeting!", req.Name)
	name := "service emojis"
	log.Printf("Service: %s greeting!", name)
	return &proto.HelloResponse{Name:name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Listening on %s", lis.Addr())

	s := grpc.NewServer()
	proto.RegisterEmojiServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
