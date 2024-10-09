/**
 * @File : server.go
 * @Description : 请填写文件描述
 * @Author : Junxi You
 * @Date : 2024-10-08
 */
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
	"protobuf_grpc_advance/protobuf_test/proto"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &proto.HelloReply{Message: "Hello " + in.GetName()}, nil
	}
	for k, v := range md {
		fmt.Println("Server received metadata: ", k, "=", v)
	}

	repMD := metadata.New(map[string]string{
		"Server-ID1": "111",
		"Server-ID2": "222",
		"Server-ID3": "333",
	})
	newctx := metadata.NewOutgoingContext(ctx, repMD)
	grpc.SetHeader(newctx, repMD)
	return &proto.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}
}
