### `strings` 包常用函数总结表

| 函数        | 描述                                                 | 示例使用                                    | 示例输出                  |
| ----------- | ---------------------------------------------------- | ------------------------------------------- | ------------------------- |
| `Contains`  | 判断子字符串是否存在于字符串中                       | `strings.Contains("hello", "ll")`           | `true`                    |
| `Count`     | 统计子字符串在字符串中出现的次数                     | `strings.Count("hello", "l")`               | `2`                       |
| `Split`     | 按指定的分隔符将字符串拆分成切片                     | `strings.Split("a,b,c", ",")`               | `[]string{"a", "b", "c"}` |
| `HasPrefix` | 判断字符串是否以指定前缀开头                         | `strings.HasPrefix("hello", "he")`          | `true`                    |
| `HasSuffix` | 判断字符串是否以指定后缀结尾                         | `strings.HasSuffix("hello", "lo")`          | `true`                    |
| `Index`     | 返回子字符串在字符串中第一次出现的位置               | `strings.Index("hello", "l")`               | `2`                       |
| `IndexRune` | 返回给定 Unicode 字符在字符串中第一次出现的位置      | `strings.IndexRune("gophers", 'h')`         | `3`                       |
| `Replace`   | 将字符串中指定的子串替换为新的子串，支持替换次数控制 | `strings.Replace("oink oink", "o", "a", 1)` | `aink oink`               |
| `ToLower`   | 将字符串中的所有字符转换为小写                       | `strings.ToLower("GoLang")`                 | `golang`                  |
| `ToUpper`   | 将字符串中的所有字符转换为大写                       | `strings.ToUpper("GoLang")`                 | `GOLANG`                  |
| `Trim`      | 去除字符串两端的指定字符（默认为空格）               | `strings.Trim("  hello  ", " ")`            | `hello`                   |

---

### 详细说明与示例

#### 1. **`Contains`**
`Contains` 用于判断子字符串是否存在于另一个字符串中。如果存在则返回 `true`，否则返回 `false`。

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.Contains("hello", "ll"))   // 输出: true
    fmt.Println(strings.Contains("hello", "world")) // 输出: false
}
```

#### 2. **`Count`**
`Count` 返回子字符串在父字符串中出现的次数。

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.Count("hello", "l"))  // 输出: 2
    fmt.Println(strings.Count("hello", "x"))  // 输出: 0
}
```

#### 3. **`Split`**
`Split` 根据指定的分隔符将字符串拆分成切片。 

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    result := strings.Split("a,b,c", ",")
    fmt.Println(result)  // 输出: [a b c]
}
```

#### 4. **`HasPrefix`**
`HasPrefix` 用于判断字符串是否以指定的前缀开头。

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.HasPrefix("hello", "he"))  // 输出: true
    fmt.Println(strings.HasPrefix("hello", "lo"))  // 输出: false
}
```

#### 5. **`HasSuffix`**
`HasSuffix` 用于判断字符串是否以指定的后缀结尾。

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.HasSuffix("hello", "lo"))  // 输出: true
    fmt.Println(strings.HasSuffix("hello", "he"))  // 输出: false
}
```

#### 6. **`Index`**
`Index` 返回子字符串在父字符串中第一次出现的位置。如果未找到子字符串，返回 `-1`。

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.Index("hello", "l"))  // 输出: 2
    fmt.Println(strings.Index("hello", "x"))  // 输出: -1
}
```

#### 7. **`IndexRune`**
`IndexRune` 返回给定 Unicode 字符在字符串中的第一次出现位置。

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.IndexRune("gophers", 'h'))  // 输出: 3
}
```

#### 8. **`Replace`**
`Replace` 用于将字符串中的部分内容替换为新的字符串。可以控制替换的次数，若设置为 `-1` 则全部替换。

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.Replace("oink oink", "o", "a", 1)) // 输出: aink oink
    fmt.Println(strings.Replace("oink oink", "o", "a", -1)) // 输出: aink aink
}
```

#### 9. **`ToLower`**
`ToLower` 将字符串中的所有字母转换为小写。

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.ToLower("GoLang"))  // 输出: golang
}
```

#### 10. **`ToUpper`**
`ToUpper` 将字符串中的所有字母转换为大写。

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.ToUpper("GoLang"))  // 输出: GOLANG
}
```

#### 11. **`Trim`**
`Trim` 去除字符串两端的指定字符，默认去除空格。

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.Trim("  hello  ", " "))  // 输出: hello
    fmt.Println(strings.Trim("!!!hello!!!", "!"))  // 输出: hello
}
```

---

### 官方文档
你可以查阅 [strings 包官方文档](https://pkg.go.dev/strings) 以获取更多关于这些函数的详细信息和最新更新。

如果你需要深入了解其他内容或继续整理其他模块的笔记，随时告诉我！