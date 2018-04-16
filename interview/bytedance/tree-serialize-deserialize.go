package main

import ds "github.com/berryjam/leetcode-golang/datastructure"

func serialize(root *ds.TreeNode) string {

}

func deserialize(data string) *ds.TreeNode {

}

func main() {
	n3 := &ds.TreeNode{Val: 3}
	n9 := &ds.TreeNode{Val: 9}
	n20 := &ds.TreeNode{Val: 20}
	n15 := &ds.TreeNode{Val: 15}
	n7 := &ds.TreeNode{Val: 7}

	n3.Left = n9
	n3.Right = n20
	n20.Left = n15
	n20.Right = n7

	
}
