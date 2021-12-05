package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {

	n := 10
	aq := InitArrayQueue(n)

	for i := 0; i < n-5; i++ {
		if !aq.Enqueue(fmt.Sprintf("%s, %d", "hello", i)) {
			t.Log("queue is full")
		}
	}

	for i := 0; i < n-5; i++ {
		t.Log(aq.Dequeue())
	}

	for i := 0; i < n; i++ {
		if !aq.Enqueue(fmt.Sprintf("%s, %d", "hello", i)) {
			t.Logf("queue is full. hello %d\n", i)
		}
	}

	for i := 0; i < n; i++ {
		t.Log(aq.Dequeue())
	}
}
