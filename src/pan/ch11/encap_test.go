package ch11

import (
	"fmt"
	"testing"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreate(t *testing.T) {
	e := Employee{"1", "xiaomi", 11}
	e1 := Employee{Name: "mike", Age: 8}
	e2 := new(Employee) //返回指针
	e2.Id = "0"
	e2.Name = "xixi"
	e2.Age = 3

	t.Log(e)
	t.Log(e1)
	t.Log(e2)
	t.Logf("%T", e) //打印数据类型
	t.Logf("%T", e2)
}

// 与结构绑定的方法会复制对象
func (e Employee) String() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

// 指针指向同一个属性 推荐这个哦哦哦
func (e *Employee) String22() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestName(t *testing.T) {
	e := Employee{"0", "12", 12}
	// 方法可已绑定到结构体  类似java impl ？
	t.Log(e.String22())
}
