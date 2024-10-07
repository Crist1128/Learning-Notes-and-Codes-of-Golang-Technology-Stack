在 Go 语言的并发编程中，**`sync.WaitGroup`** 是一个用于同步多个 Goroutine 的计数器工具，常用来等待一组 Goroutine 完成。通过 `sync.WaitGroup`，可以确保主 Goroutine（或其他 Goroutine）在其他并发 Goroutine 运行完毕之前不会退出。

### `sync.WaitGroup` 的基本操作

`sync.WaitGroup` 提供了三个主要的方法：
- **`Add(delta int)`**：添加或减少等待的 Goroutine 计数。
- **`Done()`**：每当一个 Goroutine 完成时，调用 `Done()` 减少等待的 Goroutine 计数。
- **`Wait()`**：阻塞主 Goroutine，直到所有 Goroutine 完成（即 `Add()` 计数为零）。

### `sync.WaitGroup` 的基本工作流程

1. 调用 `Add()` 方法设置需要等待的 Goroutine 数量。
2. 各个 Goroutine 执行时，每当完成任务后调用 `Done()` 方法。
3. 主 Goroutine 调用 `Wait()`，直到所有 Goroutine 调用 `Done()` 并使计数器归零，主 Goroutine 才继续执行。

---

### `sync.WaitGroup` 的重点提炼表

| 方法         | 描述                                                         | 示例代码    |
| ------------ | ------------------------------------------------------------ | ----------- |
| **`Add(n)`** | 增加需要等待的 Goroutine 数量，`n` 是增量（可以为负数）      | `wg.Add(1)` |
| **`Done()`** | 每个 Goroutine 任务完成时调用 `Done()`，相当于 `Add(-1)`     | `wg.Done()` |
| **`Wait()`** | 阻塞，直到所有 Goroutine 完成任务，`Add()` 计数归零时解除阻塞 | `wg.Wait()` |

---

### 示例：基本使用 `sync.WaitGroup`

在这个例子中，主 Goroutine 启动了两个 Goroutine，并通过 `sync.WaitGroup` 等待它们完成：

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()  // 任务完成后调用 Done
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)  // 模拟工作
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup  // 创建 WaitGroup 实例

    wg.Add(2)  // 需要等待两个 Goroutine

    go worker(1, &wg)  // 启动第一个 Goroutine
    go worker(2, &wg)  // 启动第二个 Goroutine

    wg.Wait()  // 等待两个 Goroutine 完成
    fmt.Println("All workers done")
}
```

#### **解释**：
1. **`wg.Add(2)`**：表示我们将启动两个 Goroutine，需要等待这两个 Goroutine 完成。
2. **`go worker(1, &wg)`** 和 **`go worker(2, &wg)`**：启动两个 Goroutine，它们会执行 `worker` 函数。
3. **`defer wg.Done()`**：确保 Goroutine 完成后调用 `Done()` 减少计数。
4. **`wg.Wait()`**：阻塞主 Goroutine，直到所有的 `worker` Goroutine 完成，才继续往下执行。

### 输出示例：

```
Worker 1 starting
Worker 2 starting
Worker 1 done
Worker 2 done
All workers done
```

---

### 使用 `sync.WaitGroup` 的注意事项

1. **确保每个 Goroutine 都调用 `Done()`**：
   - 每个 Goroutine 执行完毕时必须调用 `Done()`，否则 `Wait()` 永远不会结束，程序可能会因为等待而挂起。

2. **正确设置 `Add()` 调用次数**：
   - `Add()` 必须在所有 Goroutine 启动之前调用，以确保 `Wait()` 等待正确的数量。
   - `Add()` 应在启动 Goroutine 之前调用，避免主 Goroutine 过早调用 `Wait()` 而导致死锁。

3. **避免 `WaitGroup` 重复使用**：
   - 一个 `sync.WaitGroup` 通常在一个并发过程完成后不能再复用（除非你明确了解如何重置和重新使用它）。因此在每次并发任务开始时，应该创建新的 `WaitGroup` 实例。

---

### 使用 `WaitGroup` 同时处理多个任务

在实际应用中，`sync.WaitGroup` 通常用于等待一组并发任务完成。下面是一个使用 `WaitGroup` 处理多个并发任务的示例：

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d started\n", id)
    time.Sleep(time.Second)  // 模拟任务执行
    fmt.Printf("Worker %d finished\n", id)
}

func main() {
    var wg sync.WaitGroup
    workerCount := 5

    // 启动多个并发任务
    for i := 1; i <= workerCount; i++ {
        wg.Add(1)  // 启动一个 Goroutine 前增加计数
        go worker(i, &wg)
    }

    wg.Wait()  // 等待所有 Goroutine 完成
    fmt.Println("All workers finished")
}
```

#### **解释**：
- 这个例子启动了 `5` 个并发的 `worker` 任务，每个任务执行完毕后调用 `wg.Done()`，主程序使用 `wg.Wait()` 等待所有 `worker` 任务完成。

---

### 总结

1. **`sync.WaitGroup`**：是 Go 并发编程中用于等待多个 Goroutine 完成任务的同步原语。
2. **主要方法**：
   - `Add(n)`：增加需要等待的 Goroutine 数量。
   - `Done()`：每个 Goroutine 完成后调用，减少计数。
   - `Wait()`：阻塞主 Goroutine，直到所有 Goroutine 完成。
3. **使用场景**：适合需要同步多个并发任务的场景，确保所有任务完成后继续执行主程序。

通过 `sync.WaitGroup`，Go 能够高效地同步 Goroutine 之间的执行，确保程序能够正确处理并发逻辑。如果你有更多问题，随时告诉我！