package main

import (
	ds "github.com/berryjam/leetcode-golang/datastructure"
	"fmt"
	"strconv"
	"strings"
)

var dArr []string
var idx = -1

func serialize(root *ds.TreeNode) string {
	res := ""
	if root == nil {
		return "#,"
	}

	res += strconv.Itoa(root.Val) + ","
	res += serialize(root.Left)
	res += serialize(root.Right)

	return res
}

func deserialize(data string) *ds.TreeNode {
	if len(dArr) == 0 {
		dArr = strings.Split(data, ",")
	}

	idx++
	if dArr[idx] == "#" {
		return nil
	}

	val, err := strconv.Atoi(dArr[idx])
	if err != nil {
		panic(err)
	}
	node := &ds.TreeNode{Val: val}
	left := deserialize(data)
	right := deserialize(data)

	node.Left = left
	node.Right = right

	return node
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

	str := serialize(n3)
	fmt.Printf("tree=%s\n", str[:len(str)-1])
	node := deserialize(str)
	fmt.Printf("tree=%+v\n", node)
}
