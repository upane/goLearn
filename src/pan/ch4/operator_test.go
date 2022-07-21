package ch4

import "testing"

const (
	// 比特位表示 状态码 第一位为可读 第二位可写
	// 连续赋值牛的 常量和位运算都行
	Readable = 1 << iota
	Writable
	Executable
)

// 可比较 相同维度数 数组的值相等
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 2, 3, 4, 5}
	d := c

	// 长度不同不可比较
	// invalid operation: b == c (mismatched types [4]int and [5]int)
	t.Log(a == b, c == d)
}

// &^ 按位清零   右边的二进制为1 则左边的操作数清零
// new 特有功能

func TestBitClear(t *testing.T) {
	a := 7 //0111
	a = a &^ Readable
	// 与运算
	t.Log(a&Readable == Readable, a&Writable == Writable, a)
}
