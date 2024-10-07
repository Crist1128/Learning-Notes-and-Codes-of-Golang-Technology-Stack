/**
 * @File : main.go
 * @Description : 请填写文件描述
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-09-26
 */
package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

func Worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker out...")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(time.Second)
		}
	}

}

func Processer(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Processer out...")
			return
		default:
			fmt.Println("Processing...")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Wait()
	ctx, cancel := context.WithCancel(context.Background())
	go Worker(ctx, &wg)
	go Processer(ctx, &wg)
	// 等待用户按下回车键
	fmt.Println("Press Enter to stop...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // 等待用户输入
	cancel()
	fmt.Println("Received cancel signal, exiting...")
}
