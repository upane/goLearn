package ch10

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 函数是一等公民 go中函数为传值！！！

// 多返回值
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFunc(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
	spent := timeSpent(slowFun)
	t.Log(spent(10))
}

// 可变长参数 函数
func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op

	}
	return ret
}
func TestName(t *testing.T) {
	t.Log(sum(1, 2, 3, 4, 5))
}

// defer（延迟） 函数 相当于finally

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("hhaha") //匿名函数
	}()
	fmt.Println("zzz")
	panic("err") // 在panic后 defer还是会执行
	fmt.Println("mmmm")
}
