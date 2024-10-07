package main

import "fmt"

// Sayer 接口定义
type Sayer interface {
	Say() string
}

// Dog 结构体定义
type Dog struct{}

// Say 方法实现 Dog 的叫声
func (d Dog) Say() string {
	return "Woof"
}

// Cat 结构体定义
type Cat struct{}

// Say 方法实现 Cat 的叫声
func (c Cat) Say() string {
	return "Meow"
}

// Duck 结构体定义
type Duck struct{}

// Say 方法实现 Duck 的叫声
func (d Duck) Say() string {
	return "Quack"
}

// MakeSound 接收一个 Sayer 接口类型并打印叫声
func MakeSound(name string, s Sayer) {
	fmt.Printf("%s says: %s\n", name, s.Say())
}

func main() {
	// 使用值类型调用
	MakeSound("Dog", Dog{})
	MakeSound("Cat", Cat{})
	MakeSound("Duck", Duck{})
}
