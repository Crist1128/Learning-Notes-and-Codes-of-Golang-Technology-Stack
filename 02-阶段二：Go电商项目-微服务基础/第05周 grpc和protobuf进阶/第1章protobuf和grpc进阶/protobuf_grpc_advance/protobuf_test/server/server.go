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
	"net"
	"protobuf_grpc_advance/protobuf_test/proto"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	fmt.Println(in.RequestTime.AsTime())
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
