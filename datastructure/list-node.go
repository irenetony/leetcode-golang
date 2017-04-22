package datastructure

type ListNode struct {
	Val  int
	Next *ListNode
}

type DoubleLinkedListNode struct {
	Val  int
	Pre  *DoubleLinkedListNode
	Next *DoubleLinkedListNode
}
