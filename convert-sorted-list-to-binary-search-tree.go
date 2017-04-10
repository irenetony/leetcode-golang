package main

import (
	ds "leetcode/datastructure"
	"fmt"
)

func sortedListToBST(head *ds.ListNode) *ds.TreeNode {
	array := make([]int, 0)
	for head != nil {
		array = append(array, head.Val)
		head = head.Next
	}

	if len(array) == 0 {
		return nil
	}

	res := getTreeNode(array, 0, len(array)-1)

	return res
}

func getTreeNode(array []int, left, right int) *ds.TreeNode {
	if left <= right {
		mid := (left + right) / 2
		root := &ds.TreeNode{Val: array[mid]}

		leftNode := getTreeNode(array, left, mid-1)
		if leftNode != nil {
			root.Left = leftNode
		}

		rightNode := getTreeNode(array, mid+1, right)
		if rightNode != nil {
			root.Right = rightNode
		}

		return root
	} else {
		return nil
	}
}

func main() {
	one := &ds.ListNode{Val: 1}
	three := &ds.ListNode{Val: 3}
	seven := &ds.ListNode{Val: 7}
	eight := &ds.ListNode{Val: 8}
	ten := &ds.ListNode{Val: 10}
	twelve := &ds.ListNode{Val: 12}
	fifteen := &ds.ListNode{Val: 15}
	eighteen := &ds.ListNode{Val: 18}
	twenty := &ds.ListNode{Val: 20}
	twentyOne := &ds.ListNode{Val: 21}

	one.Next = three
	three.Next = seven
	seven.Next = eight
	eight.Next = ten
	ten.Next = twelve
	twelve.Next = fifteen
	fifteen.Next = eighteen
	eighteen.Next = twenty
	twenty.Next = twentyOne

	res := sortedListToBST(nil)
	fmt.Printf("root : %d", res.Val)
}
