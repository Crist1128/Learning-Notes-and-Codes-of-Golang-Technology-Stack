/**
 * @File : server.go
 * @Description : 启动 JSON-RPC 服务，处理客户端请求
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-10-06
 */
package main

import (
	"grpc_test/full_rpc/handler"      // 引入处理业务逻辑的包
	"grpc_test/full_rpc/server_proxy" // 引入服务代理注册的包
	"log"                             // 导入日志包，便于记录日志
	"net"                             // 导入网络包，用于监听 TCP 连接
	"net/rpc"                         // 导入 RPC 包，用于处理远程过程调用
)

func main() {
	// 创建监听器，监听 TCP 端口 1234
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		// 使用 log 记录错误，避免 panic
		log.Fatalf("监听失败: %v", err)
	}

	// 注册 Hello 服务
	err = server_proxy.RegisterHelloService(&handler.HelloServer{})
	if err != nil {
		// 错误处理，确保服务注册成功
		log.Fatalf("服务注册失败: %v", err)
	}

	// 循环处理客户端连接
	for {
		conn, err := listener.Accept() // 接收客户端连接
		if err != nil {
			// 继续处理其他连接
			log.Printf("连接接受失败: %v", err)
			continue
		}
		// 使用 goroutine 并发处理客户端请求，避免阻塞
		go rpc.ServeConn(conn)
	}
}
