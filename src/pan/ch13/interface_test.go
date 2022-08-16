package ch13

import (
	"fmt"
	"testing"
)

// 定义别名  其实code代表 string
// 可类比返回复杂type实体类
type Code string

// 定义接口 重写方法
type Programmer interface {
	WriteHello() Code
}
type GoProgrammer struct {
}

// 定义接口
func (p *GoProgrammer) WriteHello() Code {
	return "fmt.Println(\"hello\")"
}

type JavaProgrammer struct {
}

// 定义接口
func (p *JavaProgrammer) WriteHello() Code {
	return "System.out.Println(\"hello\")"
}

func writeFirstProgrammer(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHello())
	// %T 输出传入符号的类型
	// interface最好为指针

}

// 多态实现
func TestClient(t *testing.T) {
	//gopro := new(GoProgrammer)  最好这种吧 maybe  new(GoProgrammer)
	gopro := &GoProgrammer{}
	//fmt.Println(gopro.WriteHello())
	writeFirstProgrammer(gopro)

	jvaPro := new(JavaProgrammer)
	//fmt.Println(jvaPro.WriteHello())
	writeFirstProgrammer(jvaPro)

}
