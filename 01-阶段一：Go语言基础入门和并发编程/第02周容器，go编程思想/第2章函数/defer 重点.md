在 Go 语言中，`defer` 语句用于延迟函数的执行，直到其所在的函数返回时才会被执行。它通常用于执行资源清理操作，例如关闭文件、解锁资源或释放内存等。`defer` 的主要优势在于，即使函数在中途发生错误或提前返回，`defer` 语句仍然会确保被执行。

---

### `defer` 重点提炼表

| 特性              | 描述                                         | 示例                               |
| ----------------- | -------------------------------------------- | ---------------------------------- |
| 执行时机          | `defer` 语句延迟到所在函数即将返回时才会执行 | `defer fmt.Println("done")`        |
| 多个 `defer` 语句 | 多个 `defer` 语句按**后进先出**的顺序执行    | -                                  |
| 捕获值            | `defer` 会立即捕获相关的参数值，但延迟执行   | `defer fmt.Println("Value:", val)` |
| 典型用途          | 资源管理，如关闭文件、解锁互斥锁、回收资源   | `defer file.Close()`               |

---

### 详细说明

#### 1. **基本用法**
`defer` 语句会将它所调用的函数的执行延迟到封装它的函数即将返回时。这对于在函数结束时确保某些清理操作（如关闭文件、解锁资源等）非常有用。

```go
package main

import "fmt"

func main() {
    fmt.Println("Start")
    defer fmt.Println("Done") // 这句将在函数返回时执行
    fmt.Println("Working...")
}
```

**输出:**
```
Start
Working...
Done
```

在这个例子中，`defer` 延迟了 `fmt.Println("Done")` 的执行，直到 `main()` 函数即将结束时才输出 `"Done"`。

#### 2. **多个 `defer` 的执行顺序**
如果有多个 `defer` 语句，它们会按照**后进先出（LIFO, Last In First Out）**的顺序执行。也就是说，最后一个 `defer` 语句会最先执行。

```go
package main

import "fmt"

func main() {
    defer fmt.Println("First")
    defer fmt.Println("Second")
    defer fmt.Println("Third")
    fmt.Println("Working...")
}
```

**输出:**
```
Working...
Third
Second
First
```

#### 3. **`defer` 捕获值**
需要注意的是，`defer` 在声明时会立即评估它所引用的参数或变量的值，但不会立即执行函数。换句话说，`defer` 捕获的值是当时的值，而不是 `defer` 执行时的值。

```go
package main

import "fmt"

func main() {
    val := 10
    defer fmt.Println("Deferred Value:", val)
    val = 20
    fmt.Println("Updated Value:", val)
}
```

**输出:**
```
Updated Value: 20
Deferred Value: 10
```

在这个例子中，尽管 `val` 在后面被修改成了 `20`，但 `defer` 捕获的是 `val` 的初始值 `10`。

#### 4. **常见用例**

##### 4.1 关闭文件
`defer` 常被用于在文件操作完成后关闭文件。

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()  // 确保文件操作结束后关闭文件

    // 读取文件内容或其他操作
}
```

无论文件操作成功或失败，`defer` 都会确保文件在函数结束时被关闭。

##### 4.2 互斥锁解锁
在并发编程中，`defer` 可以确保在使用互斥锁的过程中，即使发生错误也能够正确解锁。

```go
package main

import (
    "fmt"
    "sync"
)

var mu sync.Mutex

func main() {
    mu.Lock()
    defer mu.Unlock()  // 确保函数返回时解锁

    // 处理关键区代码
    fmt.Println("Critical section")
}
```

#### 5. **defer 与匿名函数**
有时候为了延迟执行某些操作并捕获执行时的值，可以结合匿名函数使用 `defer`。

```go
package main

import "fmt"

func main() {
    val := 10
    defer func() {
        fmt.Println("Deferred value with closure:", val)
    }()
    val = 20
    fmt.Println("Updated value:", val)
}
```

**输出:**
```
Updated value: 20
Deferred value with closure: 20
```

这里的匿名函数延迟了执行，并且它捕获了当时的 `val` 的最终值（即 `20`），而不是 `defer` 声明时的值。

---

### `defer` 结合 `panic` 和 `recover`
Go 语言支持 `panic` 和 `recover` 机制，用于异常处理。`defer` 可以与 `recover` 搭配使用，确保在发生 `panic` 时进行适当的清理操作。

```go
package main

import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()

    fmt.Println("Starting the program")
    panic("A critical error occurred!")
    fmt.Println("This will not be printed")
}
```

**输出:**
```
Starting the program
Recovered from: A critical error occurred!
```

当 `panic` 触发时，程序不会立即崩溃，因为 `recover` 在 `defer` 中捕获到了 `panic`，并进行了处理。

---

### 官方文档
可以参考 Go 的[官方文档](https://go.dev/doc/effective_go#defer) 了解 `defer` 语句的更多详细信息和最佳实践。

