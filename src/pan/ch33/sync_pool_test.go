package ch33

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// sync.pool 受gc影响 不建议创建池  缓存对象涉及锁啥的
func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new object.")
			return 100
		},
	}
	// 取出值
	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	runtime.GC() // 没被gc  wtf  是更新golang了吗
	v1, _ := pool.Get().(int)
	fmt.Println(v1)

}
