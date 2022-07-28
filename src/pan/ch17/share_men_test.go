package ch17

import (
	"sync"
	"testing"
	"time"
)

//共享内存  类似于java 锁。
func TestName(t *testing.T) {
	mut := sync.Mutex{}
	//var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			// 加锁解锁 线程安全
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++

		}()
	}
	time.Sleep(1 * time.Second)
	t.Log(counter)
}

//共享内存  类似于java 锁。
func TestNameWaitGroup(t *testing.T) {
	mut := sync.Mutex{}
	// 读锁不是互斥锁 所以可以用
	var wg sync.WaitGroup
	//var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1) //写2出现报错死锁：all goroutines are asleep - deadlock!
		go func() {
			// 加锁解锁 线程安全
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	//time.Sleep(1 * time.Second)
	t.Log(counter)
}
