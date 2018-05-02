package main

import (
	ds "github.com/berryjam/leetcode-golang/datastructure"
	"fmt"
)

func buildTree(preorder []int, inorder []int) *ds.TreeNode {
	if len(preorder) == 1 {
		return &ds.TreeNode{Val: preorder[0]}
	}

	if len(inorder) == 1 {
		return &ds.TreeNode{Val: inorder[0]}
	}

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	rootNode := &ds.TreeNode{Val: rootVal}
	leftPre, leftIn, rightPre, rightIn := getPart(rootVal, preorder, inorder)

	leftNode := buildTree(leftPre, leftIn)
	rightNode := buildTree(rightPre, rightIn)

	rootNode.Left = leftNode
	rootNode.Right = rightNode

	return rootNode
}

func getPart(root int, preorder, inorder []int) (leftPre, leftIn, rightPre, rightIn []int) {
	var leftInIdx int
	for i, v := range inorder {
		if v != root {
			leftInIdx = i
		} else {
			break
		}
	}
	if (leftInIdx + 2) < len(inorder) {
		rightIn = inorder[leftInIdx+2:]
	}
	if root != inorder[0] {
		leftIn = inorder[:leftInIdx+1]
	} else {
		if root != inorder[leftInIdx+1] {
			rightIn = inorder[leftInIdx+1:]
		}
	}

	var leftPreIdx int
	for i, v := range preorder[1:] {
		if inSlice(v, rightIn) {
			leftPreIdx = i + 1
			break
		}
	}
	if leftPreIdx >= 1 {
		leftPre = preorder[1:leftPreIdx]
	}
	rightPre = preorder[leftPreIdx:]

	return
}

func inSlice(val int, slice []int) bool {
	res := false

	for _, v := range slice {
		if val == v {
			return true
		}
	}

	return res
}

func main() {
	//preorder := []int{3, 9, 20, 15, 7}
	//inorder := []int{9, 3, 15, 20, 7}
	preorder := []int{1, 2}
	inorder := []int{2, 1}
	//preorder := []int{1, 5, 2, 3, 4, 6}
	//inorder := []int{1, 2, 3, 4, 5, 6}
	root := buildTree(preorder, inorder)
	fmt.Printf("tree=%+v", root)
}
