package ch22

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("id is %d", id)
}

// 仅需任意任务完成 就返回
func FirstResponse() string {
	numOfRunner := 10
	////	Channel: The channel's buffer is initialized with the specified
	////	buffer capacity. If zero, or the size is omitted, the channel is
	////	unbuffered.  设置大小可以buffer channel 防止内存泄露 本来等待的协程为11 现在  为0
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFistResponse(t *testing.T) {
	// NumGoroutine returns the number of goroutines that currently exist.
	// 可以打印协程数量
	t.Log("Before", runtime.NumGoroutine())
	t.Log(FirstResponse())
	t.Log("After", runtime.NumGoroutine())
}
