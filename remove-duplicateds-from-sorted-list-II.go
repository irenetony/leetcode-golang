package main

import (
	"fmt"
	"strings"
	"github.com/berryjam/leetcode-golang/util"
	ds "github.com/berryjam/leetcode-golang/datastructure"
)

func deleteDuplicated(head *ds.ListNode) *ds.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var res *ds.ListNode
	cur := head.Next
	pre := head
	var preDiff *ds.ListNode
	if cur.Val != pre.Val {
		preDiff = pre
		res = preDiff
	}
	for cur != nil {
		if cur.Val != pre.Val {
			if preDiff == nil {
				if cur.Next == nil || cur.Next != nil && cur.Next.Val != cur.Val {
					preDiff = cur
					res = preDiff
				}
			} else {
				if cur.Next == nil || cur.Next != nil && cur.Next.Val != cur.Val {
					preDiff.Next = cur
					preDiff = cur
				}
			}
		} else {
			if cur.Next == nil && preDiff != nil {
				preDiff.Next = nil
			}
		}
		pre = cur
		cur = cur.Next
	}

	return res
}

func TestDeleteDuplicated1() {
	firstNode := &ds.ListNode{Val: 1}
	secondNode := &ds.ListNode{Val: 2}
	thirdNode := &ds.ListNode{Val: 3}
	fourthNode := &ds.ListNode{Val: 3}
	fifthNode := &ds.ListNode{Val: 4}
	sixthNode := &ds.ListNode{Val: 4}
	seventhNode := &ds.ListNode{Val: 5}

	firstNode.Next = secondNode
	secondNode.Next = thirdNode
	thirdNode.Next = fourthNode
	fourthNode.Next = fifthNode
	fifthNode.Next = sixthNode
	sixthNode.Next = seventhNode

	TestDeleteDuplicated(firstNode)
}

func TestDeleteDuplicated2() {
	firstNode := &ds.ListNode{Val: 1}
	secondNode := &ds.ListNode{Val: 1}
	thirdNode := &ds.ListNode{Val: 1}
	fourthNode := &ds.ListNode{Val: 2}
	fifthNode := &ds.ListNode{Val: 3}

	firstNode.Next = secondNode
	secondNode.Next = thirdNode
	thirdNode.Next = fourthNode
	fourthNode.Next = fifthNode

	TestDeleteDuplicated(firstNode)
}

func TestDeleteDuplicated3() {
	firstNode := &ds.ListNode{Val: 1}
	secondNode := &ds.ListNode{Val: 2}
	thirdNode := &ds.ListNode{Val: 2}

	firstNode.Next = secondNode
	secondNode.Next = thirdNode

	TestDeleteDuplicated(firstNode)
}

func TestDeleteDuplicated4() {
	firstNode := &ds.ListNode{Val: 1}
	secondNode := &ds.ListNode{Val: 1}

	firstNode.Next = secondNode

	TestDeleteDuplicated(firstNode)
}

func TestDeleteDuplicated(head *ds.ListNode) {
	if head == nil {
		return
	}
	var inputArray []int
	curNode := head
	for curNode != nil {
		inputArray = append(inputArray, curNode.Val)
		curNode = curNode.Next
	}
	fmt.Printf("Input: %s \n", strings.Join(util.IntArray2StringArray(inputArray), " -> "))
	resHead := deleteDuplicated(head)
	curNode = resHead
	var outputArray []int
	for curNode != nil {
		outputArray = append(outputArray, curNode.Val)
		curNode = curNode.Next
	}
	fmt.Printf("Output: %s \n", strings.Join(util.IntArray2StringArray(outputArray), " -> "))
}

func main() {
	//TestDeleteDuplicated1()
	//TestDeleteDuplicated2()
	//TestDeleteDuplicated3()
	TestDeleteDuplicated4()
}
