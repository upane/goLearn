package ch9

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFun(t *testing.T) {

	s := "A,B,C"
	// 字符串分割
	part := strings.Split(s, ",")
	for _, c := range part {
		t.Log(c)
	}
	// 字符串连接
	t.Log(strings.Join(part, "z"))
}
func TestName(t *testing.T) {
	// 整数转字符串
	s := strconv.Itoa(10)
	t.Log("str" + s)
	// 字符串转int 需要捕获错误，判断 err
	if i, err := strconv.Atoi("100"); err == nil {
		t.Log(10 + i)
	}
}
