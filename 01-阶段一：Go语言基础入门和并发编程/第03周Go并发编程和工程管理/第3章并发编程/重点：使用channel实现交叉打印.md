# 重点：使用channel实现交叉打印

#### 题目：
编写两个 Goroutine，一个打印数字，一个打印字母，要求交替输出数字和字母，例如输出 "12AB34CD..."，直到数字 28 和字母 Z 完成。使用无缓冲通道同步 Goroutine 的执行顺序。

```bash
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
```

#### 涉及的知识点：
1. **Goroutine 并发编程**：利用 Goroutine 并发执行任务，允许多个函数或方法同时运行。
2. **通道（Channel）同步**：通道是 Goroutine 之间通信和同步的方式。**无缓冲通道在没有接收方接收时，发送方会阻塞。**
3. **通道死锁（Deadlock）**：当 Goroutine 因为无法接收到通道的信号而阻塞，并且没有其他 Goroutine 能够打破这个循环时，就会发生死锁。
4. **缓冲通道（Buffered Channel）**：缓冲通道允许发送方在通道未满时继续发送数据，即使没有接收方。
5. **happen-before 机制**：确保通道发送和接收的执行顺序，发送操作发生在接收操作之前，通道保证发送和接收之间的同步性。

第一次编写的代码出现了死锁，因为只是单纯地使用两个无缓冲channel进行同步，而忽略了在最后一步输出数字的时候，会向一个没有goroutine使用的channel进行发送而形成了死锁：
```go
package main

import (
	"fmt"
)

// 打印数字
func PrintNum(chNum chan struct{}, chLetter chan struct{}, done chan struct{}) {
	for i := 1; i <= 28; i += 2 {
		<-chNum                    // 等待数字信号
		fmt.Printf("%v%v", i, i+1) // 输出两个数字
		chLetter <- struct{}{}     // 发送字母信号
	}
	done <- struct{}{}
}

// 打印字母
func PrintLetter(chNum chan struct{}, chLetter chan struct{}) {
	for i := 0; i < 26; i += 2 {
		<-chLetter                         // 等待字母信号
		fmt.Printf("%c%c", 'A'+i, 'A'+i+1) // 输出两个字母
		chNum <- struct{}{}                // 发送数字信号
	}
}

func main() {

	chNum := make(chan struct{})       // 数字信号通道
	chLetter := make(chan struct{}, 1) // 字母信号通道
	done := make(chan struct{}, 1)

	// 启动两个 Goroutine，分别用于打印数字和字母
	go PrintNum(chNum, chLetter, done)
	go PrintLetter(chNum, chLetter)

	// 首先发送一个数字信号，开始打印数字
	chNum <- struct{}{}
	<-done

}

```

#### 错误分析：
1. **错误表现**：程序在最后一次循环时陷入死锁，无法完成所有打印任务。
   
2. **错误原因**：
   - 在最后一轮 `PrintNum` 完成数字打印后，仍然试图向 `chLetter` 通道发送信号（`chLetter <- struct{}{}`）。然而，`PrintLetter` 已经在完成最后一轮字母打印后退出循环，不再等待信号。因为通道是无缓冲的，发送方需要等待接收方的接收，但此时没有 Goroutine 接收这个信号，导致 `PrintNum` Goroutine 卡住，形成死锁。
   
3. **为什么会错**：
   - 没有正确处理最后一次循环的信号传递，导致多余的信号发送给已经不再活跃的接收方（`PrintLetter`）。这是因为每次打印的 Goroutine 都严格依赖另一个 Goroutine 的信号来继续进行，在最后一次时没有停止信号传递，导致通道阻塞。

#### 两种解决方案：

##### 1. **控制最后一次不进行通道传输**
   ```go
   // 打印数字
   func PrintNum(chNum chan struct{}, chLetter chan struct{}, done chan struct{}) {
       for i := 1; i <= 28; i += 2 {
           <-chNum                    // 等待数字信号
           fmt.Printf("%v%v", i, i+1) // 输出两个数字
           if i < 27 {                // 控制最后一次不发送信号
               chLetter <- struct{}{}  // 发送字母信号
           }
       }
       done <- struct{}{}  // 通知主 Goroutine 完成
   }
   ```

   - **解决思路**：
     通过在最后一次数字打印后不再向 `chLetter` 发送信号，避免了多余的信号传递。由于最后一次 `PrintLetter` 已经不再接收信号，这样可以避免 `PrintNum` 阻塞在发送信号的操作上，从而避免死锁。
   
   - **解决角度**：
     从**逻辑控制**的角度解决，通过条件判断控制最后一次的信号传递，确保 Goroutine 正常退出，避免阻塞。

##### 2. **使用有缓冲区的通道**
   ```go
   func main() {
       chNum := make(chan struct{})       // 数字信号通道
       chLetter := make(chan struct{}, 1) // 字母信号通道，带缓冲区
       done := make(chan struct{})

       go PrintNum(chNum, chLetter, done)
       go PrintLetter(chNum, chLetter)

       chNum <- struct{}{}  // 首先发送一个数字信号
       <-done  // 等待任务完成
   }
   ```

   - **解决思路**：
     给 `chLetter` 通道增加一个缓冲区（容量为 1）。这样即使在最后一次 `PrintNum` 发送信号时，`chLetter` 没有立即接收，也不会导致阻塞，因为缓冲区可以暂时存储这个信号，避免死锁。
   
   - **解决角度**：
     从**通道机制**的角度解决，通过缓冲区允许通道发送方不必等待接收方立即处理数据，避免了因没有接收方而导致的阻塞。

#### 总结：
- 错误源自于 Goroutine 之间的信号传递在最后一次循环中没有正确处理，导致死锁。
- 第一个方案通过逻辑控制，避免了多余的信号传递，确保 Goroutine 正常退出。
- 第二个方案通过引入缓冲通道，允许最后一次的信号传递能够存储在缓冲区中，不会导致发送方阻塞。