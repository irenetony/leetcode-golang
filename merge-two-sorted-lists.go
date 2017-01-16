package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil && l2 != nil {
		return l2
	} else if l1 != nil && l2 == nil {
		return l1
	} else {
		var res *ListNode
		var last *ListNode

		p1 := l1
		p2 := l2

		for p1 != nil && p2 != nil {
			var tmp *ListNode
			if p1.Val <= p2.Val {
				tmp = &ListNode{
					Val: p1.Val,
				}
				p1 = p1.Next
			} else {
				tmp = &ListNode{
					Val:p2.Val,
				}
				p2 = p2.Next
			}
			if last != nil {
				last.Next = tmp
				last = tmp
			} else {
				res = tmp
				last = tmp
			}
		}

		for p1 != nil {
			tmp := &ListNode{
				Val :p1.Val,
			}
			last.Next = tmp
			last = tmp
			p1 = p1.Next
		}

		for p2 != nil {
			tmp := &ListNode{
				Val :p2.Val,
			}
			last.Next = tmp
			last = tmp
			p2 = p2.Next
		}

		return res
	}
}

func TestMergeTwoLists1() {
	l1 := &ListNode{
		Val:1,
	}
	l3 := &ListNode{
		Val:3,
	}
	l4 := &ListNode{
		Val:4,
	}
	l6 := &ListNode{
		Val:6,
	}
	l7 := &ListNode{
		Val:7,
	}

	l2 := &ListNode{
		Val:2,
	}
	l8 := &ListNode{
		Val:8,
	}
	l9 := &ListNode{
		Val:9,
	}
	l10 := &ListNode{
		Val:10,
	}

	l1.Next = l3
	l3.Next = l4
	l4.Next = l6
	l6.Next = l7

	l2.Next = l8
	l8.Next = l9
	l9.Next = l10

	res := MergeTwoLists(l1, l2)

	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
}

func TestMergeTwoLists2() {
	l1 := &ListNode{
		Val:1,
	}
	l3 := &ListNode{
		Val:3,
	}
	l4 := &ListNode{
		Val:4,
	}
	l6 := &ListNode{
		Val:6,
	}
	l7 := &ListNode{
		Val:7,
	}
	l11 := &ListNode{
		Val:11,
	}
	l12 := &ListNode{
		Val:12,
	}
	l13 := &ListNode{
		Val:13,
	}

	l2 := &ListNode{
		Val:2,
	}
	l8 := &ListNode{
		Val:8,
	}
	l9 := &ListNode{
		Val:9,
	}
	l10 := &ListNode{
		Val:10,
	}

	l1.Next = l3
	l3.Next = l4
	l4.Next = l6
	l6.Next = l7
	l7.Next = l11
	l11.Next = l12
	l12.Next = l13

	l2.Next = l8
	l8.Next = l9
	l9.Next = l10

	res := MergeTwoLists(l1, l2)

	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
}

func TestMergeTwoLists3() {
	var l1 *ListNode = nil

	l2 := &ListNode{
		Val:2,
	}
	l8 := &ListNode{
		Val:8,
	}
	l9 := &ListNode{
		Val:9,
	}
	l10 := &ListNode{
		Val:10,
	}

	l2.Next = l8
	l8.Next = l9
	l9.Next = l10

	res := MergeTwoLists(l1, l2)

	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
}

func main() {
	//TestMergeTwoLists2()
	TestMergeTwoLists3()
}
