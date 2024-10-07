/**
 * @File : concat_test.go
 * @Description : 请填写文件描述
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-09-19
 */
package stringconcatbenchmark

import (
	"strconv"
	"testing"
)

func BenchmarkConcatWithPlus(b *testing.B) {
	if b.N > 100000 {
		b.N = 100000
	}
	b.ResetTimer()
	str := ""
	for i := 0; i < b.N; i++ {
		str = ConcatWithPlus(str, strconv.Itoa(i))
	}
}

func BenchmarkConcatWithSprintf(b *testing.B) {
	if b.N > 100000 {
		b.N = 100000
	}
	b.ResetTimer()
	str := ""
	for i := 0; i < b.N; i++ {
		str = ConcatWithSprintf(str, strconv.Itoa(i))
	}
}

func BenchmarkConcatWithBuilder(b *testing.B) {
	if b.N > 100000 {
		b.N = 100000
	}
	b.ResetTimer()
	str := ""
	for i := 0; i < b.N; i++ {
		str = ConcatWithBuilder(str, strconv.Itoa(i))
	}
}
