package main

import (
	ds "github.com/berryjam/leetcode-golang/datastructure"
	"fmt"
	"strconv"
	"strings"
	"math"
)

func mergeKLists(lists []*ds.ListNode) *ds.ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	eachListCurNodes := make([]*ds.ListNode, 0)
	for i := 0; i < len(lists); i++ {
		eachListCurNodes = append(eachListCurNodes, lists[i])
	}

	var resListHeadNode, resListTailNode *ds.ListNode = nil, nil
	notNullListIdx := -1

	for hasTwoOrMoreToBeMergedLists(eachListCurNodes, &notNullListIdx) {
		min := math.MaxInt64
		minIdx := -1
		for i := 0; i < len(eachListCurNodes); i++ {
			if eachListCurNodes[i] != nil && eachListCurNodes[i].Val < min {
				min = eachListCurNodes[i].Val
				minIdx = i
			}
		}
		nextNode := eachListCurNodes[minIdx].Next
		if resListHeadNode == nil {
			resListHeadNode = eachListCurNodes[minIdx]
			resListTailNode = eachListCurNodes[minIdx]
		} else {
			resListTailNode.Next = eachListCurNodes[minIdx]
			resListTailNode = eachListCurNodes[minIdx]
		}
		eachListCurNodes[minIdx].Next = nil
		eachListCurNodes[minIdx] = nextNode
	}

	if resListTailNode != nil {
		resListTailNode.Next = eachListCurNodes[notNullListIdx]
	} else {
		if notNullListIdx != -1 {
			return lists[notNullListIdx]
		}
	}

	return resListHeadNode
}

func hasTwoOrMoreToBeMergedLists(lists []*ds.ListNode, notNullListIdx *int) bool {
	count := 0
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			count++
			*notNullListIdx = i
		}
	}
	if count >= 2 {
		return true
	} else {
		return false
	}
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
	node9 := &ds.ListNode{Val: 9}
	node4.Next = node5
	node5.Next = node6
	node6.Next = node9

	node7 := &ds.ListNode{Val: 7}
	node10 := &ds.ListNode{Val: 10}
	node11 := &ds.ListNode{Val: 11}
	node12 := &ds.ListNode{Val: 12}
	node13 := &ds.ListNode{Val: 13}
	node7.Next = node10
	node10.Next = node11
	node11.Next = node12
	node12.Next = node13

	lists := make([]*ds.ListNode, 0)
	//lists = append(lists, node1)
	//lists = append(lists, node4)
	//lists = append(lists, node7)
	lists = append(lists, nil)
	lists = append(lists, &ds.ListNode{Val: 1})
	TestMergeKLists(lists)
}
