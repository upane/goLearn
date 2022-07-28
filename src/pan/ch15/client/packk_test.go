package client

import (
	"fmt"
	"goLearn/src/pan/ch15/series"
	"testing"
)

// 包引用测试  首字母小写 的申明方法不能在go外访问 会未定义错误
func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(10))
}

// init 在所有包里都会main之前执行
// init可能有多个init
func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println("init222")
}
