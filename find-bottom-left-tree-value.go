package main

import (
	ds "github.com/berryjam/leetcode/datastructure"
	"fmt"
)

func findBottomLeftValue(root *ds.TreeNode) int {
	res := root.Val
	queue1 := ds.NewQueue()
	queue2 := ds.NewQueue()
	queue1.Enqueue(root)

	for !queue1.IsEmpty() || !queue2.IsEmpty() {
		isLastRow := true

		findLeftMostValQueue(queue1, &res, &isLastRow, queue2)
		findLeftMostValQueue(queue2, &res, &isLastRow, queue1)

		if isLastRow {
			return res
		}
	}

	return res
}

func findLeftMostValQueue(queue *ds.Queue, res *int, isLastRow *bool, nextQueue *ds.Queue) {
	if !queue.IsEmpty() {
		resRoot, _ := queue.Dequeue()
		resNode := resRoot.(*ds.TreeNode)
		*res = resNode.Val
		if resNode.Left != nil {
			*isLastRow = false
			nextQueue.Enqueue(resNode.Left)
		}
		if resNode.Right != nil {
			*isLastRow = false
			nextQueue.Enqueue(resNode.Right)
		}
		for !queue.IsEmpty() {
			curRoot, _ := queue.Dequeue()
			curNode := curRoot.(*ds.TreeNode)
			if curNode.Left != nil {
				*isLastRow = false
				nextQueue.Enqueue(curNode.Left)
			}
			if curNode.Right != nil {
				*isLastRow = false
				nextQueue.Enqueue(curNode.Right)
			}
		}
	}
}

func TestFindBottomLeftValue(root *ds.TreeNode) {
	res := findBottomLeftValue(root)
	fmt.Printf("Output:\n%d\n", res)
}

func main() {
	root1 := ds.NewTreeNode(2)
	one1 := ds.NewTreeNode(1)
	three1 := ds.NewTreeNode(3)
	root1.Left = one1
	root1.Right = three1
	TestFindBottomLeftValue(root1)

	root2 := ds.NewTreeNode(1)
	two2 := ds.NewTreeNode(2)
	three2 := ds.NewTreeNode(3)
	four2 := ds.NewTreeNode(4)
	five2 := ds.NewTreeNode(5)
	six2 := ds.NewTreeNode(6)
	seven2 := ds.NewTreeNode(7)
	eight2 := ds.NewTreeNode(8)
	seven2.Right = eight2
	root2.Left = two2
	root2.Right = three2
	two2.Left = four2
	three2.Left = five2
	three2.Right = six2
	five2.Left = seven2
	TestFindBottomLeftValue(root2)
}
