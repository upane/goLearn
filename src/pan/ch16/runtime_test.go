package ch16

import (
	"fmt"
	"testing"
	"time"
)

// 协程的调用  方法面前加 go
func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)

		// 会出现竞争 运行有问题
		//go func() {
		// i变为共享变量需要竞争  会出现都是一个数 需要锁
		//	fmt.Println(i)
		//}()
	}
	time.Sleep((time.Second * 3))

}
