/**
 * @File : server.go
 * @Description : HTTP-RPC Server Example
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-10-06
 */
package main

import (
	"io"
	"net/http"
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

	// 注册服务，服务名称为 "HelloService"
	err := rpc.RegisterName("HelloService", &HelloServer{})
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: request.Body,
			Writer:     writer,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":1234", nil)
}
