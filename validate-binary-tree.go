package main

import (
	"fmt"
	ds "github.com/berryjam/leetcode/datastructure"
)

func IsValidBST(node *ds.TreeNode) bool {
	if node == nil {
		return true
	}

	stack := ds.NewStack()
	stack.Push(node)
	for !stack.IsEmpty() {
		curNode := stack.Pop().(*ds.TreeNode)
		if curNode.Right != nil {
			if curNode.Right.Val <= curNode.Val {
				return false
			} else {
				stack.Push(curNode.Right)
			}
		}
		if curNode.Left != nil {
			if curNode.Left.Val >= curNode.Val {
				return false
			} else {
				stack.Push(curNode.Left)
			}
		}
	}

	return true
}

func TestIsValidBST1() {
	node1 := ds.NewTreeNode(1)
	node2 := ds.NewTreeNode(2)
	node3 := ds.NewTreeNode(3)
	node2.Left = node1
	node2.Right = node3
	fmt.Printf("Tree1 is %t BST\n", IsValidBST(node1))
}

func TestIsValidBST2() {
	node1 := ds.NewTreeNode(1)
	node2 := ds.NewTreeNode(2)
	node3 := ds.NewTreeNode(3)
	node1.Left = node2
	node1.Right = node3
	fmt.Printf("Tree2 is %t BST\n", IsValidBST(node1))
}

func TestIsValidBST3() {
	node1 := ds.NewTreeNode(1)
	node2 := ds.NewTreeNode(2)
	node3 := ds.NewTreeNode(3)
	node4 := ds.NewTreeNode(4)
	node5 := ds.NewTreeNode(5)
	node3.Left = node2
	node2.Left = node1
	node3.Right = node5
	node5.Left = node4
	fmt.Printf("Tree3 is %t BST\n", IsValidBST(node3))
}

func TestIsValidBST4() {
	node1 := ds.NewTreeNode(1)
	node2 := ds.NewTreeNode(2)
	node3 := ds.NewTreeNode(3)
	node4 := ds.NewTreeNode(4)
	node5 := ds.NewTreeNode(5)
	node3.Left = node2
	node2.Left = node1
	node3.Right = node5
	node5.Right = node4
	fmt.Printf("Tree4 is %t BST\n", IsValidBST(node3))
}

func main() {
	TestIsValidBST1()
	TestIsValidBST2()
	TestIsValidBST3()
	TestIsValidBST4()
}
