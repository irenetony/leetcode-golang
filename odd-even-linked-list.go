package main

import (
	"fmt"
	ds "github.com/berryjam/leetcode/datastructure"
)

func OddEvenList(head *ds.ListNode) *ds.ListNode {
	if head == nil {
		return nil
	}

	var oddPtr *ds.ListNode = head
	var evenPtr *ds.ListNode = oddPtr.Next

	var firstEvenPtr *ds.ListNode = evenPtr

	var nextOddPtr, nextEvenPtr, lastOddPtr *ds.ListNode

	lastOddPtr = oddPtr
	for oddPtr != nil || evenPtr != nil {
		if evenPtr != nil {
			nextOddPtr = evenPtr.Next
			oddPtr.Next = nextOddPtr
			if nextOddPtr != nil {
				lastOddPtr = nextOddPtr
				nextEvenPtr = nextOddPtr.Next
				evenPtr.Next = nextEvenPtr
				evenPtr = nextEvenPtr
			} else {
				evenPtr = nil
			}
			oddPtr = nextOddPtr
		} else {
			break
		}
	}

	lastOddPtr.Next = firstEvenPtr

	return head
}

func TestOddEvenList0() {
	var l1 *ds.ListNode = nil
	res := OddEvenList(l1)
	for res != nil {
		fmt.Printf("%d->", res.Val)
		res = res.Next
	}
	fmt.Print("NULL")
}

func TestOddEvenList1() {
	l1 := &ds.ListNode{
		Val: 1,
	}
	res := OddEvenList(l1)
	for res != nil {
		fmt.Printf("%d->", res.Val)
		res = res.Next
	}
	fmt.Print("NULL")
}

func TestOddEvenList2() {
	l1 := &ds.ListNode{
		Val: 1,
	}
	l2 := &ds.ListNode{
		Val: 2,
	}
	l1.Next = l2

	res := OddEvenList(l1)
	for res != nil {
		fmt.Printf("%d->", res.Val)
		res = res.Next
	}
	fmt.Print("NULL")
}

func TestOddEvenList3() {
	l1 := &ds.ListNode{
		Val: 1,
	}
	l2 := &ds.ListNode{
		Val: 2,
	}
	l3 := &ds.ListNode{
		Val: 3,
	}
	l1.Next = l2
	l2.Next = l3

	res := OddEvenList(l1)
	for res != nil {
		fmt.Printf("%d->", res.Val)
		res = res.Next
	}
	fmt.Print("NULL")
}

func TestOddEvenList4() {
	l1 := &ds.ListNode{
		Val: 1,
	}
	l2 := &ds.ListNode{
		Val: 2,
	}
	l3 := &ds.ListNode{
		Val: 3,
	}
	l4 := &ds.ListNode{
		Val: 4,
	}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4

	res := OddEvenList(l1)
	for res != nil {
		fmt.Printf("%d->", res.Val)
		res = res.Next
	}
	fmt.Print("NULL")
}

func TestOddEvenList5() {
	l1 := &ds.ListNode{
		Val: 1,
	}
	l2 := &ds.ListNode{
		Val: 2,
	}
	l3 := &ds.ListNode{
		Val: 3,
	}
	l4 := &ds.ListNode{
		Val: 4,
	}
	l5 := &ds.ListNode{
		Val: 5,
	}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5

	res := OddEvenList(l1)
	for res != nil {
		fmt.Printf("%d->", res.Val)
		res = res.Next
	}
	fmt.Print("NULL")
}

func TestOddEvenList6() {
	l1 := &ds.ListNode{
		Val: 1,
	}
	l2 := &ds.ListNode{
		Val: 2,
	}
	l3 := &ds.ListNode{
		Val: 3,
	}
	l4 := &ds.ListNode{
		Val: 4,
	}
	l5 := &ds.ListNode{
		Val: 5,
	}
	l6 := &ds.ListNode{
		Val: 6,
	}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6

	res := OddEvenList(l1)
	for res != nil {
		fmt.Printf("%d->", res.Val)
		res = res.Next
	}
	fmt.Print("NULL")
}

func main() {
	//TestOddEvenList0()
	//TestOddEvenList1()
	//TestOddEvenList2()
	//TestOddEvenList3()
	//TestOddEvenList4()
	//TestOddEvenList5()
	//TestOddEvenList6()
}
