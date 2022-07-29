package ch19

import (
	"fmt"
	"testing"
	"time"
)

// select 多路选择 可以实现超时
// case 为并排的关系  非顺序，谁返回早就直接返回
func service() string {
	time.Sleep(time.Millisecond * 200)
	return "Done"
}
func AsyncService() chan string {
	//retCh := make(chan string)    // 携程会被阻塞
	retCh := make(chan string, 1) // buffer channel 不会被阻塞

	go func() {
		ret := service()
		fmt.Println("return result")
		retCh <- ret                  // channel 放数据
		fmt.Println("service exited") // 这里会存在阻塞和不阻塞
	}()
	return retCh
}
func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("task is done")
}
func TestService(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")

	}
}
