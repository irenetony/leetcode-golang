package main

import (
	ds "github.com/berryjam/leetcode/datastructure"
	"fmt"
)

func partition(head *ds.ListNode, x int) *ds.ListNode {
	if head == nil {
		return nil
	}

	var lessHeadNode, lessTailNode *ds.ListNode
	var greatHeadNode, greatTailNode *ds.ListNode

	for ; head != nil; head = head.Next {
		if head.Val < x {
			if lessHeadNode == nil {
				lessHeadNode = head
				lessTailNode = head
			} else {
				lessTailNode.Next = head
				lessTailNode = head
			}
		} else {
			if greatHeadNode == nil {
				greatHeadNode = head
				greatTailNode = head
			} else {
				greatTailNode.Next = head
				greatTailNode = head
			}
		}
	}

	if lessTailNode != nil {
		lessTailNode.Next = greatHeadNode
	}

	if greatTailNode != nil {
		greatTailNode.Next = nil
	}

	if lessHeadNode != nil {
		return lessHeadNode
	}

	return greatHeadNode
}

func main() {
	oneNode := &ds.ListNode{Val: 1}
	fourNode := &ds.ListNode{Val: 4}
	threeNode := &ds.ListNode{Val: 3}
	twoNode := &ds.ListNode{Val: 2}
	fiveNode := &ds.ListNode{Val: 5}
	twotwoNode := &ds.ListNode{Val: 2}

	oneNode.Next = fourNode
	fourNode.Next = threeNode
	threeNode.Next = twoNode
	twoNode.Next = fiveNode
	fiveNode.Next = twotwoNode

	res := partition(oneNode, 3)
	fmt.Println(res.Val)
}
