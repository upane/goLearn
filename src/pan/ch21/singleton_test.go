package ch21

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

// 单例模式 懒汉式
type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func GetSingletoneObj() *Singleton {
	once.Do(func() {
		fmt.Println("create obj")
		// new 返回的是指针  所以singleInstance 得是指针
		singleInstance = new(Singleton)
	})
	return singleInstance

}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletoneObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
