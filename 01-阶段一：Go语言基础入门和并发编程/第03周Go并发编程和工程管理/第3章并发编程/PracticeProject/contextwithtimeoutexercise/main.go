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

func Worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Worker task timeout, exiting...")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(time.Second)
		}
	}
}

func Processer(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 4; i++ {
		fmt.Println("Processing...")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	wg.Add(2)
	go Worker(ctx, &wg)
	go Processer(&wg)
	wg.Wait()
	fmt.Println("All tasks finished.")
}
