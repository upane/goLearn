package ch5

import "testing"

func TestIfMultiSec(t *testing.T) {

	if a := 1 == 1; a {
		t.Log("1==1")
	}

	// if var 定义变量：boolean判断条件

}

// go switch
// 默认不需要加break 类似多个if else
func TestCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 1:
			t.Log("111")
		case 2, 3:
			t.Log("222")
		default:
			t.Log("hahha")

		}
	}
}
