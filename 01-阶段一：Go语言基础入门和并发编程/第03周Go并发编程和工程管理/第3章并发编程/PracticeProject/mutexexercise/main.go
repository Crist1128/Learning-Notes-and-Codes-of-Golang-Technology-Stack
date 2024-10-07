/**
 * @File : main.go
 * @Description : 请填写文件描述
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-09-21
 */
package main

import (
	"fmt"
	"sync"
)

var count int

func Counter(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		count++
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Counter(&wg, &mu)
	}
	wg.Wait()
	fmt.Println("Final counter value: ", count)
}
