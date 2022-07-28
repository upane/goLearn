package ch14

import (
	"errors"
	"fmt"
	"testing"
)

// panic 用于不可恢复错误  会输出调用栈信息
// defer   在panic会被执行 在os.Exit不会执行
func TestPanic(t *testing.T) {
	defer func() {
		// recover xixixixi SomeThing wrong
		// 是恶魔 尽量谨慎
		if err := recover(); err != nil {
			fmt.Println("xixixixi", err)
		}
	}()
	fmt.Println("hahha")
	//os.Exit(222)
	panic(errors.New("SomeThing wrong"))
}
