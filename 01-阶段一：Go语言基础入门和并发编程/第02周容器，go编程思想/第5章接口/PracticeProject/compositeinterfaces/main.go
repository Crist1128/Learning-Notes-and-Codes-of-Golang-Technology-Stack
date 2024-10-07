/**
 * @File : main.go
 * @Description : 请填写文件描述
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-09-19
 */
package main

import "fmt"

type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

type PS interface {
	Printer
	Scanner
}

type MultiFunctionDevice struct {
}

func (mfd *MultiFunctionDevice) Print() {
	fmt.Println("Device is printing...")
}

func (mfd *MultiFunctionDevice) Scan() {
	fmt.Println("Device is scanning...")
}

func OperateDevice(ps PS) {
	ps.Print()
	ps.Scan()
}

func main() {
	OperateDevice(&MultiFunctionDevice{})
}
