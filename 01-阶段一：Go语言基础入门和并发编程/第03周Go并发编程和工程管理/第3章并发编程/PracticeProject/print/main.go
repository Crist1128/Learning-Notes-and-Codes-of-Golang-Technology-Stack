package main

import (
	"fmt"
)

// 打印数字
func PrintNum(chNum chan struct{}, chLetter chan struct{}, done chan struct{}) {
	for i := 1; i <= 28; i += 2 {
		<-chNum                    // 等待数字信号
		fmt.Printf("%v%v", i, i+1) // 输出两个数字
		chLetter <- struct{}{}     // 发送字母信号
	}
	done <- struct{}{}
}

// 打印字母
func PrintLetter(chNum chan struct{}, chLetter chan struct{}) {
	for i := 0; i < 26; i += 2 {
		<-chLetter                         // 等待字母信号
		fmt.Printf("%c%c", 'A'+i, 'A'+i+1) // 输出两个字母
		chNum <- struct{}{}                // 发送数字信号
	}
}

func main() {

	chNum := make(chan struct{})       // 数字信号通道
	chLetter := make(chan struct{}, 1) // 字母信号通道
	done := make(chan struct{})

	// 启动两个 Goroutine，分别用于打印数字和字母
	go PrintNum(chNum, chLetter, done)
	go PrintLetter(chNum, chLetter)

	// 首先发送一个数字信号，开始打印数字
	chNum <- struct{}{}
	<-done

}
