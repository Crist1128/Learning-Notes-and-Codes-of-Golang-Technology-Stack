/**
 * @File : server.go
 * @Description : 请填写文件描述
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-10-06
 */
package main

import (
	"net"
	"net/rpc"
)

type HelloServer struct{}

func (s *HelloServer) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}
func main() {
	listener, _ := net.Listen("tcp", ":1234")
	_ = rpc.RegisterName("HelloService", &HelloServer{})
	conn, _ := listener.Accept()
	rpc.ServeConn(conn)
}
