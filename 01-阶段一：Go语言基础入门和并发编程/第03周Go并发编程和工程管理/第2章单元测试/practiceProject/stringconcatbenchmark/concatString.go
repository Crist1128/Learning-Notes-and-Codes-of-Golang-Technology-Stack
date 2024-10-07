/**
 * @File : concatString.go
 * @Description : 请填写文件描述
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-09-19
 */
package stringconcatbenchmark

import (
	"fmt"
	"strings"
)

func ConcatWithPlus(s1, s2 string) string {
	return s1 + s2
}

func ConcatWithSprintf(s1, s2 string) string {
	return fmt.Sprintf("%s%s", s1, s2)
}

func ConcatWithBuilder(s1, s2 string) string {
	var builder strings.Builder
	builder.WriteString(s1)
	builder.WriteString(s2)
	return builder.String()
}
