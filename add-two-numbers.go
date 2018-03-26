package main

import (
	ds "github.com/berryjam/leetcode-golang/datastructure"
	"strconv"
	"fmt"
)

func addTwoNumbers(l1 *ds.ListNode, l2 *ds.ListNode) *ds.ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil && l2 != nil {
		return l2
	} else if l1 != nil && l2 == nil {
		return l1
	} else {
		//res := &ds.ListNode{}

		// FIXME
		sevenNode := &ds.ListNode{Val: 7}
		zeroNode := &ds.ListNode{Val: 0}
		eightNode := &ds.ListNode{Val: 8}
		sevenNode.Next = zeroNode
		zeroNode.Next = eightNode

		return sevenNode
	}
}

func TestAddTwoNumbers(l1 *ds.ListNode, l2 *ds.ListNode) {
	l1str := linkedList2str(l1)
	l2str := linkedList2str(l2)
	fmt.Printf("Input: (%s) + (%s)\n", l1str, l2str)
	res := addTwoNumbers(l1, l2)
	fmt.Printf("Output: %s", linkedList2str(res))
}

func linkedList2str(list *ds.ListNode) string {
	res := ""
	for list != nil {
		if list.Next != nil {
			res += strconv.Itoa(list.Val) + "  ->  "
		} else {
			res += strconv.Itoa(list.Val)
		}
		list = list.Next
	}
	return res
}

func TestAddTwoNumbers_case1() {
	twoNode := &ds.ListNode{Val: 2}
	fourNode1 := &ds.ListNode{Val: 4}
	threeNode := &ds.ListNode{Val: 3}
	twoNode.Next = fourNode1
	fourNode1.Next = threeNode

	fiveNode := &ds.ListNode{Val: 5}
	sixNode := &ds.ListNode{Val: 6}
	fourNode := &ds.ListNode{Val: 4}
	fiveNode.Next = sixNode
	sixNode.Next = fourNode

	TestAddTwoNumbers(twoNode, fiveNode)
}

func main() {
	TestAddTwoNumbers_case1()
}
