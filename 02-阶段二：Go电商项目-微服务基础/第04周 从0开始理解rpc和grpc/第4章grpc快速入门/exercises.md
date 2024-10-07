# 知识点：gRPC 服务器流式传输

**软件包名：grpc_server_streaming**

#### 题目：gRPC 服务器流式传输的简单练习

**题目描述：**
编写一个 gRPC 服务器与客户端程序，实现服务器流式传输（Server Streaming）的基本功能。客户端发起一次请求，服务器连续发送多条响应消息，客户端会一直接收消息直到服务器结束流。

**要求：**
1. 创建一个 gRPC 服务，其中包含一个服务器流式方法 `StreamNumbers`，该方法接受客户端发送的请求，返回一系列连续的整数。
2. 服务器：当服务器接收到客户端请求后，发送从 `1` 到 `10` 的整数流，每次发送一个数字并等待 1 秒后继续发送。
3. 客户端：客户端发起请求后，循环接收服务器发来的数字，并将其打印到控制台。直到流结束。
4. 使用 `sync.WaitGroup` 确保所有消息被正确处理后程序退出。

**提示：**
- 使用 `proto` 文件定义 `StreamNumbers` RPC 服务。
- 在服务器端实现 `StreamNumbers` 方法，并使用 `stream.Send()` 向客户端发送消息。
- 在客户端使用 `stream.Recv()` 接收消息。

**输出示例：**
服务器端：
```bash
Sending number: 1
Sending number: 2
Sending number: 3
...
Sending number: 10
```

客户端：
```bash
Received number: 1
Received number: 2
Received number: 3
...
Received number: 10
```

完成这道题后提交代码，我会帮你审阅！

```go
syntax = "proto3";

option go_package = ".;proto";

service Greeter {
  rpc StreamNumbers(StreamRequest) returns (stream StreamResponse);
}

message StreamRequest {
  string data = 1;
}

message StreamResponse {
  string data = 1;
}
```

```go
/**
 * @File : server.go
 * @Description : 请填写文件描述
 * @Author : Junxi You
 * @Date : 2024-10-07
 */
package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_protoc/grpc_server_streaming/proto"
	"net"
	"strconv"
	"time"
)

// server 结构体实现了 proto 定义的 Greeter 服务
type server struct {
	proto.UnimplementedGreeterServer
}

// 必须包含未实现的嵌入方法（强制要求）
func (s *server) mustEmbedUnimplementedGreeterServer() {
	//TODO implement me
	panic("implement me")
}

// StreamNumbers 是服务器流式传输的核心逻辑
// 它会向客户端发送 1 到 10 的连续数字，逐次发送后等待 1 秒
func (s *server) StreamNumbers(req *proto.StreamRequest, res proto.Greeter_StreamNumbersServer) error {
	i := 1
	for {
		// 如果 i 大于 10，结束流式传输
		if i > 10 {
			break
		}
		// 打印服务器正在发送的数字
		fmt.Println("Sending number: " + strconv.Itoa(i))

		// 通过 res.Send 方法向客户端发送 StreamResponse 消息
		err := res.Send(&proto.StreamResponse{
			Data: fmt.Sprintf("%d", i), // 将数字 i 转换为字符串并放入消息中
		})
		if err != nil {
			return err // 传输错误处理
		}

		i++
		time.Sleep(time.Second) // 每次发送后等待 1 秒
	}
	return nil
}

func main() {
	// 监听 50051 端口，准备接受 gRPC 请求
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	// 创建 gRPC 服务器
	s := grpc.NewServer()

	// 注册 Greeter 服务到服务器
	proto.RegisterGreeterServer(s, &server{})

	// 启动服务器，监听传入的 gRPC 请求
	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}
}

```

```go
/**
 * @File : client.go
 * @Description : 请填写文件描述
 * @Author : Junxi You
 * @Date : 2024-10-07
 */
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_protoc/grpc_server_streaming/proto"
	"io"
)

func main() {
	// 连接到 gRPC 服务器，使用不安全凭证（没有 TLS 加密）
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))//注意，dial已经弃用
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 创建 Greeter 客户端
	c := proto.NewGreeterClient(conn)

	// 调用 StreamNumbers 以开始服务器流
	r, err := c.StreamNumbers(context.Background(), &proto.StreamRequest{})
	if err != nil {
		panic(err)
	}
	// 接收来自服务器的消息流，直到流结束
	for {
		a, err := r.Recv()
		if err == io.EOF {
			fmt.Println("All numbers received.")
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("Received number: " + a.Data)
	}

}

```

### 代码检查和解释

你的代码已经接近完成了一个简单的 gRPC 服务器流式传输（Server Streaming）的功能，服务器将从 1 到 10 的数字发送给客户端，客户端接收并打印这些数字。以下是代码的详细解释及建议改进：

---

### `server.go` - 服务器端代码

```go
package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_protoc/grpc_server_streaming/proto"
	"net"
	"strconv"
	"time"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) mustEmbedUnimplementedGreeterServer() {
	// TODO implement me
	panic("implement me")
}
```

- 这里实现了 `server` 结构体，并嵌入了 `proto.UnimplementedGreeterServer`，这是为了兼容 gRPC 生成的代码。在较新的版本中，你必须显式嵌入这个未实现的服务接口。
- `mustEmbedUnimplementedGreeterServer` 方法是生成的一个默认方法，它并不直接参与业务逻辑，可以保持不变。

---

```go
func (s *server) StreamNumbers(req *proto.StreamRequest, res proto.Greeter_StreamNumbersServer) error {
	i := 1
	for {
		if i > 10 {
			break
		}
		fmt.Println("Sending number: " + strconv.Itoa(i))
		err := res.Send(&proto.StreamResponse{
			Data: fmt.Sprintf("%d", i),
		})
		if err != nil {
			return err
		}
		i++
		time.Sleep(time.Second) // 每发送一次数据暂停1秒
	}
	return nil
}
```

- `StreamNumbers` 是核心的 gRPC 服务方法，负责处理服务器流式传输的逻辑。
- 每次客户端请求该方法时，服务器会循环从 `1` 发送到 `10`，每次发送后暂停 1 秒。
- 通过 `res.Send(&proto.StreamResponse)` 发送消息，数据包装在 `StreamResponse` 消息中。`Send` 返回一个 `error`，如果出错将终止流的传输。
- 当所有消息发送完毕，`for` 循环结束，返回 `nil`，表示成功完成流式传输。

---

```go
func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}
}
```

- 服务器监听 `50051` 端口等待客户端的 gRPC 请求。
- `grpc.NewServer()` 创建了一个 gRPC 服务器实例。
- `proto.RegisterGreeterServer(s, &server{})` 注册了我们自定义的 `server` 实例。
- `s.Serve(listen)` 启动 gRPC 服务器，等待客户端连接。

---

### `client.go` - 客户端代码

```go
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_protoc/grpc_server_streaming/proto"
	"io"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
```

- `grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))` 创建了一个到 gRPC 服务器的连接，这里使用了不加密的连接（insecure credentials），适用于开发环境。
- 连接成功后，创建 `GreeterClient` 客户端实例，用来发起 RPC 请求。

---

```go
	r, err := c.StreamNumbers(context.Background(), &proto.StreamRequest{})
	for {
		a, err := r.Recv()
		if err == io.EOF {
			fmt.Println("All numbers received.")
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("Received number: " + a.Data)
	}
}
```

- `c.StreamNumbers(context.Background(), &proto.StreamRequest{})` 向服务器发送请求，获取一个流式的响应（流）。
- 使用 `r.Recv()` 循环接收服务器发来的流数据，每次接收到数据后打印。
- `r.Recv()` 返回 `io.EOF` 表示流结束。
- 如果有错误发生，直接 `panic`，打印错误堆栈。

---

### `proto` 文件

```proto
syntax = "proto3";

option go_package = ".;proto";

service Greeter {
  rpc StreamNumbers(StreamRequest) returns (stream StreamResponse);
}

message StreamRequest {
  string data = 1;
}

message StreamResponse {
  string data = 1;
}
```

- 定义了 `Greeter` 服务，其中包含一个 `StreamNumbers` 方法，这个方法接受 `StreamRequest` 请求，并返回流式的 `StreamResponse` 响应。
- 服务器流式传输通过 `returns (stream StreamResponse)` 实现，表示服务器将返回一个流，而不是一次性返回所有数据。

---

### 改进建议

1. **客户端的连接方式有误**：
   - 你在客户端使用了 `grpc.NewClient`，应该改为 `grpc.Dial`，以连接服务器：
   
   ```go
   conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
   ```

2. **客户端接收流的循环逻辑：**
   - 你已经正确实现了流的接收逻辑。需要确保服务器发送的消息数量和客户端预期的一致，若有超时或其他原因导致流中断，应处理错误。

3. **代码注释**：
   - 在正式项目中，注释文件头中的描述部分应填写详细内容，比如文件功能、依赖关系等。

---

### 完整注释示例：

```go
// server.go
package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_protoc/grpc_server_streaming/proto"
	"net"
	"strconv"
	"time"
)

// server 是 gRPC 服务器的实现，继承了 proto 中定义的接口
type server struct {
	proto.UnimplementedGreeterServer
}

// StreamNumbers 实现服务器流式传输，从1发送到10，每次发送后暂停1秒
func (s *server) StreamNumbers(req *proto.StreamRequest, res proto.Greeter_StreamNumbersServer) error {
	i := 1
	for {
		if i > 10 {
			break
		}
		fmt.Println("Sending number: " + strconv.Itoa(i))
		err := res.Send(&proto.StreamResponse{
			Data: fmt.Sprintf("%d", i),
		})
		if err != nil {
			return err
		}
		i++
		time.Sleep(time.Second) // 模拟延迟
	}
	return nil
}

func main() {
	// 启动 TCP 监听
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer() // 创建 gRPC 服务器实例
	proto.RegisterGreeterServer(s, &server{}) // 注册服务
	err = s.Serve(listen) // 开始服务
	if err != nil {
		panic(err)
	}
}
```

# 知识点：gRPC 客户端流式传输

**软件包名：grpc_client_streaming**

#### 题目：gRPC 简单客户端流式传输练习

**题目描述：**
编写一个简单的 gRPC 服务和客户端，使用客户端流式传输的方式向服务器发送多个请求数据。服务器在接收完所有数据后，返回一个总结结果给客户端。

**要求：**
1. 定义一个 gRPC 服务 `SumService`，包含一个 `StreamSum` 方法，接受客户端流式传输的数据，并在客户端完成发送后，返回所有数值的总和。
2. 客户端使用流式传输方式发送一组整数给服务器。
3. 服务端在客户端发送完数据后，计算所有接收到的整数的和，并将结果返回给客户端。
4. 使用 `proto` 文件定义服务和消息格式，并生成相应的 gRPC 代码。
5. 编写服务端和客户端的 Go 实现，确保客户端能够流式发送多条数据给服务端，并正确接收服务端返回的总和结果。

**proto 文件示例：**
```proto
syntax = "proto3";

service SumService {
    rpc StreamSum(stream SumRequest) returns (SumResponse);
}

message SumRequest {
    int32 number = 1;
}

message SumResponse {
    int32 sum = 1;
}
```

**功能要求：**
- 客户端能够通过流的形式发送多次 `SumRequest` 给服务器。
- 服务端计算所有接收到的 `SumRequest` 中的 `number` 之和，并返回 `SumResponse` 给客户端。
- 客户端接收到服务器的总和后，打印出总和结果。

**输出示例：**
客户端输出：
```bash
Sending: 1
Sending: 2
Sending: 3
Sending: 4
Received sum: 10
```

服务端输出：
```bash
Received number: 1
Received number: 2
Received number: 3
Received number: 4
Returning sum: 10
```

**提示：**
- 使用 `proto` 文件定义 `SumService` 和消息格式。
- 客户端通过 `gRPC` 的流式调用发送多个 `SumRequest`。
- 服务端接收流数据并计算和，返回 `SumResponse`。

完成这道题后，提交代码我会帮你审阅！

```go
syntax = "proto3";

option go_package=".;proto";

service SumService {
  rpc StreamSum(stream SumRequest) returns (SumResponse);
}

message SumRequest {
  int32 number = 1;
}

message SumResponse {
  int32 sum = 1;
}

```

```go
/**
 * @File : client.go
 * @Description : 请填写文件描述
 * @Author : Junxi You
 * @Date : 2024-10-07
 */
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_protoc/grpc_client_streaming/proto"
	"time"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewSumServiceClient(conn)

	r, err := c.StreamSum(context.Background())
	if err != nil {
		panic(err)
	}
	for i := 1; i <= 10; i++ {
		err := r.Send(&proto.SumRequest{Number: int32(i)})
		if err != nil {
			panic(err)
		}
		fmt.Println("Sent number: " + fmt.Sprintf("%d", i))
		time.Sleep(time.Second)
	}
	a, err := r.CloseAndRecv()
	if err != nil {
		panic(err)
	}
	fmt.Println("Received sum: " + fmt.Sprint(a.Sum))
}

```

```go
/**
 * @File : server.go
 * @Description : 请填写文件描述
 * @Author : Junxi You
 * @Date : 2024-10-07
 */
package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_protoc/grpc_client_streaming/proto"
	"io"
	"net"
)

type server struct {
	proto.UnimplementedSumServiceServer
}

func (s *server) StreamSum(stream proto.SumService_StreamSumServer) error {
	var sum int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += req.Number
		fmt.Println("Received number: " + fmt.Sprint(req.Number))
	}
	fmt.Println("Returning sum: " + fmt.Sprint(sum))
	return stream.SendAndClose(&proto.SumResponse{Sum: sum})
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterSumServiceServer(s, &server{})
	s.Serve(listen)
}

```

你的代码符合题目需求，实现了一个 gRPC 客户端流式传输的例子，其中客户端发送一系列 `SumRequest` 消息，服务器接收并返回计算后的总和结果。整体逻辑是正确的。以下是详细的注释和一些补充说明：

### 客户端代码 `client.go` 注释：

```go
/**
 * @File : client.go
 * @Description : gRPC 客户端，发送 1 到 10 的数字到服务器，接收服务器返回的总和
 * @Author : Junxi You
 * @Date : 2024-10-07
 */
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_protoc/grpc_client_streaming/proto"
	"time"
)

func main() {
    // 连接到 gRPC 服务器
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// 如果连接失败，打印错误信息并终止程序
		panic(err)
	}
	defer conn.Close() // 程序结束时关闭连接

    // 创建一个 SumService 客户端
	c := proto.NewSumServiceClient(conn)

    // 开启客户端流式传输，通过 StreamSum 向服务器发送一系列请求
	stream, err := c.StreamSum(context.Background())
	if err != nil {
		// 如果流创建失败，打印错误信息并终止程序
		panic(err)
	}

    // 模拟发送 1 到 10 的整数到服务器，每发送一次 Sleep 1 秒
	for i := 1; i <= 10; i++ {
		// 将每个数字封装成 SumRequest 并发送到服务器
		err := stream.Send(&proto.SumRequest{Number: int32(i)})
		if err != nil {
			// 发送失败时终止程序
			panic(err)
		}
		fmt.Println("Sent number: " + fmt.Sprintf("%d", i)) // 打印发送的数字
		time.Sleep(time.Second) // 模拟延迟，等待一秒
	}

    // 关闭发送流并接收服务器的返回结果
	response, err := stream.CloseAndRecv()
	if err != nil {
		// 处理关闭流或接收响应时的错误
		panic(err)
	}
	
	// 打印从服务器接收到的总和
	fmt.Println("Received sum: " + fmt.Sprint(response.Sum))
}
```

### 关键点：
1. **gRPC 客户端连接**：通过 `grpc.Dial` 创建与服务器的连接。
2. **流式传输**：客户端通过 `stream.Send()` 多次发送数据给服务器。
3. **流关闭和结果接收**：`stream.CloseAndRecv()` 用来结束客户端的发送，并接收服务器的最终响应。

---

### 服务器代码 `server.go` 注释：

```go
/**
 * @File : server.go
 * @Description : gRPC 服务器，接收客户端流式传输的数字并返回总和
 * @Author : Junxi You
 * @Date : 2024-10-07
 */
package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_protoc/grpc_client_streaming/proto"
	"io"
	"net"
)

// 定义 gRPC 服务器结构体，继承生成的 SumService 接口
type server struct {
	proto.UnimplementedSumServiceServer
}

// 实现 StreamSum 方法，接收客户端流式传输并返回计算后的总和
func (s *server) StreamSum(stream proto.SumService_StreamSumServer) error {
	var sum int32 // 用于计算总和的变量
	for {
		// 从客户端接收数据
		req, err := stream.Recv()
		if err == io.EOF {
			// 当客户端关闭流时，结束接收，退出循环
			break
		}
		if err != nil {
			// 处理接收过程中发生的错误
			return err
		}
		// 累加收到的数字
		sum += req.Number
		fmt.Println("Received number: " + fmt.Sprint(req.Number)) // 打印接收到的数字
	}
	
	// 当所有数据接收完毕后，发送响应并关闭流
	fmt.Println("Returning sum: " + fmt.Sprint(sum))
	return stream.SendAndClose(&proto.SumResponse{Sum: sum}) // 返回计算的总和
}

func main() {
    // 启动 TCP 监听器，监听 50051 端口
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	
	// 创建 gRPC 服务器实例
	s := grpc.NewServer()
	
	// 注册 SumService 服务
	proto.RegisterSumServiceServer(s, &server{})
	
	// 启动服务器，监听来自客户端的请求
	s.Serve(listen)
}
```

### 关键点：
1. **流接收**：服务器通过 `stream.Recv()` 循环接收客户端发来的每一个 `SumRequest`。
2. **总和计算**：服务器接收每个 `SumRequest` 中的数字，累加起来。
3. **返回结果**：服务器通过 `stream.SendAndClose()` 返回计算的总和，并关闭连接。

---

### proto 文件内容：
```proto
syntax = "proto3";

option go_package=".;proto";

// 定义 SumService 服务，包含一个客户端流式传输的方法 StreamSum
service SumService {
  rpc StreamSum(stream SumRequest) returns (SumResponse);
}

// 请求消息：包含一个整型字段 number
message SumRequest {
  int32 number = 1;
}

// 响应消息：包含计算后的总和
message SumResponse {
  int32 sum = 1;
}
```

### 总结：
你的客户端和服务器代码实现了 gRPC 客户端流式传输的功能。客户端发送多个数字到服务器，服务器接收所有数字后返回它们的总和。代码逻辑清晰且符合 gRPC 流式传输的标准用法，建议你在测试时，确保客户端和服务器能够正确通信并传输数据。

#### 改进建议：
- **错误处理**：在 `CloseAndRecv` 和 `Send` 之后，可以加入更多的错误处理和日志记录，确保对潜在的网络或流错误有更好的控制。