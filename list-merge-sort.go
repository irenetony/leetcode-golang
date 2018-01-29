package main

import (
	ds "leetcode/datastructure"
	"fmt"
	"strconv"
	"strings"
)

func LinkedListMergeSort(list *ds.ListNode) *ds.ListNode {
	if list == nil || list.Next == nil {
		return list
	}

	onePtr := list
	twoPtr := list.Next
	for twoPtr != nil {
		if twoPtr.Next == nil {
			break
		} else {
			onePtr = onePtr.Next
			twoPtr = twoPtr.Next.Next
		}
	}

	left := list
	right := onePtr.Next
	onePtr.Next = nil
	sortedLeft := LinkedListMergeSort(left)
	sortedRight := LinkedListMergeSort(right)
	return mergeKLists([]*ds.ListNode{sortedLeft, sortedRight})
}

func mergeKLists(lists []*ds.ListNode) *ds.ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else {
		var res, cur *ds.ListNode
		listPtr := make(map[*ds.ListNode]bool)
		for _, node := range lists {
			if node != nil {
				listPtr[node] = true
			}
		}
		for len(listPtr) >= 2 {
			var minNode *ds.ListNode = nil
			for node := range listPtr {
				if minNode == nil {
					minNode = node
				} else {
					if minNode.Val > node.Val {
						minNode = node
					}
				}
			}
			if minNode.Next != nil {
				listPtr[minNode.Next] = true
			}
			delete(listPtr, minNode)

			if res == nil {
				res = minNode
				cur = minNode
			} else {
				cur.Next = minNode
				cur = minNode
			}
		}

		for node := range listPtr {
			if res == nil {
				res = node
			} else {
				cur.Next = node
			}
		}

		return res
	}
}

func TestLinkedListMergeSort(list *ds.ListNode) {
	fmt.Println("Input:")
	fmt.Print("[ ")
	for cur := list; cur != nil; cur = cur.Next {
		fmt.Print(strconv.Itoa(cur.Val) + " ")
	}
	fmt.Println("]\n")
	fmt.Println("\nOutput:")
	res := LinkedListMergeSort(list)
	outlist := make([]string, 0)
	for ; res != nil; res = res.Next {
		outlist = append(outlist, strconv.Itoa(res.Val))
	}
	fmt.Println("[ " + strings.Join(outlist, " ") + " ]")
}

func main() {
	node8 := &ds.ListNode{Val: 8}
	node3 := &ds.ListNode{Val: 3}
	node5 := &ds.ListNode{Val: 5}
	node1 := &ds.ListNode{Val: 1}
	node4 := &ds.ListNode{Val: 4}
	node7 := &ds.ListNode{Val: 7}
	node6 := &ds.ListNode{Val: 6}

	node8.Next = node3
	node3.Next = node5
	node5.Next = node1
	node1.Next = node4
	node4.Next = node7
	node7.Next = node6

	TestLinkedListMergeSort(node8)

}
