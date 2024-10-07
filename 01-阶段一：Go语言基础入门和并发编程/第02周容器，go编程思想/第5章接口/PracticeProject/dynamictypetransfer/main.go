package main

import (
	"fmt"
	"strings"
)

// Processor 接口，包含 Process 方法
type Processor interface {
	Process(data interface{})
}

// IntProcessor 结构体，实现 Processor 接口
type IntProcessor struct{}

// 实现 Process 方法，用于处理整数类型
func (ip IntProcessor) Process(data interface{}) {
	// 使用类型断言确保传入的数据是整数
	if v, ok := data.(int); ok {
		fmt.Printf("IntProcessor received %d, processing: %d\n", v, v*v)
	} else {
		fmt.Println("IntProcessor can only process integers!")
	}
}

// StringProcessor 结构体，实现 Processor 接口
type StringProcessor struct{}

// 实现 Process 方法，用于处理字符串类型
func (sp StringProcessor) Process(data interface{}) {
	// 使用类型断言确保传入的数据是字符串
	if v, ok := data.(string); ok {
		fmt.Printf("StringProcessor received \"%s\", processing: \"%s\"\n", v, strings.ToUpper(v))
	} else {
		fmt.Println("StringProcessor can only process strings!")
	}
}

// HandleProcess 函数，根据传入的 Processor 类型处理不同类型的数据
func HandleProcess(p Processor, data interface{}) {
	p.Process(data)
}

func main() {
	// 创建处理器实例
	intProcessor := IntProcessor{}
	stringProcessor := StringProcessor{}

	// 传入不同类型的数据给处理器
	HandleProcess(intProcessor, 4)
	HandleProcess(stringProcessor, "hello")
}
