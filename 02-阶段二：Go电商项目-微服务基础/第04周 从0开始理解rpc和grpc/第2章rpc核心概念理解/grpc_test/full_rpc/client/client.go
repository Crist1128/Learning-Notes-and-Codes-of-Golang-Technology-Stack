/**
 * @File : client.go
 * @Description : 客户端主程序，调用远程服务
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-10-06
 */
package main

import (
	"fmt"                             // 导入 fmt 包，用于输出
	"grpc_test/full_rpc/client_proxy" // 引入客户端代理包
)

func main() {
	// 创建与服务端的连接代理，使用 TCP 协议连接 127.0.0.1:1234 地址
	conn := client_proxy.NewHelloServiceStub("tcp", "127.0.0.1:1234")

	var reply string // 用于接收服务端返回的数据
	// 调用远程 Hello 方法，传入请求 "cc"
	err := conn.Hello("cc", &reply)
	if err != nil {
		// 记录调用失败的错误日志
		fmt.Printf("远程调用失败: %v\n", err)
		return
	}

	// 输出服务端返回的响应结果
	fmt.Println("服务端响应:", reply)
}
