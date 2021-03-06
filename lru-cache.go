package main

import "fmt"

type DoubleLinkedlistNode struct {
	Key  int
	Val  int
	Pre  *DoubleLinkedlistNode
	Next *DoubleLinkedlistNode
}

type LRUCache struct {
	Cap   int
	Cache map[int]*DoubleLinkedlistNode
	Head  *DoubleLinkedlistNode
	Tail  *DoubleLinkedlistNode
}

func Constructor1(capacity int) LRUCache {
	return LRUCache{Cap: capacity, Cache: make(map[int]*DoubleLinkedlistNode)}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Cache[key]; ok {
		pre := node.Pre
		next := node.Next
		if pre != nil {
			pre.Next = next
		}
		if next != nil {
			if node != this.Head {
				next.Pre = pre
			}
		}
		if node == this.Tail {
			if pre != nil {
				this.Tail = pre
			} else {
				this.Tail = node
			}
		}
		this.Head.Pre = node
		node.Next = this.Head
		node.Pre = nil
		this.Head = node
		return node.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Cache[key]; ok {
		node.Key = key
		node.Val = value
		this.Get(key)
	} else {
		node = &DoubleLinkedlistNode{Key: key, Val: value}
		if len(this.Cache) < this.Cap {
			if this.Head == nil {
				this.Head = node
			} else {
				this.Head.Pre = node
				node.Next = this.Head
				this.Head = node
			}
			if this.Tail == nil {
				this.Tail = node
			}
		} else {
			this.Head.Pre = node
			node.Next = this.Head
			lruKey := this.Tail.Key
			pre := this.Tail.Pre
			if pre != nil {
				pre.Next = nil
				this.Tail = pre
			}

			this.Head = node
			delete(this.Cache, lruKey)
		}
		this.Cache[key] = node
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	//cache := Constructor1(2)
	//cache.Put(1, 1)
	//cache.Put(2, 2)
	//fmt.Printf("get(1)=%d\n", cache.Get(1))
	//cache.Put(3, 3)
	//fmt.Printf("get(2)=%d\n", cache.Get(2))
	//cache.Put(4, 4)
	//fmt.Printf("get(1)=%d\n", cache.Get(1))
	//fmt.Printf("get(3)=%d\n", cache.Get(3))
	//fmt.Printf("get(4)=%d\n", cache.Get(4))

	//cache := Constructor1(1)
	//cache.Put(2, 1)
	//fmt.Printf("get(2)=%d\n", cache.Get(2))
	//cache.Put(3, 2)
	//fmt.Printf("get(2)=%d\n", cache.Get(2))
	//fmt.Printf("get(3)=%d\n", cache.Get(3))

	cache := Constructor1(2)
	cache.Put(2, 1)
	cache.Put(3, 2)
	fmt.Printf("get(3)=%d\n", cache.Get(3))
	fmt.Printf("get(2)=%d\n", cache.Get(2))
	cache.Put(4, 3)
	fmt.Printf("get(1)=%d\n", cache.Get(2))
	fmt.Printf("get(3)=%d\n", cache.Get(3))
	fmt.Printf("get(4)=%d\n", cache.Get(4))
}
