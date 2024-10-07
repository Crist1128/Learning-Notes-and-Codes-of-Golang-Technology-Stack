# RPC 实践笔记

## 重点提炼表

| 项目         | 描述                                                         |
| ------------ | ------------------------------------------------------------ |
| **RPC**      | 远程过程调用，允许在不同地址空间的程序之间调用过程。         |
| **目录结构** | 包括 server、server_proxy、handler、client 和 client_proxy 五个部分。 |
| **信息流**   | 客户端通过 client_proxy 发送请求，server_proxy 注册服务，handler 实现业务逻辑，server 处理连接并调用 handler。 |
| **实现方式** | 使用 Go 的 net/rpc 包进行 TCP 通信，并通过 JSON-RPC 序列化请求和响应。 |
| **最佳实践** | 通过接口进行服务注册和调用，确保代码的可扩展性和可维护性。   |

---

## 1. 目录结构分析

### 1.1 server.go
负责创建 TCP 监听器，注册服务，并接收来自客户端的连接。代码结构如下：

```go
package main

import (
	"grpc_test/full_rpc/handler" // 引入 handler 包以访问服务实现
	"grpc_test/full_rpc/server_proxy" // 引入 server_proxy 包以注册服务
	"net"
	"net/rpc"
)

func main() {
	// 创建监听器，监听 TCP 连接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	// 注册服务
	server_proxy.RegisterHelloService(&handler.HelloServer{})

	// 接收并处理连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue // 忽略错误，继续接受其他连接
		}
		go rpc.ServeConn(conn) // 并发处理每个连接
	}
}
```

### 1.2 server_proxy.go
该文件负责服务的注册。通过接口定义服务，使得代码更加灵活。代码如下：

```go
package server_proxy

import (
	"grpc_test/full_rpc/handler"
	"net/rpc"
)

// HelloServicer 定义了 Hello 服务的接口
type HelloServicer interface {
	Hello(request string, reply *string) error
}

// RegisterHelloService 注册服务到 RPC 框架
func RegisterHelloService(srv HelloServicer) error {
	return rpc.RegisterName(handler.HelloServiceName, srv)
}
```

### 1.3 handler.go
实现具体的服务逻辑，包括处理请求和返回响应。代码如下：

```go
package handler

const HelloServiceName = "handler/HelloService" // 服务名称

type HelloServer struct{}

// Hello 方法实现具体的业务逻辑
func (s *HelloServer) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}
```

### 1.4 client_proxy.go
客户端代理，用于封装 RPC 调用，简化客户端与服务端之间的交互。代码如下：

```go
package client_proxy

import (
	"grpc_test/full_rpc/handler"
	"net/rpc"
)

// HelloServiceStub 用于客户端调用服务
type HelloServiceStub struct {
	*rpc.Client // 嵌入 rpc.Client 以使用其方法
}

// NewHelloServiceStub 创建新的服务代理
func NewHelloServiceStub(protocol, addr string) HelloServiceStub {
	conn, err := rpc.Dial(protocol, addr) // 连接到服务端
	if err != nil {
		panic("connect err!")
	}
	return HelloServiceStub{conn}
}

// Hello 封装 RPC 调用
func (s *HelloServiceStub) Hello(request string, reply *string) error {
	err := s.Call(handler.HelloServiceName+".Hello", request, reply) // 调用服务的方法
	if err != nil {
		return err
	}
	return nil
}
```

### 1.5 client.go
客户端入口，创建客户端代理并发起请求。代码如下：

```go
package main

import (
	"fmt"
	"grpc_test/full_rpc/client_proxy" // 引入客户端代理
)

func main() {
	conn := client_proxy.NewHelloServiceStub("tcp", "127.0.0.1:1234") // 创建客户端代理
	var reply string
	conn.Hello("cc", &reply) // 调用服务
	fmt.Println(reply) // 打印响应
}
```

---

## 2. 信息传递流程

1. **客户端请求**：
   - 客户端通过 `client_proxy` 创建一个连接，构造请求并调用 `Hello` 方法。
   - 请求数据包括方法名和参数。

2. **服务代理处理**：
   - `client_proxy` 中的 `HelloServiceStub` 接收到请求后，调用 `rpc.Client` 的 `Call` 方法。
   - 该方法负责通过网络发送请求到服务端，并等待响应。

3. **服务注册**：
   - 在 `server.go` 中，服务被注册到 RPC 框架，使用 `server_proxy.RegisterHelloService` 方法。
   - 注册服务时指定服务名称和实现该服务的具体逻辑（`handler.HelloServer`）。

4. **请求处理**：
   - 服务端在 `main` 函数中监听 TCP 连接，并在接收到连接后使用 `rpc.ServeConn` 来处理请求。
   - 请求被转发到相应的服务实现（`handler.HelloServer`）的 `Hello` 方法进行处理。

5. **响应返回**：
   - `Hello` 方法处理完请求后，将结果通过 `reply` 参数返回给服务代理。
   - 服务代理将响应返回给客户端，客户端接收并打印响应内容。

### 信息流示意图

```plaintext
客户端                      服务端
+------------+             +------------------+
|            |             |                  |
| client.go  |  发送请求   |  server.go       |
|            | ----------> |  监听连接        |
|            |             |                  |
| client_proxy | <---调用服务|  server_proxy     |
|            |             |                  |
|            | <---处理请求 |  handler          |
|            |             |                  |
|            |  返回响应   |                  |
|            | <----------  |                  |
|            |             |                  |
+------------+             +------------------+
```

---

## 3. 总结

通过将 RPC 实现拆分为多个模块，代码结构更加清晰，易于扩展和维护。各个模块的分离使得信息流动更加高效，有利于团队协作开发和后续的功能扩展。使用接口定义服务可以为将来的服务替换或添加提供便利。