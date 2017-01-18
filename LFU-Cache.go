package main

import "fmt"

type LFUCache struct {
	capacity      int
	cacheMap      map[int]int
	frequencyList FrequencyList
}

type FrequencyList struct {
	head *FrequencyNode
	tail *FrequencyNode
}

type FrequencyNode struct {
	val  int
	next *FrequencyNode
}

func (cache *LFUCache) Get(key int) int {
	if val, ok := cache.cacheMap[key]; ok {
		// update data structure
		var preNode *FrequencyNode = nil
		for node := cache.frequencyList.head; node != nil; {
			if node.val == key {
				if preNode != nil {
					preNode.next = node.next
				} else {
					cache.frequencyList.head = cache.frequencyList.head.next
				}
				cache.frequencyList.tail.next = node
				cache.frequencyList.tail = node
				node.next = nil
				break
			} else {
				preNode = node
				node = node.next
			}
		}
		return val
	} else {
		return -1
	}
}

func (cache *LFUCache) Put(key, value int) {
	if len(cache.cacheMap) < cache.capacity {
		cache.cacheMap[key] = value
		newTailNode := &FrequencyNode{val:key}
		if cache.frequencyList.tail == nil {
			cache.frequencyList.tail = newTailNode
		} else {
			cache.frequencyList.tail.next = newTailNode
			cache.frequencyList.tail = newTailNode
		}

		if cache.frequencyList.head == nil {
			cache.frequencyList.head = newTailNode
		}
	} else {
		// evicts least recently used key and update node
		delete(cache.cacheMap, cache.frequencyList.head.val)
		cache.frequencyList.head = cache.frequencyList.head.next
		cache.cacheMap[key] = value
	}
}

func TestCache1() {
	cache := &LFUCache{
		capacity:      2,
		cacheMap:      make(map[int]int),
		frequencyList: FrequencyList{},
	}
	cache.Put(1, 1)
	fmt.Printf("cache.put(1,1);\n")
	cache.Put(2, 2)
	fmt.Printf("cache.put(2,2);\n")
	fmt.Printf("cache.get(1);       // returns %d\n", cache.Get(1))
	cache.Put(3, 3)
	fmt.Printf("cache.put(3, 3);    // evicts key 2\n")
	fmt.Printf("cache.get(2);       // returns %d. (not found)\n", cache.Get(2))
	fmt.Printf("cache.get(3);       // returns %d\n.", cache.Get(3))
	cache.Put(4, 4)
	fmt.Printf("cache.put(4, 4);    // evicts key 1.\n")
	fmt.Printf("cache.get(1);       // returns %d (not found)\n", cache.Get(1))
	fmt.Printf("cache.get(3);       // returns %d\n", cache.Get(3))
	fmt.Printf("cache.get(4);       // returns %d\n", cache.Get(4))
}

func TestCache2() {
	cache := &LFUCache{
		capacity:      3,
		cacheMap:      make(map[int]int),
		frequencyList: FrequencyList{},
	}
	cache.Put(1, 1)
	fmt.Printf("cache.put(1,1);\n")
	cache.Put(2, 2)
	fmt.Printf("cache.put(2,2);\n")
	cache.Put(3, 3)
	fmt.Printf("cache.put(3,3);\n")
	fmt.Printf("cache.get(1);       // returns %d\n", cache.Get(1))
	cache.Put(5, 5)
	fmt.Printf("cache.put(5,5);\n")
	fmt.Printf("cache.get(2);       // returns %d\n", cache.Get(2))
}

func main() {
	//TestCache1()
	TestCache2()
}
