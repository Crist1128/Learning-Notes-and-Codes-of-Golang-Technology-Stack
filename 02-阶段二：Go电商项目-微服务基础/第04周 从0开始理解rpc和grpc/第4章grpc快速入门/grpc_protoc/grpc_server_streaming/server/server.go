/**
 * @File : server.go
 * @Description : 请填写文件描述
 * @Author : Junxi You
 * @Date : 2024-10-07
 */
package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_protoc/grpc_server_streaming/proto"
	"net"
	"strconv"
	"time"
)

// server 结构体实现了 proto 定义的 Greeter 服务
type server struct {
	proto.UnimplementedGreeterServer
}

// StreamNumbers 是服务器流式传输的核心逻辑
// 它会向客户端发送 1 到 10 的连续数字，逐次发送后等待 1 秒
func (s *server) StreamNumbers(req *proto.StreamRequest, res proto.Greeter_StreamNumbersServer) error {
	i := 1
	for {
		// 如果 i 大于 10，结束流式传输
		if i > 10 {
			break
		}
		// 打印服务器正在发送的数字
		fmt.Println("Sending number: " + strconv.Itoa(i))

		// 通过 res.Send 方法向客户端发送 StreamResponse 消息
		err := res.Send(&proto.StreamResponse{
			Data: fmt.Sprintf("%d", i), // 将数字 i 转换为字符串并放入消息中
		})
		if err != nil {
			return err // 传输错误处理
		}

		i++
		time.Sleep(time.Second) // 每次发送后等待 1 秒
	}
	return nil
}

func main() {
	// 监听 50051 端口，准备接受 gRPC 请求
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	// 创建 gRPC 服务器
	s := grpc.NewServer()

	// 注册 Greeter 服务到服务器
	proto.RegisterGreeterServer(s, &server{})

	// 启动服务器，监听传入的 gRPC 请求
	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}
}
