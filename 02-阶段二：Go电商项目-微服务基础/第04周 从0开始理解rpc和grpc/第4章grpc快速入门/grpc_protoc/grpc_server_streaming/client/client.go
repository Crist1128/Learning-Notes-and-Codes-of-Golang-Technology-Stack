/**
 * @File : client.go
 * @Description : 请填写文件描述
 * @Author : Junxi You
 * @Date : 2024-10-07
 */
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_protoc/grpc_server_streaming/proto"
	"io"
)

func main() {
	// 连接到 gRPC 服务器，使用不安全凭证（没有 TLS 加密）
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 创建 Greeter 客户端
	c := proto.NewGreeterClient(conn)

	// 调用 StreamNumbers 以开始服务器流
	r, err := c.StreamNumbers(context.Background(), &proto.StreamRequest{})
	if err != nil {
		panic(err)
	}
	// 接收来自服务器的消息流，直到流结束
	for {
		a, err := r.Recv()
		if err == io.EOF {
			fmt.Println("All numbers received.")
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("Received number: " + a.Data)
	}

}
