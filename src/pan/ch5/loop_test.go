package ch5

import "testing"

func TestWithLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}
