package main

import (
	"fmt"
	ds "github.com/berryjam/leetcode/datastructure"
	"github.com/berryjam/leetcode/util"
	"strings"
)

func ZigzagLevelOrder(root *ds.TreeNode) (res [][]int) {
	res = [][]int{}
	if root == nil {
		return
	}

	stack := ds.NewStack()
	queue := ds.NewQueue()
	queue.Enqueue(root)
	idx := 0
	for !queue.IsEmpty() || !stack.IsEmpty() {

		currArray := []int{}
		for !queue.IsEmpty() {
			curNode, err := queue.Dequeue()
			if err != nil {
				panic(err)
			}
			currArray = append(currArray, curNode.(*ds.TreeNode).Val)

			if idx%2 == 0 { // 自左向右把子节点放到stack
				leftNode := curNode.(*ds.TreeNode).Left
				if leftNode != nil {
					stack.Push(leftNode)
				}
				rightNode := curNode.(*ds.TreeNode).Right
				if rightNode != nil {
					stack.Push(rightNode)
				}
			} else { // 自右向左把自节点放到stack
				rightNode := curNode.(*ds.TreeNode).Right
				if rightNode != nil {
					stack.Push(rightNode)
				}
				leftNode := curNode.(*ds.TreeNode).Left
				if leftNode != nil {
					stack.Push(leftNode)
				}
			}
		}
		if len(currArray) > 0 {
			res = append(res, currArray)
		}

		for !stack.IsEmpty() {
			curNode := stack.Pop()
			queue.Enqueue(curNode)
		}

		idx++
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
