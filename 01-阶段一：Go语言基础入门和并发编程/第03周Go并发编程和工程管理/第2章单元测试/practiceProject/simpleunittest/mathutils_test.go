/**
 * @File : mathutils_test.go
 * @Description : 请填写文件描述
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-09-19
 */
package simpleunittest

import "testing"

func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		a, b int
		exp  int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, -2, -3},
	}
	for _, tt := range tests {
		re := Add(tt.a, tt.b)
		if re != tt.exp {
			t.Errorf("Add(%v, %v) = %v, expect %v", tt.a, tt.b, re, tt.exp)
		}
	}
}

func TestMultiplyTableDriven(t *testing.T) {
	tests := []struct {
		a, b int
		exp  int
	}{
		{1, 2, 2},
		{0, 0, 0},
		{-1, -2, 2},
	}
	for _, tt := range tests {
		re := Multiply(tt.a, tt.b)
		if re != tt.exp {
			t.Errorf("Multiply(%v, %v) = %v, expect %v", tt.a, tt.b, re, tt.exp)
		}
	}
}
