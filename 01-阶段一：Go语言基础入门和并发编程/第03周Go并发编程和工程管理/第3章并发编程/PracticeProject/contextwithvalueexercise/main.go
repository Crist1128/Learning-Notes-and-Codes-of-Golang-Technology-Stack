/**
 * @File : main.go
 * @Description : 请填写文件描述
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-09-26
 */
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const UserID = "userID"
const RequestID = "requestID"

func Caller(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Caller out...")
			return
		default:
			fmt.Println("User ID: ", ctx.Value(UserID), "is working")
			time.Sleep(time.Second)
		}
	}
}

func Responder(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Responder out...")
			return
		default:
			fmt.Println("Request ID: ", ctx.Value(RequestID), "is working")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, UserID, "123")
	ctx = context.WithValue(ctx, RequestID, "abc123")
	var wg sync.WaitGroup
	wg.Add(2)
	go Caller(ctx, &wg)
	go Responder(ctx, &wg)
	wg.Wait()
	fmt.Println("Main Goroutine done")
}
