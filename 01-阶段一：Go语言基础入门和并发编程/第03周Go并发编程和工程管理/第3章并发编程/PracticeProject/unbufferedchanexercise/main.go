/**
 * @File : main.go
 * @Description : 请填写文件描述
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-09-23
 */
package main

import (
	"fmt"
	"sync"
)

func Sender(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	ch <- "Hello from Goroutine 1!"
}

func Receiver(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	fmt.Println("Received message: ", <-ch)
}
func main() {
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(2)
	go Sender(&wg, ch)
	go Receiver(&wg, ch)
	wg.Wait()
}
