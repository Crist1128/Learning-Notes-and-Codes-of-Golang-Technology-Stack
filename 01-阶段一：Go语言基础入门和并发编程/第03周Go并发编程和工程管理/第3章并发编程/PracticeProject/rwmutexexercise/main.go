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
	"time"
)

type Data struct {
	value   int64
	rwMutex sync.RWMutex
}

func ReadData(wg *sync.WaitGroup, data *Data) {
	defer wg.Done()
	data.rwMutex.RLock()
	defer data.rwMutex.RUnlock()
	fmt.Println("Reading data: ", data.value)
	time.Sleep(200 * time.Millisecond)
}

func WriteData(wg *sync.WaitGroup, data *Data) {
	defer wg.Done()
	data.rwMutex.Lock()
	defer data.rwMutex.Unlock()
	atomic.AddInt64(&data.value, 1)
	fmt.Println("Writing data: ", data.value)
	time.Sleep(500 * time.Millisecond)
}

func main() {
	var wg sync.WaitGroup
	data := &Data{value: 0}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go ReadData(&wg, data)
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go WriteData(&wg, data)
	}
	wg.Wait()
	fmt.Println("Final data value：", data.value)
}
