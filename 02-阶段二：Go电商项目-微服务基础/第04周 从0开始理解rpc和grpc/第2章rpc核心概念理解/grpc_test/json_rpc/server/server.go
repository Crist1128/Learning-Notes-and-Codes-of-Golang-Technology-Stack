/**
 * @File : server.go
 * @Description : JSON-RPC Server Example
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-10-06
 */
package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloServer struct{}

// 服务端方法，接收 request 并返回处理后的 reply
func (s *HelloServer) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

func main() {
	// 创建监听器，监听 TCP 连接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	// 注册服务，服务名称为 "HelloService"
	err = rpc.RegisterName("HelloService", &HelloServer{})
	if err != nil {
		panic(err)
	}

	// 接收并处理连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// 使用 JSON-RPC 代替默认的 RPC 序列化协议
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
