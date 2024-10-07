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
	"sync/atomic"
)

var count int64

func Counter(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&count, 1)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Counter(&wg)
	}
	wg.Wait()
	fmt.Println("Final count value: ", count)
}
