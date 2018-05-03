package main

import (
	ds "github.com/berryjam/leetcode-golang/datastructure"
	"fmt"
)

func kthSmallest(root *ds.TreeNode, k int) int {
	leftCount := treeCount(root.Left)
	if leftCount+1 == k {
		return root.Val
	} else if leftCount >= k {
		return kthSmallest(root.Left, k)
	} else {
		return kthSmallest(root.Right, k-leftCount-1)
	}
}

func treeCount(root *ds.TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else {
		left := treeCount(root.Left)
		right := treeCount(root.Right)
		return 1 + left + right
	}
}

func main() {
	n3 := &ds.TreeNode{Val: 3}
	n2 := &ds.TreeNode{Val: 2}
	n1 := &ds.TreeNode{Val: 1}
	n5 := &ds.TreeNode{Val: 5}
	n6 := &ds.TreeNode{Val: 6}
	n8 := &ds.TreeNode{Val: 8}

	n3.Left = n1
	n1.Right = n2

	n3.Right = n6
	n6.Left = n5
	n6.Right = n8
	fmt.Printf("res=%+v\n", kthSmallest(n3, 1))
}
