/**
 * @File : main.go
 * @Description : 请填写文件描述
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-09-24
 */
package main

import (
	"fmt"
	"sync"
)

func Sender(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func Receiver(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	for data := range ch {
		fmt.Println("Received message: Message ", data)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 5)
	wg.Add(2)
	defer wg.Wait()
	go Sender(&wg, ch)
	go Receiver(&wg, ch)
}
