package ch14

import (
	"errors"
	"testing"
)

// 也就是单个这样定义 较为常用
var ShouldBetweenTwoAndTen = errors.New("n should be in[2,100]")

func GetFibonacci(n int) ([]int, error) {

	if n < 2 || n > 100 {
		return nil, ShouldBetweenTwoAndTen
	}

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil

}

// 避免嵌套 及早失败的角度让代码更加清晰
func TestGetFibonacci(t *testing.T) {
	if err, v := GetFibonacci(200); nil != err {
		//可以判断错误
		t.Error(err)
	} else {
		t.Log(v)
	}
	//n should be in[2,100]
}
