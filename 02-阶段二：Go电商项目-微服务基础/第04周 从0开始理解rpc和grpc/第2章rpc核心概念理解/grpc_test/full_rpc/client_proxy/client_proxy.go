/**
 * @File : client_proxy.go
 * @Description : 封装客户端的调用逻辑，提供 Stub 代理
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-10-06
 */
package client_proxy

import (
	"grpc_test/full_rpc/handler" // 引入处理业务逻辑的包
	"log"                        // 导入日志包，便于记录日志
	"net/rpc"                    // 导入 RPC 包，处理远程过程调用
)

// HelloServiceStub 是客户端调用远程服务的代理
// 通过此结构体封装与服务端的连接和调用逻辑
type HelloServiceStub struct {
	*rpc.Client // 嵌入 rpc.Client，实现 RPC 调用功能
}

// NewHelloServiceStub 函数：用于创建一个新的服务代理实例
// 参数 protcol 是连接协议，如 "tcp"，addr 是服务端地址
func NewHelloServiceStub(protocol, addr string) HelloServiceStub {
	// 建立与服务端的连接
	conn, err := rpc.Dial(protocol, addr)
	if err != nil {
		// 日志记录连接错误，避免 panic
		log.Fatalf("连接失败: %v", err)
	}
	// 返回封装好的 HelloServiceStub 实例
	return HelloServiceStub{conn}
}

// Hello 方法：调用服务端的 Hello 方法，发送请求并接收响应
// 参数 request 是发送给服务端的请求数据，reply 是服务端返回的响应
func (s *HelloServiceStub) Hello(request string, reply *string) error {
	// 调用服务端的 Hello 方法
	err := s.Call(handler.HelloServiceName+".Hello", request, reply)
	if err != nil {
		// 记录错误日志，返回错误
		log.Printf("RPC 调用失败: %v", err)
		return err
	}
	return nil // 返回 nil 表示调用成功
}
