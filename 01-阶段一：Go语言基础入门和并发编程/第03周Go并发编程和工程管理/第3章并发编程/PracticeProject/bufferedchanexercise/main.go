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
	ch <- "Hello"
	ch <- "World"
	ch <- "from Goroutine"
	close(ch)
}

func Receiver(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	for data := range ch {
		fmt.Println("Received message: ", data)
	}
	fmt.Println("Received all message!")
}
func main() {
	var wg sync.WaitGroup
	ch := make(chan string, 3)
	wg.Add(2)
	go Receiver(&wg, ch)
	go Sender(&wg, ch)
	wg.Wait()
}
