package main

import (
	ds "leetcode/datastructure"
	"fmt"
)

func diameterOfBinaryTree(root *ds.TreeNode) int {
	if root == nil || root.Left == nil && root.Right == nil {
		return 0
	} else {
		leftDiameter := 0
		leftHeight := 0
		crossDiameter := 0
		if root.Left != nil {
			leftDiameter = diameterOfBinaryTree(root.Left)
			leftHeight = heightOfBinaryTree(root.Left)
			crossDiameter++
		}
		rightDiameter := 0
		rightHeight := 0
		if root.Right != nil {
			rightDiameter = diameterOfBinaryTree(root.Right)
			rightHeight = heightOfBinaryTree(root.Right)
			crossDiameter++
		}
		crossDiameter += leftHeight + rightHeight
		return max(max(leftDiameter, rightDiameter), crossDiameter)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func heightOfBinaryTree(root *ds.TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 0
	} else {
		leftHeight := 0
		if root.Left != nil {
			leftHeight = heightOfBinaryTree(root.Left)
		}
		rightHeight := 0
		if root.Right != nil {
			rightHeight = heightOfBinaryTree(root.Right)
		}
		if leftHeight > rightHeight {
			return leftHeight + 1
		} else {
			return rightHeight + 1
		}
	}
}

func main() {
	oneNode := &ds.TreeNode{Val: 1}
	twoNode := &ds.TreeNode{Val: 2}
	threeNode := &ds.TreeNode{Val: 3}
	fourNode := &ds.TreeNode{Val: 4}
	fiveNode := &ds.TreeNode{Val: 5}

	oneNode.Left = twoNode
	oneNode.Right = threeNode
	twoNode.Left = fourNode
	twoNode.Right = fiveNode

	res := diameterOfBinaryTree(oneNode)
	fmt.Printf("%d", res)
}
