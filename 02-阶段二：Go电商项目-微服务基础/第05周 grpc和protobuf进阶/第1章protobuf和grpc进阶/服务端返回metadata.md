在 gRPC 中，服务端也可以使用 `metadata.AppendToOutgoingContext` 向客户端返回元数据。这通常用于在服务端处理请求后，将额外的信息通过元数据返回给客户端。例如，服务端可以返回一些请求处理的状态信息、认证信息或其他自定义数据。

### 服务端返回元数据的过程

服务端在处理客户端的请求时，可以在处理完成后，使用 `metadata.AppendToOutgoingContext` 将元数据附加到响应的上下文中。然后，客户端可以从响应的元数据中获取这些信息。

### 服务端返回元数据的示例

假设我们有一个简单的 gRPC 服务 `SayHello`，服务端在返回 `HelloResponse` 的同时，还希望返回一些元数据，比如响应的处理时间。

#### 服务端代码示例

```go
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"time"

	"my_project/proto" // 假设这是你的 gRPC 生成代码
)

// 实现 HelloServiceServer 接口
type Server struct {
	proto.UnimplementedHelloServiceServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	// 在处理请求之前生成一些元数据，比如时间戳
	timestamp := time.Now().Format(time.RFC3339)

	// 创建元数据
	md := metadata.Pairs("response-time", timestamp)

	// 将元数据附加到上下文
	newCtx := metadata.AppendToOutgoingContext(ctx, "response-time", timestamp)

	// 打印日志信息
	fmt.Println("Processing request for:", request.Name)

	// 返回响应和上下文
	grpc.SetHeader(newCtx, md) // 将元数据返回给客户端
	return &proto.HelloResponse{Message: "Hello, " + request.Name}, nil
}

func main() {
	// 创建监听器
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 创建 gRPC 服务器
	grpcServer := grpc.NewServer()
	proto.RegisterHelloServiceServer(grpcServer, &Server{})

	// 启动服务器
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
```

#### 代码解读：
- 在 `SayHello` 方法中，我们创建了元数据 `md`，其中键为 `"response-time"`，值为当前的时间戳。
- 然后使用 `metadata.AppendToOutgoingContext` 将元数据附加到上下文 `newCtx` 中。
- `grpc.SetHeader(newCtx, md)` 负责将元数据传递给客户端。

#### 特别注意：

Q:为什么这一步似乎是在反复将md放到上下文中？在创建md的时候已经放了一次数据，附加到ctx里面的时候又放了一次数据，如何最好setheader将md和已经设置好的newctx又放一起了？

**区别**：

1. **`metadata.AppendToOutgoingContext`**：
   - 作用：将一对键值（或多对键值）附加到现有的上下文（`context.Context`）中，形成新的上下文。这通常用于客户端或服务端在 RPC 请求或响应的过程中携带额外信息，例如认证信息或请求相关的元数据。
   - 适用场景：它主要用于携带需要传递的元数据，而这些元数据在调用链条中可能会被进一步传递和处理。

2. **`grpc.SetHeader`**：
   - 作用：直接设置 HTTP/2 头中的元数据，以便服务端在响应给客户端时携带这些数据。这个方法通常在服务端调用，用于明确告诉 gRPC 框架把特定的元数据发送回客户端。
   - 适用场景：它是在 RPC 方法执行过程中明确设置响应的元数据。

**为什么看似重复了**？

在这段代码中：

```go
// 创建元数据
md := metadata.Pairs("response-time", timestamp)

// 将元数据附加到上下文
newCtx := metadata.AppendToOutgoingContext(ctx, "response-time", timestamp)

// 返回响应和上下文
grpc.SetHeader(newCtx, md) // 将元数据返回给客户端
```

1. **`metadata.Pairs("response-time", timestamp)`**：这里创建了一个元数据 `md`，它本质上是一个 `map[string][]string` 结构，用来存储要发送的元数据。
   
2. **`metadata.AppendToOutgoingContext(ctx, "response-time", timestamp)`**：这里将元数据 `response-time` 附加到上下文中，并返回一个新的 `newCtx`。这一步其实是将元数据与当前上下文绑定，以便在调用链中传递给下一个 handler 或传递给 gRPC 库。

3. **`grpc.SetHeader(newCtx, md)`**：这一步显式地将元数据 `md` 设置到 HTTP/2 头中，以便通过 gRPC 的响应传回给客户端。

**问题所在**：

- **重复附加**：这里 `metadata.AppendToOutgoingContext` 和 `grpc.SetHeader` 都分别处理了 `response-time` 的元数据。`AppendToOutgoingContext` 是为了保证上下文中的元数据可被传递，但并未真正发送。而 `SetHeader` 是用于明确发送响应头中的元数据。最终看起来像是重复操作。

**优化方式**：

要避免重复附加元数据的操作，可以只使用 `grpc.SetHeader` 来设置发送给客户端的元数据，而不必再使用 `AppendToOutgoingContext`。`AppendToOutgoingContext` 主要用于链式调用或希望上下文携带的元数据继续传递到下游处理逻辑。

优化后的代码：

```go
// 创建元数据
md := metadata.Pairs("response-time", timestamp)

// 打印日志信息
fmt.Println("Processing request for:", request.Name)

// 直接设置响应头的元数据
grpc.SetHeader(ctx, md) // 将元数据返回给客户端

// 返回响应
return &proto.HelloResponse{Message: "Hello, " + request.Name}, nil
```

### 总结：

- **`AppendToOutgoingContext`** 主要用于将元数据附加到上下文中，以便它可以在调用链中传递。
- **`grpc.SetHeader`** 是明确用于设置 gRPC 响应头中的元数据，以便传回客户端。
- 如果你只是在当前服务端设置元数据返回给客户端，那么只需使用 `grpc.SetHeader` 即可，无需 `AppendToOutgoingContext`。

### 客户端如何接收元数据

客户端可以在发起请求后，通过调用 `grpc.Header()` 方法来获取从服务端返回的元数据。

#### 客户端代码示例

```go
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"

	"my_project/proto" // 假设这是你的 gRPC 生成代码
)

func main() {
	// 创建 gRPC 连接
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// 创建 gRPC 客户端
	client := proto.NewHelloServiceClient(conn)

	// 创建上下文，用于存放响应头的元数据
	md := metadata.MD{}
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// 调用远程方法，并接收响应头中的元数据
	var header metadata.MD
	response, err := client.SayHello(ctx, &proto.HelloRequest{Name: "Junxi"}, grpc.Header(&header))
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	// 打印服务端返回的消息
	fmt.Println("Response Message:", response.Message)

	// 打印元数据中的 "response-time"
	if responseTime, ok := header["response-time"]; ok {
		fmt.Println("Response Time:", responseTime[0])
	} else {
		fmt.Println("No response-time metadata received")
	}
}
```

#### 代码解读：
1. **`grpc.Header(&header)`**：客户端在发起 `SayHello` 请求时，传入了一个 `header` 变量，这个变量用于存放从服务端返回的元数据。
2. **解析元数据**：在收到响应后，通过 `header["response-time"]` 获取服务端返回的 `"response-time"` 元数据，并打印出来。

### 小结
- **`metadata.AppendToOutgoingContext`** 可以用于在 **服务端** 向客户端返回额外的元数据。
- 服务端通过 `grpc.SetHeader` 将元数据附加到响应中，并且客户端通过 `grpc.Header()` 获取这些元数据。
- 这种方式非常适合在请求响应之外传递一些额外的状态信息，如处理时间、认证状态等。