package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mutex  sync.Mutex
	cond   = sync.NewCond(&mutex)
	buffer int
)

func Sender(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for i := 1; i <= 7; i++ {
		ch <- i
		mutex.Lock()
		fmt.Printf("Sent: %d (buffer size after sending: %d)\n", i, len(ch))
		if len(ch) == cap(ch) {
			fmt.Println("Channel is full, waiting for a receiver...")
		}
		cond.Signal() // 通知接收方
		mutex.Unlock()
		time.Sleep(200 * time.Millisecond)
	}
	close(ch)
}

func Receiver(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		mutex.Lock()
		fmt.Printf("Receiving: %d (buffer size after receiving: %d)\n", data, len(ch))
		cond.Signal() // 通知发送方
		mutex.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 3)

	wg.Add(2)

	go Sender(&wg, ch)
	go Receiver(&wg, ch)

	wg.Wait()
}
