// demo_15.go
package main

import (
	"fmt"
)

func main() {

	var a = 1
	var b = 2

	defer fmt.Println(a + b)

	a = 2

	fmt.Println("main")
}
