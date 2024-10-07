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
	"grpc_protoc/grpc_test/proto"
)

func main() {
	/*	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()*/

	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewHelloServiceClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "Alice"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
