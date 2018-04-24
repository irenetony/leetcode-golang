package main

import (
	ds "github.com/berryjam/leetcode-golang/datastructure"
	"fmt"
)

func reverseBetween(head *ds.ListNode, m int, n int) *ds.ListNode {
	res := head
	tmp := head
	var mNode, premNode, nNode, postnNode *ds.ListNode
	count := 1
	for tmp != nil {
		if m == 1 {
			mNode = head
		}
		if count == m-1 {
			premNode = tmp
			mNode = premNode.Next
		}
		if count == n {
			nNode = tmp
			postnNode = nNode.Next
		}
		tmp = tmp.Next
		count++
	}

	newHead, newTail := reverseLinkedList(mNode, nNode)
	if premNode != nil {
		premNode.Next = newHead
	} else {
		res = newHead
	}
	newTail.Next = postnNode
	return res
}

func reverseLinkedList(head *ds.ListNode, tail *ds.ListNode) (newHead, newTail *ds.ListNode) {
	newTail = head
	var pre *ds.ListNode
	cur := head
	for {
		next := cur.Next
		cur.Next = pre
		pre = cur
		if cur == tail {
			break
		}
		cur = next
	}
	newHead = pre
	return
}

func main() {
	n1 := &ds.ListNode{Val: 1}
	n2 := &ds.ListNode{Val: 2}
	n3 := &ds.ListNode{Val: 3}
	n4 := &ds.ListNode{Val: 4}
	n5 := &ds.ListNode{Val: 5}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	res := reverseBetween(n1, 1, 5)
	for res != nil {
		if res.Next != nil {
			fmt.Printf("%d -> ", res.Val)
		} else {
			fmt.Printf("%d", res.Val)
		}
		res = res.Next
	}
}
