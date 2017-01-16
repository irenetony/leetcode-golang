package main

import "fmt"

type Dicitionary interface {
	AddWord(word string)
	Search(word string) bool
}

type MyDicitionary struct {
}

func (d *MyDicitionary) AddWord(word string) {

}

func (d *MyDicitionary) Search(word string) bool {
	return false
}

func TestSearch() {
	md := &MyDicitionary{}
	var i Dicitionary = md
	word := "bad"
	fmt.Printf("word [%s] exist? %t", word, i.Search(word))
}

func main() {
	TestSearch()
}
