package ch8

import "testing"

// map的工厂模式
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

// map设置set  无内置set
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Log("1111")
	} else {
		t.Log("22222")
	}
	t.Log(len(mySet))

	// 删除key
	delete(mySet, 1)
	t.Log(len(mySet))
}
