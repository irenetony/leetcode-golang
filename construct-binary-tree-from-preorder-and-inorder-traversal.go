package main

import (
	ds "github.com/berryjam/leetcode-golang/datastructure"
	"fmt"
)

func buildTree(preorder []int, inorder []int) *ds.TreeNode {
	if inorder == nil || len(inorder) == 0 || preorder == nil || len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 && len(inorder) == 1 {
		return &ds.TreeNode{Val: preorder[0]}
	}

	root := &ds.TreeNode{Val: preorder[0]}

	leftVals := getLeftVals(inorder, root.Val)
	rightVals := getRightVals(inorder, root.Val)
	left := buildTree(preorder[1:], leftVals)
	right := buildTree(preorder[len(leftVals)+1:], rightVals)

	root.Left = left
	root.Right = right

	return root
}

func getLeftVals(inorder []int, root int) []int {
	idx := 0
	for inorder[idx] != root {
		idx++
	}
	if idx == 0 {
		return nil
	}

	return inorder[:idx]
}

func getRightVals(inorder []int, root int) []int {
	idx := 0
	for inorder[idx] != root {
		idx++
	}
	if idx == len(inorder)-1 {
		return nil
	}
	return inorder[idx+1:]
}

func preorders(root *ds.TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = preorders(root.Left, res)
	res = preorders(root.Right, res)
	return res
}

func inorders(root *ds.TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = inorders(root.Left, res)
	res = append(res, root.Val)
	res = inorders(root.Right, res)
	return res
}

func main() {
	//preorder := []int{3, 9, 20, 15, 7}
	//inorder := []int{9, 3, 15, 20, 7}
	//preorder := []int{1, 2}
	//inorder := []int{2, 1}
	preorder := []int{1, 5, 2, 3, 4, 6}
	inorder := []int{1, 2, 3, 4, 5, 6}
	root := buildTree(preorder, inorder)
	fmt.Printf("preorder=%+v\n", preorders(root, []int{}))
	fmt.Printf("inorders=%+v\n", inorders(root, []int{}))
}
