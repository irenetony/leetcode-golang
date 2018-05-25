package main

import (
	ds "github.com/berryjam/leetcode-golang/datastructure"
	"fmt"
)

func recoverTree(root *ds.TreeNode) {
	var lastNode, firstNode, backupNode *ds.TreeNode
	stack := ds.NewStack()
	cur := root
	for cur != nil {
		stack.Push(cur)
		cur = cur.Left
	}

	for !stack.IsEmpty() {
		val := stack.Pop()
		cur = val.(*ds.TreeNode)
		if lastNode == nil {
			lastNode = cur
		} else {
			if cur.Val < lastNode.Val {
				if firstNode == nil {
					firstNode = lastNode
					backupNode = cur
				} else {
					cur.Val, firstNode.Val = firstNode.Val, cur.Val
					return
				}
				lastNode = cur
			} else {
				lastNode = cur
			}
		}

		if cur.Right != nil {
			cur = cur.Right
			for cur != nil {
				stack.Push(cur)
				cur = cur.Left
			}
		}
	}

	firstNode.Val, backupNode.Val = backupNode.Val, firstNode.Val
}

func inorderTree(root *ds.TreeNode) {
	stack := ds.NewStack()
	cur := root
	for cur != nil {
		stack.Push(cur)
		cur = cur.Left
	}

	for !stack.IsEmpty() {
		val := stack.Pop()
		cur = val.(*ds.TreeNode)
		fmt.Printf("%d ", cur.Val)
		if cur.Right != nil {
			cur = cur.Right
			for cur != nil {
				stack.Push(cur)
				cur = cur.Left
			}
		}
	}
}

func main() {
	//one := &ds.TreeNode{Val: 1}
	//two := &ds.TreeNode{Val: 2}
	//three := &ds.TreeNode{Val: 3}

	//one.Left = three
	//three.Right = two

	//three.Left = one
	//four := &ds.TreeNode{Val:4}
	//three.Right = four
	//four.Left = two
	//root := three

	//root := three
	//recoverTree(root)
	//fmt.Printf("one=%+v", one)
	//three.Left = one
	//one.Right = two
	//n50 := &ds.TreeNode{Val: 50}
	//n40 := &ds.TreeNode{Val: 40}
	//n60 := &ds.TreeNode{Val: 60}
	//n30 := &ds.TreeNode{Val: 30}
	//n52 := &ds.TreeNode{Val: 52}
	//n70 := &ds.TreeNode{Val: 70}
	//n45 := &ds.TreeNode{Val: 45}
	//n44 := &ds.TreeNode{Val: 44}
	//n47 := &ds.TreeNode{Val: 47}
	//n20 := &ds.TreeNode{Val: 20}
	//n10 := &ds.TreeNode{Val: 10}
	//n21 := &ds.TreeNode{Val: 21}
	//n11 := &ds.TreeNode{Val: 11}
	//n43 := &ds.TreeNode{Val: 43}
	//n46 := &ds.TreeNode{Val: 46}
	//n48 := &ds.TreeNode{Val: 48}
	//n42 := &ds.TreeNode{Val: 42}
	//
	//n50.Right = n60
	//n60.Left = n52
	////n60.Right = n70
	//n60.Right = n40
	////n50.Left = n40
	//n50.Left = n70
	////n40.Left = n30
	////n40.Right = n45
	//n70.Left = n30
	//n70.Right = n45
	//n30.Left = n20
	//n45.Left = n44
	//n45.Right = n47
	//
	//n44.Left = n43
	//n43.Left = n42
	//n47.Left = n46
	//n47.Right = n48
	//n20.Left = n10
	//
	//n20.Right = n21
	//n10.Right = n11
	//inorderTree(n50)
	//fmt.Println()
	//recoverTree(n50)
	//inorderTree(n50)

	n1 := &ds.TreeNode{Val: 1}
	n2 := &ds.TreeNode{Val: 2}
	n3 := &ds.TreeNode{Val: 3}
	n4 := &ds.TreeNode{Val: 4}
	n3.Left = n1
	n3.Right = n4
	n4.Left = n2

	inorderTree(n3)
	fmt.Println()
	recoverTree(n3)
	inorderTree(n3)

}
