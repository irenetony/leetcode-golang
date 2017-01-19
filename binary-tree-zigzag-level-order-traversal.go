package main

import (
	"fmt"
	ds "leetcode/datastructure"
	"strings"
	"leetcode/util"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func ZigzagLevelOrder(root *TreeNode) (res [][]int) {
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
			currArray = append(currArray, curNode.(*TreeNode).val)

			if idx%2 == 0 { // 自左向右把子节点放到stack
				leftNode := curNode.(*TreeNode).left
				if leftNode != nil {
					stack.Push(leftNode)
				}
				rightNode := curNode.(*TreeNode).right
				if rightNode != nil {
					stack.Push(rightNode)
				}
			} else { // 自右向左把自节点放到stack
				rightNode := curNode.(*TreeNode).right
				if rightNode != nil {
					stack.Push(rightNode)
				}
				leftNode := curNode.(*TreeNode).left
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
	root := &TreeNode{val:3}
	nineNode := &TreeNode{val:9}
	twentyNode := &TreeNode{val:20}
	root.left = nineNode
	root.right = twentyNode

	fifteenNode := &TreeNode{val:15}
	sevenNode := &TreeNode{val:7}
	twentyNode.left = fifteenNode
	twentyNode.right = sevenNode

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
	root := &TreeNode{val:1}
	twoNode := &TreeNode{val:2}
	root.right = twoNode

	threeNode := &TreeNode{val:3}
	fourNode := &TreeNode{val:4}
	twoNode.left = threeNode
	twoNode.right = fourNode

	fiveNode := &TreeNode{val:5}
	sixNode := &TreeNode{val:6}
	threeNode.left = fiveNode
	threeNode.right = sixNode

	sevenNode := &TreeNode{val:7}
	eightNode := &TreeNode{val:8}
	fourNode.left = sevenNode
	fourNode.right = eightNode

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
