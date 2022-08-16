package ch39

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

// 别名 type myInt  int   转 int 可以转换  //强制转换会错误不可取
// atomic 原子操作

// 把读写到新的buffer 然后重新指向 需要unsafe.Pointer
// 共享buffer 的安全读写

// 追求更高效的读写
func TestAtomic(t *testing.T) {
	var shareBuffer unsafe.Pointer
	writeDataFuncs := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBuffer, unsafe.Pointer(&data))
	}

	readDataFn := func() {
		data := atomic.LoadPointer(&shareBuffer)
		fmt.Println(data, *(*[]int)(shareBuffer))
	}
	var wg sync.WaitGroup
	writeDataFuncs()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFuncs()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readDataFn()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()

	}

}
