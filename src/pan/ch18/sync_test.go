package ch18

import (
	"fmt"
	"testing"
	"time"
)

// channel and buffer channel difference
func service() string {
	time.Sleep(time.Millisecond * 30)
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
	fmt.Println(service())
	otherTask()
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh) // channel 取数据
	time.Sleep(time.Second * 1)

}
