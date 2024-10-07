/**
 * @File : server.go
 * @Description : 请填写文件描述
 * @Author : Junxi You
 * @Date : 2024-10-07
 */
package main

import (
	"context"
	"google.golang.org/grpc"
	"net"

	"grpc_protoc/grpc_test/proto"
)

type Server struct {
	proto.UnimplementedHelloServiceServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{Message: "hello" + request.Name}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterHelloServiceServer(g, &Server{})

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	err = g.Serve(listener)
	if err != nil {
		panic(err)
	}

}
