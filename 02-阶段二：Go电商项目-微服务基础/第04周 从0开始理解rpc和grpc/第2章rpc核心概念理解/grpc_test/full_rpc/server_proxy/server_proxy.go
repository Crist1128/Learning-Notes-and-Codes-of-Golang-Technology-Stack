/**
 * @File : server_proxy.go
 * @Description : 负责服务端服务的注册逻辑和接口定义
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-10-06
 */
package server_proxy

import (
	"grpc_test/full_rpc/handler" // 引入业务逻辑所在的包
	"net/rpc"                    // 导入 RPC 包，处理远程过程调用
)

// HelloServicer 接口：定义了服务的行为规范
// 任何实现此接口的服务都需要实现 Hello 方法
type HelloServicer interface {
	Hello(request string, reply *string) error
}

// RegisterHelloService 函数：将服务实例注册到 RPC 系统
// 参数 srv 是实现了 HelloServicer 接口的服务
// 返回值是注册服务的错误（如果存在）
func RegisterHelloService(srv HelloServicer) error {
	// 注册服务名称，handler.HelloServiceName 定义了服务名称
	return rpc.RegisterName(handler.HelloServiceName, srv)
}
