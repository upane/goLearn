package ch20

import (
	"fmt"
	"sync"
	"testing"
)

// 生产者
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 7; i++ {
			ch <- i
		}
		close(ch) // channel可以被关闭
		//ch <- 11 //panic: send on closed channel
		wg.Done()
	}()
}

// 消费者
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			if data, ok := <-ch; ok {
				// true  为未关闭通道 正常接收
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestName(t *testing.T) {
	var wg sync.WaitGroup

	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()

}
