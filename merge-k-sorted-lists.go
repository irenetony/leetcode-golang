package main

import (
	ds "leetcode/datastructure"
	"fmt"
	"strconv"
	"strings"
)

func mergeKLists(lists []*ds.ListNode) *ds.ListNode {
	// FIXME
	return lists[0]
}

func TestMergeKLists(lists []*ds.ListNode) {
	fmt.Println("Input:")
	for i := 0; i < len(lists); i++ {
		tmp := lists[i]
		fmt.Print("[ ")
		for ; tmp != nil; tmp = tmp.Next {
			fmt.Print(strconv.Itoa(tmp.Val) + " ")
		}
		fmt.Print("]")
		fmt.Println()

	}
	fmt.Println("\nOutput:")
	res := mergeKLists(lists)
	list := make([]string, 0)
	for ; res != nil; res = res.Next {
		list = append(list, strconv.Itoa(res.Val))
	}
	fmt.Println("[ " + strings.Join(list, " ") + " ]")
}

func main() {
	node1 := &ds.ListNode{Val: 1}
	node2 := &ds.ListNode{Val: 2}
	node3 := &ds.ListNode{Val: 3}
	node1.Next = node2
	node2.Next = node3

	node4 := &ds.ListNode{Val: 4}
	node5 := &ds.ListNode{Val: 5}
	node6 := &ds.ListNode{Val: 6}
	node4.Next = node5
	node5.Next = node6

	lists := make([]*ds.ListNode, 0)
	lists = append(lists, node1)
	lists = append(lists, node4)
	TestMergeKLists(lists)
}
