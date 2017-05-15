package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type RandomListNodeSolution struct {
	head *ListNode
}

/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func RandomListNodeConstructor(head *ListNode) RandomListNodeSolution {
	return RandomListNodeSolution{head: head}
}

/** Returns a random node's value. */
func (this *RandomListNodeSolution) GetRandom() int {
	res := 0
	p := this.head
	for cnt := 0; p != nil; p = p.Next {
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		if random.Intn(cnt+1) == 0 {
			res = p.Val
		}
		cnt++
	}

	return res
}

func main() {
	// Init a singly linked list [1,2,3].
	head := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	head.Next = node2
	node2.Next = node3
	solution := RandomListNodeConstructor(head)

	// getRandom() should return either 1, 2, or 3 randomly. Each element should have equal probability of returning.
	fmt.Printf("val:%d\n", solution.GetRandom())
}
