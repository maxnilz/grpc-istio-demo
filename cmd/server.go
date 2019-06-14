package main

import (
	"context"
	"log"
	"net"

	"github.com/maxnilz/grpc-istio-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	emoji "gopkg.in/kyokomi/emoji.v1"
)

type server struct{}

type AttributeCondition struct {
	key  string
	desc string
}

func (ac *AttributeCondition) String() {
	log.Printf("%v: %v", ac.key, ac.desc)
}

var ac1 = AttributeCondition{key: "c1", desc: "Can only perform the ops(update, delete, read) on the jobs he created"}
var ac2 = AttributeCondition{key: "c2", desc: "can only update the salary field of the jobs that he has been assigned to"}
var ac3 = AttributeCondition{key: "c3", desc: "can only update the requirements field of the jobs that he has been assigned to"}

//var APIAcMetas = []APIAcMeta{
//	{
//		api: "/api/say",
//		supportedACs: []AttributeCondition{
//			ac1,
//			ac2,
//			ac3,
//		},
//	},
//}

func IsAttributeConditionEnabled(ctx context.Context, ac AttributeCondition) bool {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false
	}

	res := md.Get(ac.key)
	if len(res) == 0 {
		return false
	}
	return true
}

func (s *server) InsertEmojis(ctx context.Context, req *proto.EmojiRequest) (*proto.EmojiResponse, error) {
	log.Printf("Client says: %s", req.InputText)
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Println(md)
	}

	if IsAttributeConditionEnabled(ctx, ac1) {
		log.Printf("Hi, we got ac in InsertEmojis: %v\n", ac1)
	}
	if IsAttributeConditionEnabled(ctx, ac2) {
		log.Printf("Hi, we got ac in InsertEmojis: %v\n", ac2)
	}
	if IsAttributeConditionEnabled(ctx, ac3) {
		log.Printf("Hi, we got ac in InsertEmojis: %v\n", ac3)
	}

	outputText := emoji.Sprint(req.InputText)
	log.Printf("Response: %s", outputText)
	return &proto.EmojiResponse{OutputText: outputText}, nil
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	log.Printf("Client: %s greeting!", req.Name)
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Println(md)
	}

	if IsAttributeConditionEnabled(ctx, ac1) {
		log.Printf("Hi, we got ac in SayHello: %v\n", ac1)
	}
	if IsAttributeConditionEnabled(ctx, ac2) {
		log.Printf("Hi, we got ac in SayHello: %v\n", ac2)
	}
	if IsAttributeConditionEnabled(ctx, ac3) {
		log.Printf("Hi, we got ac in SayHello: %v\n", ac3)
	}

	name := "service emojis"
	log.Printf("Service: %s greeting!", name)
	return &proto.HelloResponse{Name: name}, nil
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
