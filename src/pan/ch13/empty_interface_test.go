package ch13

import (
	"fmt"
	"testing"
)

// go 倾向于定义小接口 大接口，为使用小接口组合的方法
// 只依赖表功能的小接口 增加复用性
// 空接口
func DoSomething(p interface{}) {
	//if _, ok := p.(int); ok {
	//	fmt.Println("Integer")
	//	return
	//}
	//if _, ok := p.(string); ok {
	//	fmt.Println("String")
	//	return
	//}
	//fmt.Println("UnKnow")

	//定义了值必须得用到 不然后报错
	//v declared but not used
	//switch v := p.(type) {
	switch p.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("SSS")

	}
}

// 哦哦哦 Test必须以Test开头
func TestName(t *testing.T) {
	DoSomething(1)

}
