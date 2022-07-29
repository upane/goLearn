package ch20_1

import (
	"fmt"
	"testing"
	"time"
)

func isCanceled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}

}

func cnacel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}
func cnacel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

// 取消任务
func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)

	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCanceled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i, cancelChan)
	}
	cnacel_2(cancelChan)
	time.Sleep(time.Second * 1)
}
