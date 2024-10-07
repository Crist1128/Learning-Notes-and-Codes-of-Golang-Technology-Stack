/**
 * @File : handler.go
 * @Description : 定义业务处理逻辑及服务名称常量
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-10-06
 */
package handler

// HelloServiceName 是服务名称常量
// 客户端在调用远程方法时会使用此名称
const HelloServiceName = "handler/HelloService"

// HelloServer 结构体实现了 HelloServicer 接口，提供具体业务逻辑
type HelloServer struct{}

// Hello 方法：实现业务逻辑，接收客户端的请求并返回响应
// request 是客户端发送的请求数据，reply 是返回给客户端的响应数据
func (s *HelloServer) Hello(request string, reply *string) error {
	*reply = "hello, " + request // 简单的业务逻辑，将 "hello, " 和客户端请求的字符串拼接在一起
	return nil                   // 返回 nil 表示处理成功
}
