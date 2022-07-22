package ch6

import "testing"

// 数组赋值
func TestArrayInit(t *testing.T) {
	//var arr [3]int  声明初始化默认值为零

	// 初始化并赋值，需要赋值 ...可不指定长度
	arr1 := [...]int{123, 12, 124}
	// index 2 is out of bounds (>= 2)  长度<指定长度

	// 多维数组赋值
	arr2 := [2][3]int{{123, 12, 124}, {123, 12, 124}}

	t.Log(arr1[0], arr1[1])

	// 数组遍历  1、for语句 2、range
	for i := 0; i < len(arr1); i++ {
		t.Log(arr1[i])
	}

	// for _, num := range arr1 {  _可以忽略值
	// idx 为索引，num为值
	for idx, num := range arr2 {
		t.Log("outline", idx, num)
		for i, d := range num {
			t.Log("inline", i, d)
		}
	}

}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	// 不支持负数倒数截取
	arr3_sec := arr3[1:]
	t.Log(arr3_sec)
}
