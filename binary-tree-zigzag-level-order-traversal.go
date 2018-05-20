package main

import (
	"fmt"
	ds "github.com/berryjam/leetcode-golang/datastructure"
	"github.com/berryjam/leetcode-golang/util"
	"strings"
)

func ZigzagLevelOrder(root *ds.TreeNode) (res [][]int) {
	if root == nil {
		return
	}

	stack := ds.NewStack()
	queue1 := ds.NewQueue()
	queue2 := ds.NewQueue()
	queue1.Enqueue(root)

	for !queue1.IsEmpty() || !queue2.IsEmpty() {
		if !queue1.IsEmpty() {
			cur := make([]int, 0)
			for !queue1.IsEmpty() {
				val, _ := queue1.Dequeue()
				node := val.(*ds.TreeNode)
				cur = append(cur, node.Val)
				if node.Left != nil {
					queue2.Enqueue(node.Left)
					stack.Push(node.Left)
				}
				if node.Right != nil {
					queue2.Enqueue(node.Right)
					stack.Push(node.Right)
				}
			}
			res = append(res, cur)
		}

		if !stack.IsEmpty() {
			cur := make([]int, 0)
			for !stack.IsEmpty() {
				val := stack.Pop()
				node := val.(*ds.TreeNode)
				cur = append(cur, node.Val)
			}
			res = append(res, cur)
		}

		for !queue2.IsEmpty() {
			val, _ := queue2.Dequeue()
			node := val.(*ds.TreeNode)
			if node.Left != nil {
				queue1.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue1.Enqueue(node.Right)
			}
		}
	}

	return
}

func TestZigzagLevelOrder1() {
	root := &ds.TreeNode{Val: 3}
	nineNode := &ds.TreeNode{Val: 9}
	twentyNode := &ds.TreeNode{Val: 20}
	root.Left = nineNode
	root.Right = twentyNode

	fifteenNode := &ds.TreeNode{Val: 15}
	sevenNode := &ds.TreeNode{Val: 7}
	twentyNode.Left = fifteenNode
	twentyNode.Right = sevenNode

	res := ZigzagLevelOrder(root)
	fmt.Printf("[\n")
	for idx, val := range res {
		tmp := util.IntArray2StringArray(val)
		if idx != len(res)-1 {
			fmt.Printf("\t[%s],\n", strings.Join(tmp, ","))
		} else {
			fmt.Printf("\t[%s]\n", strings.Join(tmp, ","))
		}
	}
	fmt.Printf("]\n")
}

func TestZigzagLevelOrder2() {
	root := &ds.TreeNode{Val: 1}
	twoNode := &ds.TreeNode{Val: 2}
	root.Right = twoNode

	threeNode := &ds.TreeNode{Val: 3}
	fourNode := &ds.TreeNode{Val: 4}
	twoNode.Left = threeNode
	twoNode.Right = fourNode

	fiveNode := &ds.TreeNode{Val: 5}
	sixNode := &ds.TreeNode{Val: 6}
	threeNode.Left = fiveNode
	threeNode.Right = sixNode

	sevenNode := &ds.TreeNode{Val: 7}
	eightNode := &ds.TreeNode{Val: 8}
	fourNode.Left = sevenNode
	fourNode.Right = eightNode

	res := ZigzagLevelOrder(root)
	fmt.Printf("[\n")
	for idx, val := range res {
		tmp := util.IntArray2StringArray(val)
		if idx != len(res)-1 {
			fmt.Printf("\t[%s],\n", strings.Join(tmp, ","))
		} else {
			fmt.Printf("\t[%s]\n", strings.Join(tmp, ","))
		}
	}
	fmt.Printf("]\n")
}

func main() {
	TestZigzagLevelOrder1()
	TestZigzagLevelOrder2()
}
