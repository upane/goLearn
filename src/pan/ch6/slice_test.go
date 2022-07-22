package ch6

import "testing"

// 切片为共享存储空间
// 数组容量为固定长度，切片可以变化
// 数组相同长度可比较，切片不可以

func TestSliceInit(t *testing.T) {
	// 声明切片
	var s0 []int
	t.Log(len(s0), cap(s0))

	//append(s0, 1) (value of type []int) is not used 需要被赋值
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	// make 创建  类型，len，cap
	s2 := make([]int, 2, 5)
	t.Log(len(s2), cap(s2))

	s2 = append(s2, 444)
	// len 初始化个数 cap容量  访问大于len 会报错
	// panic: runtime error: index out of range [3] with length 3
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

// 切片的扩容
func TestSliceGrowing(t *testing.T) {
	s := []int{1, 2, 3}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}

	// 存储空间的复制
}

func TestShareMod(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	Q2 := numbers[2:3]
	t.Log(len(Q2), cap(Q2))
	//slice_test.go:42: 1 7
	// 3, 4, 5, 6, 7, 8, 9 后面的cap 为共享空间 所以为7
	t.Log(len(numbers), cap(numbers))

	// 修改Q2  numbers也会发生改变
	Q2[0] = 999
	t.Log(numbers)
}
