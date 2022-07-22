package ch7

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2}
	t.Log(m1[2])
	t.Log(len(m1))
	// cap 不存在
	// invalid argument: m1 (variable of type map[int]int) for cap

	m2 := map[string]int{}
	m2["hahha"] = 1

	// make指定Map容量初始化     节省空间和性能优化
	m3 := make(map[int]int, 9)
	t.Logf("len =%d", len(m3))
}

func TestAccessNotExistKey(t *testing.T) {
	m1 := map[int]int{}
	m1[1] = 0
	t.Log(m1[0], m1[1])

	m1[3] = 0
	// v为值， ok 判断是否存在
	if v, ok := m1[3]; ok {
		t.Log("key3存在", v)
	} else {
		t.Log("not exist")
	}
}

// 循环
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 2, 3: 4, 5: 6}

	for k, v := range m1 {
		t.Log(k, v)
	}
}
