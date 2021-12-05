package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	as := InitArrayStack(10)
	for i := 0; i < 10; i++ {
		ret := as.Push(fmt.Sprintf("%s, %d", "Hello", i))
		if ret {
			t.Log("ArrayStack Push Success")
		} else {
			t.Log("ArrayStack Push Failure")
		}
	}

	for i := 0; i < 10; i++ {
		t.Log(as.Pop())
	}
}
