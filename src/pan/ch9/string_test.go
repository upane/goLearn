package ch9

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 初始化为”“ 不是null哦
	s = "hello"
	t.Log(len(s), s[1])
	//  不可改变值
	// s[1] = "3" cannot assign to s[1] (value of type byte)
	s = "\xE4\xB8\xA5"
	t.Log(s, len(s)) // string_test.go:13: 严 3
	s = "\xE4\xBA\xB5\xFF"
	t.Log(s)

	s = "中"
	t.Log(len(s))

	c := []rune(s)             // 去除字符串的Unicode 内置机制
	t.Logf("unicode %x", c[0]) // unicode 4e2d
	t.Logf("utf8 %x", s)       // utf8 e4b8ad

}

// String遍历rune
func TestRune(t *testing.T) {
	s := "阿斯蒂芬"
	for _, c := range s {
		t.Logf("%[1]c %[1]d %[1]x", c)
	}
}
