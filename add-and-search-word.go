package main

import "fmt"

var MY_DICITIONARY *MyDicitionary = &MyDicitionary{}

type Dicitionary interface {
	AddWord(word string)
	Search(word string) bool
}

type MyDicitionary struct {
	rootNodes map[rune]*Node
}

type Node struct {
	ChildNodes map[rune]*Node
	IsLeafNode bool
}

func (d *MyDicitionary) AddWord(word string) {
	if d.rootNodes == nil {
		d.rootNodes = make(map[rune]*Node)
	}
	curChildNodes := d.rootNodes
	var parentNode *Node

	for i, char := range word {
		if curChildNodes != nil {
			if _, ok := curChildNodes[char]; ok {
				parentNode = curChildNodes[char]
				curChildNodes = curChildNodes[char].ChildNodes
			} else {
				if i == len(word)-1 {
					curChildNodes[char] = &Node{IsLeafNode:true}
					curChildNodes = curChildNodes[char].ChildNodes
				} else {
					curChildNodes[char] = &Node{IsLeafNode:false}
					parentNode = curChildNodes[char]
					curChildNodes = curChildNodes[char].ChildNodes
				}
			}
		} else {
			curChildNodes = make(map[rune]*Node)
			if i == len(word)-1 {
				curChildNodes[char] = &Node{IsLeafNode:true}
			} else {
				curChildNodes[char] = &Node{IsLeafNode:false}
			}
			parentNode.ChildNodes = curChildNodes
			parentNode = curChildNodes[char]
			curChildNodes = curChildNodes[char].ChildNodes
		}
	}
}

func (d *MyDicitionary) Search(word string) bool {
	curChildNodes := d.rootNodes
	for i, char := range word {
		if curChildNodes == nil {
			return false
		} else {
			if char == '.' {
				if i == len(word)-1 {
					return true
				} else {
					for _, v := range curChildNodes {
						if exists := SearchStarWord(word[i + 1:], v.ChildNodes); exists {
							return true
						}
					}
					return false
				}
			} else {
				if _, ok := curChildNodes[char]; !ok {
					return false
				} else {
					if i == len(word)-1 {
						return true
					} else {
						curChildNodes = curChildNodes[char].ChildNodes
					}
				}
			}
		}
	}

	return false
}

func SearchStarWord(word string, childNodes map[rune]*Node) bool {
	for i, char := range word {
		if childNodes == nil {
			return false
		} else {
			if char == '.' {
				if i == len(word)-1 {
					return true
				} else {
					for _, v := range childNodes {
						if exists := SearchStarWord(word[i + 1:], v.ChildNodes); exists {
							return true
						}
					}
					return false
				}
			} else {
				if _, ok := childNodes[char]; !ok {
					return false
				} else {
					if i == len(word)-1 {
						return true
					} else {
						childNodes = childNodes[char].ChildNodes
					}
				}
			}
		}
	}
	return false
}

func TestCase1() {
	var i Dicitionary = MY_DICITIONARY
	i.AddWord("bad")
	i.AddWord("dad")
	i.AddWord("mad")
	fmt.Printf("word [%s] exist? %t\n", "pad", i.Search("pad"))
	fmt.Printf("word [%s] exist? %t\n", "bad", i.Search("bad"))
	fmt.Printf("word [%s] exist? %t\n", ".ad", i.Search(".ad"))
	fmt.Printf("word [%s] exist? %t\n", "b..", i.Search("b.."))
	fmt.Printf("word [%s] exist? %t\n", "mad", i.Search("mad"))
	fmt.Printf("word [%s] exist? %t\n", "badd", i.Search("badd"))
	fmt.Printf("word [%s] exist? %t\n", "...", i.Search("..."))
	fmt.Printf("word [%s] exist? %t\n", "b.e", i.Search("b.e"))
}

func main() {
	TestCase1()
}
