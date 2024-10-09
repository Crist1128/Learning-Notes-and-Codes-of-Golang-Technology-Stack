[grpc-go/Documentation/grpc-metadata.md at master · grpc/grpc-go (github.com)](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-metadata.md)

### 重点提炼表

| **知识点**             | **描述**                                                     |
| ---------------------- | ------------------------------------------------------------ |
| **gRPC Metadata 概述** | gRPC Metadata 用于传递请求的附加信息，比如认证、追踪信息等。 |
| **底层数据类型**       | Metadata 底层实现为键值对，键名为字符串，键值支持字符串或二进制格式。 |
| **Metadata 传递方式**  | 可以通过 `context.Context` 传递 Metadata，使用拦截器读取和操作。 |
| **Metadata 应用场景**  | 用于认证授权、跟踪请求状态、传递用户数据、协议版本控制等场景。 |
| **使用方式**           | 使用 `metadata.MD` 对象在客户端和服务端添加或读取 Metadata。 |

在 gRPC 中，**metadata**（元数据）是用于传递额外的键值对信息的机制，常用于传递认证、追踪、调试等辅助信息。元数据类似于 HTTP 的头信息，分为 **请求元数据** 和 **响应元数据**，可以在客户端和服务端之间进行传递。

gRPC 中处理元数据有几种常见方式，包括客户端和服务端两端的元数据传递。常用的键值对创建方式有以下几种：

#### 区别

1. **`metadata.New(map[string]string)`**:
   - 这个方法接受的是 `map[string]string` 作为参数，将其转换为 `metadata.MD`（即 `map[string][]string` 的形式）。
   - 每个 `string` 值会被放入一个长度为 1 的 `[]string` 列表中。

2. **`metadata.Pairs(kv ...string)`**:
   - 这个方法接受一系列成对的键值字符串（`kv ...string`），来生成 `metadata.MD`。
   - 每个键可以对应多个值。

#### 1. **通过 `metadata.New` 创建**
   - 使用 `metadata.New` 可以直接创建一个包含键值对的元数据对象。
   - 这种方法适合直接在客户端或服务端创建自定义的元数据。

**用法：**

```go
import (
    "google.golang.org/grpc/metadata"
)

// 创建 metadata 键值对
md := metadata.New(map[string]string{
    "authorization": "Bearer my-token",
    "user-id":       "1234",
})
```

**特点：**
- `metadata.New()` 接收一个 `map[string]string` 类型，方便地将键值对映射到 `metadata.MD` 中。

##### 2. **通过 `metadata.Pairs` 创建**
   - 使用 `metadata.Pairs` 方法可以通过键值对直接创建元数据。
   - 键值对以变长参数形式传入，适用于简单、直接创建多个键值对的场景。

**用法：**

```go
import (
    "google.golang.org/grpc/metadata"
)

// 创建 metadata 键值对，按键值成对提供
md := metadata.Pairs(
    "authorization", "Bearer my-token",
    "user-id", "1234",
)
```

**特点：**
- `metadata.Pairs` 以键值对的形式传递，可以依次传入多个键值对。适合直接、快速创建多对键值对的情况。

##### 3. **通过 `metadata.AppendToOutgoingContext` 在客户端发出请求时追加**
   - `AppendToOutgoingContext` 方法用于将元数据附加到客户端请求的上下文中。
   - 常用于客户端在发起 RPC 请求时，将元数据传递给服务端。

**用法：**

```go
import (
    "context"
    "google.golang.org/grpc/metadata"
)

// 将 metadata 添加到 context 中
ctx := metadata.AppendToOutgoingContext(context.Background(),
    "authorization", "Bearer my-token",
    "user-id", "1234",
)

// 在 gRPC 调用时，将带有元数据的 context 传递
response, err := client.SomeRPCMethod(ctx, request)
```

**特点：**
- `AppendToOutgoingContext` 方法在发起 gRPC 请求时，将元数据附加到 `context` 中，之后的所有 gRPC 请求都会携带这些元数据。
- 它是一种在客户端方便传递元数据的方式，适合在发起请求时动态添加元数据。

##### 4. **通过 `metadata.FromIncomingContext` 获取请求中的元数据**
   - `FromIncomingContext` 用于从服务端的 `context` 中提取客户端传递的元数据。
   - 常用于服务端获取客户端发来的元数据（例如认证 token、追踪 ID 等）。

**用法：**

```go
import (
    "context"
    "google.golang.org/grpc/metadata"
)

// 在服务端的 RPC 方法中提取元数据
func (s *Server) SomeRPCMethod(ctx context.Context, req *pb.Request) (*pb.Response, error) {
    md, ok := metadata.FromIncomingContext(ctx)
    if ok {
        // 获取特定的键值
        if authTokens, exists := md["authorization"]; exists {
            // 使用 authTokens 中的值
            fmt.Println("Authorization token:", authTokens[0])
        }
    }
    // 继续处理请求
    return &pb.Response{}, nil
}
```

**特点：**
- `metadata.FromIncomingContext` 从服务端的上下文中提取元数据，适合用于服务端解析客户端传递的额外信息。

##### 5. **通过 `metadata.NewOutgoingContext` 创建上下文**
   - `NewOutgoingContext` 用于创建一个携带元数据的上下文，适用于客户端发起 gRPC 请求前设置元数据。

**用法：**

```go
import (
    "context"
    "google.golang.org/grpc/metadata"
)

// 创建携带 metadata 的 context
md := metadata.Pairs("authorization", "Bearer my-token")
ctx := metadata.NewOutgoingContext(context.Background(), md)

// 通过携带 metadata 的 context 发送请求
response, err := client.SomeRPCMethod(ctx, request)
```

**特点：**
- 通过 `NewOutgoingContext` 方法创建的上下文中包含元数据，适用于静态场景。

##### 6. **通过 `metadata.AppendToIncomingContext`**
   - 用于服务端或拦截器向现有的元数据上下文中添加额外的信息。

**用法：**

```go
import (
    "context"
    "google.golang.org/grpc/metadata"
)

// 在服务端方法或拦截器中追加元数据
newCtx := metadata.AppendToIncomingContext(ctx, "new-key", "new-value")

// 在后续处理中，可以使用新添加的 metadata
```

##### 总结对比：
- **`metadata.New`**: 通过传入 `map[string]string` 映射批量创建键值对，适合静态配置。
- **`metadata.Pairs`**: 使用键值成对的形式，适合简单创建多个键值对，快速添加。
- **`metadata.AppendToOutgoingContext`**: 在客户端请求时追加元数据，适合动态传递键值对。
- **`metadata.FromIncomingContext`**: 从服务端上下文中提取元数据，用于服务端处理客户端传递的信息。
- **`metadata.NewOutgoingContext`**: 创建包含元数据的上下文，适合客户端提前设置的元数据。
- **`metadata.AppendToIncomingContext`**: 在已有的元数据上下文中追加内容，适合拦截器和服务端场景。

---

### 1. **gRPC Metadata 概述**

在 gRPC 中，`Metadata` 是一种可以在客户端和服务端之间传递的附加信息，类似于 HTTP 的请求头。它允许开发者在不修改消息本体的情况下，传递一些额外的信息。这些信息可能包括认证凭据、请求追踪 ID、用户数据等。

Metadata 在 gRPC 中并不是直接通过请求体传递的，而是作为请求或响应的元数据存在。

#### 应用场景
- **认证授权**：通过 Metadata 传递认证令牌或 API 密钥，服务端从 Metadata 中解析这些信息以验证请求的合法性。
- **分布式追踪**：在微服务架构中，使用 Metadata 传递 Trace ID，可以追踪请求在不同服务间的流转。
- **自定义信息**：可以通过 Metadata 传递客户端或者服务端特有的配置信息，譬如语言、版本号等。

---

### 2. **Metadata 的底层数据类型**

在 gRPC 中，`Metadata` 底层数据结构是键值对的形式。每一个键是一个字符串，值可以是字符串或二进制数据。

- MD数据类型实际上是：

```go
type MD map[string][]string
```

#### 键名格式
- 键名必须是小写的 ASCII 字符串，且只能包含字母、数字、连字符 (`-`)、下划线 (`_`) 和点 (`.`)。
- 键名以 `-bin` 结尾的表示二进制值，其对应的值应为 base64 编码的字符串。

#### 值的类型
- **字符串值**：标准的键可以对应一个或多个字符串值。  
- **二进制值**：以 `-bin` 结尾的键对应 base64 编码的二进制值。

---

### 3. **如何使用 gRPC Metadata**

#### 3.1 在客户端传递 Metadata

在 gRPC 客户端中，Metadata 是通过 `context.Context` 进行传递的。你可以在发起 RPC 调用时，将 Metadata 附加到上下文中。

##### 示例代码

```go
import (
    "context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/metadata"
)

func main() {
    // 创建 gRPC 客户端连接
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    // 创建 Metadata
    md := metadata.Pairs(
        "authorization", "Bearer token",
        "user-id", "12345",
    )

    // 创建新的 context，附加 Metadata
    ctx := metadata.NewOutgoingContext(context.Background(), md)

    // 使用带有 Metadata 的 context 发起 gRPC 请求
    client := pb.NewYourServiceClient(conn)
    response, err := client.YourRPCMethod(ctx, &pb.YourRequest{})
    if err != nil {
        log.Fatalf("RPC failed: %v", err)
    }

    fmt.Println(response)
}
```

在这个例子中，`metadata.Pairs` 函数用来创建键值对，并通过 `metadata.NewOutgoingContext` 将其附加到 context 中。

#### 3.2 在服务端接收 Metadata

在服务端，可以通过 `context.Context` 获取客户端传递的 Metadata。

##### 服务端代码示例

```go
import (
    "context"
    "google.golang.org/grpc/metadata"
)

func (s *server) YourRPCMethod(ctx context.Context, req *pb.YourRequest) (*pb.YourResponse, error) {
    // 从 context 中提取 Metadata
    md, ok := metadata.FromIncomingContext(ctx)
    if !ok {
        return nil, fmt.Errorf("missing metadata")
    }

    // 获取具体的 metadata 值
    if vals, ok := md["authorization"]; ok {
        fmt.Println("Authorization Token:", vals)
    }

    return &pb.YourResponse{}, nil
}
```

服务端可以通过 `metadata.FromIncomingContext` 从上下文中提取 Metadata。这个 Metadata 是客户端传递的。

---

### 4. **Metadata 的传递方式**

gRPC Metadata 通过 `context.Context` 进行传递，它可以携带在请求发起前附加的所有键值对。`metadata.MD` 类型是一个 map，可以存储多值键。

Metadata 可以在请求时通过上下文传递，也可以在响应中返回。这种设计让 Metadata 适用于各种需求，比如认证、追踪、定制信息等。

---

### 5. **gRPC Metadata 的应用场景**

- **认证和授权**：可以通过 Metadata 携带身份验证信息，比如 OAuth 的令牌或 API 密钥。
- **追踪和监控**：在分布式系统中，利用 Metadata 传递 Trace ID 和 Span ID，帮助跟踪整个请求链路。
- **自定义标头**：可以通过 Metadata 传递一些自定义的元数据，比如请求的版本号、语言等，这些可以帮助服务端处理个性化请求。
- **日志记录**：Metadata 中的一些信息如请求 ID、用户 ID 等，可以在日志记录时附加，便于问题排查。

---

### 6. **gRPC Metadata 最佳实践**

- **小心传递敏感信息**：虽然 Metadata 可以用来传递认证信息，但一定要确保安全，避免将敏感信息以明文方式传递。
- **避免 Metadata 太大**：过多的 Metadata 会增加网络开销，尽量保持 Metadata 的精简。
- **使用拦截器处理 Metadata**：在 gRPC 中，可以通过拦截器集中处理 Metadata，避免在每个方法中重复编写代码。

---

### 总结

gRPC Metadata 是一种灵活的机制，允许在 gRPC 请求和响应中传递额外的元数据。它是通过键值对的形式实现的，键名为字符串，值可以是字符串或二进制格式。Metadata 在 gRPC 中的使用非常广泛，尤其在认证、跟踪请求、传递用户信息等场景中非常有用。

通过 `context.Context`，可以在客户端发送 Metadata，在服务端接收 Metadata，并在必要时返回响应 Metadata。这种灵活性使得 Metadata 成为 gRPC 开发中的重要工具。