package ch34

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go get github.com/stretchr/testify
// 安装assert
func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expacted := [...]int{1, 4, 9}

	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expacted[i] {
			t.Errorf("%d %d %d", inputs[i], expacted[i], ret)
		}

	}
}

func TestFatalError(t *testing.T) {
	fmt.Println("began")
	t.Error("error") // 可以继续执行
	t.Fatal("error") // 中端操作
	fmt.Println("finish")
}

func TestSquareAsset(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expacted := [...]int{1, 4, 10}

	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(t, expacted[i], ret)

	}
}
