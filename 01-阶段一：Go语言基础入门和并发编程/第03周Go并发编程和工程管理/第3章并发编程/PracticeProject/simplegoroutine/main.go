package main

import (
	"fmt"
	"sync"
	"time"
)

func count(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 在 Goroutine 完成时调用 Done
	for i := 1; i <= 5; i++ {
		fmt.Printf("Task %v: %v\n", id, i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup

	// 添加 3 个 Goroutine 到 WaitGroup
	wg.Add(3)

	go count(1, &wg)
	go count(2, &wg)
	go count(3, &wg)

	// 等待所有 Goroutine 完成
	wg.Wait()

	fmt.Println("All tasks completed.")
}
