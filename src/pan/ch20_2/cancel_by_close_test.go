package ch20_1

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// Context  与取消任务  取消上下文任务
// Deadline
func isCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}

}

// 取消任务
func TestCancel(t *testing.T) {

	// 去的上下文 cancel方法 ctx 可传入子任务
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCanceled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
