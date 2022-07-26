package ch3

import (
	"testing"
)

type Myint int64

// go 对隐式转换是不允许的 可以是自动推断
// 但指定的类型不能隐式转换
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c Myint
	c = Myint(b)
	t.Log(a, b, c)
}

// go不支持指针运算
func TestPoint(t *testing.T) {
	a := 1
	// 获取指针地址
	aPtr := &a
	// 不支持指针运算
	//aPtr += 1  cannot convert 1 (untyped int constant) to *int
	// 1 0xc00000a380
	t.Log(a, aPtr)
	// 别忘记f
	// int *int
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("q" + s + "q")
	t.Log(len(s))
	//    type_test.go:36: qq
	//    type_test.go:37: 0
}
