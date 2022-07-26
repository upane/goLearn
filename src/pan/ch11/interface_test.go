package ch11

import "testing"

type Programmer interface {
	WriteHello() string
}
type GoProgrammer struct {
}

func WriteHello() string {
	return "fmt.Println(\"hello\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	t.Log(p.WriteHello())

}
