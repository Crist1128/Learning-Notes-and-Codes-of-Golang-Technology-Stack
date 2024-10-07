/**
 * @File : main.go
 * @Description : 请填写文件描述
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-09-25
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	message := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Second)
			message <- i
		}
	}()
	for {
		select {
		case mes, ok := <-message:
			if !ok {
				fmt.Println("ALL message is received!")
				return
			}
			fmt.Println("Received message: Message ", mes)
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(3 * time.Second)
		case <-timer.C:
			fmt.Println("Timeout!")
			return
		}
	}

}
