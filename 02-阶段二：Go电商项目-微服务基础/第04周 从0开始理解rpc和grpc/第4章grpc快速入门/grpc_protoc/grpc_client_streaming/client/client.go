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
	"grpc_protoc/grpc_client_streaming/proto"
	"time"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewSumServiceClient(conn)

	r, err := c.StreamSum(context.Background())
	if err != nil {
		panic(err)
	}
	for i := 1; i <= 10; i++ {
		err := r.Send(&proto.SumRequest{Number: int32(i)})
		if err != nil {
			panic(err)
		}
		fmt.Println("Sent number: " + fmt.Sprintf("%d", i))
		time.Sleep(time.Second)
	}
	a, err := r.CloseAndRecv()
	if err != nil {
		panic(err)
	}
	fmt.Println("Received sum: " + fmt.Sprint(a.Sum))
}
