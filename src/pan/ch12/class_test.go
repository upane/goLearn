package ch12

import (
	"fmt"
	"testing"
)

// 复合
type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("petSpeak")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Print("" + host)
}

type Dog struct {
	p *Pet
	//Pet // 匿名定义  可以默认有点类似继承  // 指针引用可以为复合
}

// 自己定义会重载
func (d *Dog) Speak() {
	d.p.Speak()
	fmt.Print("dogSpeak")
}

func (d *Dog) SpeakTo(host string) {
	d.p.SpeakTo(host)
}

// go 不支持继承  不能LSP转换 强制转换类型
// 也就是子类不能动态实现父类的方法 只能子类实现父类单向
func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("lilei") // 调用的为pet   复合 需要重新定义

}
