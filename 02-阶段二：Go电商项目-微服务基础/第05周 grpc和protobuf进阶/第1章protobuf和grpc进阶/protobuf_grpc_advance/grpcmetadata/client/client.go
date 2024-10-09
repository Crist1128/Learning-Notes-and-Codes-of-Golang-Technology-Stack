/**
 * @File : client.go
 * @Description : 请填写文件描述
 * @Author : Junxi You
 * @Date : 2024-10-08
 */
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"protobuf_grpc_advance/protobuf_test/proto"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	md := metadata.New(map[string]string{
		"Client-ID1": "111",
		"Client-ID2": "222",
		"Client-ID3": "333",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	serMD := metadata.MD{}
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "gRPC!"}, grpc.Header(&serMD))
	for k, v := range md {
		fmt.Println("Client sending metadata: ", k, "=", v)
	}

	if err != nil {
		panic(err)
	}

	for k, v := range serMD {
		fmt.Println("Client received metadata: ", k, "=", v)
	}
	fmt.Println(r.Message)
}
