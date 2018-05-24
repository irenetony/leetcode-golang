package main

import "fmt"

type Node struct {
	isWord bool
	ch     rune
	childs map[rune][]*Node
}

type WordDictionary struct {
	root *Node
}

/** Initialize your data structure here. **/
func Constructor() WordDictionary {
	var dict WordDictionary
	dict.root = new(Node)
	dict.root.childs = make(map[rune][]*Node)
	return dict
}

/** Adds a word into the data structure. **/
func (this *WordDictionary) AddWord(word string) {
	curNode := this.root
	for i, r := range word {
		if childs, ok := curNode.childs[r]; ok {
			for _, child := range childs {
				if child.ch == r {
					curNode = child
					break
				}
			}
		} else {
			node := &Node{ch: r, childs: make(map[rune][]*Node)}
			if i == len(word)-1 {
				node.isWord = true
			}
			curNode.childs[r] = append(curNode.childs[r], node)
			curNode = node
		}
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.
' to represent any one letter. **/
func (this *WordDictionary) Search(word string) bool {
	cur := this.root
	for i, r := range word {
		if r != '.' {
			if childs, ok := cur.childs[r]; !ok {
				return false
			} else {
				for _, child := range childs {
					if child.ch == r {
						if i == len(word)-1 && child.isWord {
							return true
						}
						cur = child
						break
					}
				}
			}
		} else {
			if i == len(word)-1 {
				return true
			}
			for _, nodes := range cur.childs {
				for _, node := range nodes {
					dict := Constructor()
					dict.root = node
					if dict.Search(word[i+1:]) {
						return true
					}
				}
			}
			return false
		}
	}

	return false
}

func TestCase1() {
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")
	fmt.Printf("word [%s] exist? %t\n", "pad", obj.Search("pad"))
	fmt.Printf("word [%s] exist? %t\n", "bad", obj.Search("bad"))
	fmt.Printf("word [%s] exist? %t\n", ".ad", obj.Search(".ad"))
	fmt.Printf("word [%s] exist? %t\n", "b..", obj.Search("b.."))
	fmt.Printf("word [%s] exist? %t\n", "mad", obj.Search("mad"))
	fmt.Printf("word [%s] exist? %t\n", "badd", obj.Search("badd"))
	fmt.Printf("word [%s] exist? %t\n", "...", obj.Search("..."))
	fmt.Printf("word [%s] exist? %t\n", "b.e", obj.Search("b.e"))
	fmt.Printf("word [%s] exist? %t\n", "ma.", obj.Search("ma."))
}

func main() {
	TestCase1()
}
package main

import "fmt"
type Node struct {
	isWord bool
	ch     rune
	childs map[rune][]*Node
}

type WordDictionary struct {
	root *Node
}

/** Initialize your data structure here. **/
func Constructor() WordDictionary {
	var dict WordDictionary
	dict.root = new(Node)
	dict.root.childs = make(map[rune][]*Node)
	return dict
}

/** Adds a word into the data structure. **/
func (this *WordDictionary) AddWord(word string) {
	curNode := this.root
	for i, r := range word {
		if childs, ok := curNode.childs[r]; ok {
			for _, child := range childs {
				if child.ch == r {
					curNode = child
					break
				}
			}
		} else {
			node := &Node{ch: r, childs: make(map[rune][]*Node)}
			if i == len(word)-1 {
				node.isWord = true
			}
			curNode.childs[r] = append(curNode.childs[r], node)
			curNode = node
		}
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.
' to represent any one letter. **/
func (this *WordDictionary) Search(word string) bool {
	cur := this.root
	for i, r := range word {
		if r != '.' {
			if childs, ok := cur.childs[r]; !ok {
				return false
			} else {
				for _, child := range childs {
					if child.ch == r {
						if i == len(word)-1 && child.isWord {
							return true
						}
						cur = child
						break
					}
				}
			}
		} else {
			if i == len(word)-1 {
				for _, nodes := range cur.childs {
					for _, node := range nodes {
						if node.isWord {
							return true
						}
					}
				}
				return false
			}
			for _, nodes := range cur.childs {
				for _, node := range nodes {
					dict := Constructor()
					dict.root = node
					if dict.Search(word[i+1:]) {
						return true
					}
				}
			}
			return false
		}
	}

	return false
}

func TestCase1() {
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")
	fmt.Printf("word [%s] exist? %t\n", "pad", obj.Search("pad"))
	fmt.Printf("word [%s] exist? %t\n", "bad", obj.Search("bad"))
	fmt.Printf("word [%s] exist? %t\n", ".ad", obj.Search(".ad"))
	fmt.Printf("word [%s] exist? %t\n", "b..", obj.Search("b.."))
	fmt.Printf("word [%s] exist? %t\n", "mad", obj.Search("mad"))
	fmt.Printf("word [%s] exist? %t\n", "badd", obj.Search("badd"))
	fmt.Printf("word [%s] exist? %t\n", "...", obj.Search("..."))
	fmt.Printf("word [%s] exist? %t\n", "b.e", obj.Search("b.e"))
	fmt.Printf("word [%s] exist? %t\n", "ma.", obj.Search("ma."))
}

func main() {
	TestCase1()
}
