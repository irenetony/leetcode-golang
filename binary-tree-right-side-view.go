package main

import (
	ds "leetcode/datastructure"
	"fmt"
	"leetcode/util"
	"strings"
)

func RightSideView(root *ds.TreeNode) (res []int) {
	queue1 := ds.NewQueue()
	queue2 := ds.NewQueue()
	queue1.Enqueue(root)

	for !queue1.IsEmpty() || !queue2.IsEmpty() {
		for !queue1.IsEmpty() {
			head, err := queue1.Dequeue()
			if err != nil {
				panic(err)
			}
			node := head.(*ds.TreeNode)
			if node.Left != nil {
				queue2.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue2.Enqueue(node.Right)
			}
			if queue1.Len() == 0 {
				res = append(res, node.Val)
			}
		}

		for !queue2.IsEmpty() {
			head, err := queue2.Dequeue()
			if err != nil {
				panic(err)
			}
			node := head.(*ds.TreeNode)
			if node.Left != nil {
				queue1.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue1.Enqueue(node.Right)
			}
			if queue2.Len() == 0 {
				res = append(res, node.Val)
			}
		}
	}
	return
}

func TestRightSideView1() {
	root := &ds.TreeNode{Val:1}
	twoNode := &ds.TreeNode{Val:2}
	threeNode := &ds.TreeNode{Val:3}
	root.Left = twoNode
	root.Right = threeNode

	fourNode := &ds.TreeNode{Val:4}
	fiveNode := &ds.TreeNode{Val:5}
	twoNode.Right = fiveNode
	threeNode.Right = fourNode

	res := RightSideView(root)
	fmt.Printf("[%s]\n", strings.Join(util.IntArray2StringArray(res), ","))
}

func TestRightSideView2() {
	root := &ds.TreeNode{Val:1}
	twoNode := &ds.TreeNode{Val:2}
	threeNode := &ds.TreeNode{Val:3}
	root.Left = twoNode
	root.Right = threeNode

	fourNode := &ds.TreeNode{Val:4}
	fiveNode := &ds.TreeNode{Val:5}
	twoNode.Right = fiveNode
	threeNode.Right = fourNode

	sixNode := &ds.TreeNode{Val:6}
	fiveNode.Left = sixNode

	res := RightSideView(root)
	fmt.Printf("[%s]\n", strings.Join(util.IntArray2StringArray(res), ","))
}

func main() {
	TestRightSideView1()
	TestRightSideView2()
}
