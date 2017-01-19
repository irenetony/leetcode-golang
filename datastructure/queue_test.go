package datastructure

import (
	"fmt"
	"testing"
)

func Test_Queue1(t *testing.T) {
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	fmt.Printf("queue is empty:%t\n", queue.IsEmpty())
	a, err := queue.Dequeue()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("a = %d\n", a.(int))
	b, err := queue.Dequeue()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("b = %d\n", b.(int))
}

func Test_Queue2(t *testing.T) {
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	fmt.Printf("queue is empty:%t\n", queue.IsEmpty())
	a, err := queue.Dequeue()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("a = %d\n", a.(int))
	b, err := queue.Dequeue()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("b = %d\n", b.(int))
	c, err := queue.Dequeue()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("c = %d\n", c.(int))
}
