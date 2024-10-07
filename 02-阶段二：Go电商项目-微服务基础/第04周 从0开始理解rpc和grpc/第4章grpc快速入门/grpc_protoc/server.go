/**
 * @File : server.go
 * @Description : 请填写文件描述
 * @Author : Junxi You
 * @Date : 2024-10-07
 */
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc_protoc/grpc_protoc/hello" // 导入生成的 protobuf 包
	"net"
)

// 定义一个服务器结构体，实现 HelloServiceServer 接口
type HelloServer struct {
	pb.UnimplementedHelloServiceServer
}

// 实现 SayHello 方法
func (s *HelloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	message := fmt.Sprintf("Hello, %s!", req.Name)
	return &pb.HelloResponse{Message: message}, nil
}

func main() {
	// 启动 gRPC 服务器
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterHelloServiceServer(server, &HelloServer{})

	fmt.Println("gRPC server listening on port 50051...")
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
