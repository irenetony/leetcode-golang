package main

import (
	"fmt"
	"strconv"
	"strings"
	"sort"
)

type EnvelopArray []Envelop

func (array EnvelopArray) Len() int {
	return len(array)
}

func (array EnvelopArray) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

func (array EnvelopArray) Less(i, j int) bool {
	if array[i].width != array[j].width {
		return array[i].width < array[j].width
	} else {
		return array[i].height < array[j].height
	}
}

type Envelop struct {
	width  int
	height int
}

func (e Envelop) ToString() string {
	return "[" + strconv.Itoa(e.width) + "," + strconv.Itoa(e.height) + "]"
}

func EnvelopesArray2StringArray(envelopes []Envelop) []string {
	res := make([]string, 0)
	for _, val := range envelopes {
		res = append(res, val.ToString())
	}
	return res
}

func MaxEnvelopes(envelopes []Envelop) int {
	sort.Sort(EnvelopArray(envelopes))

	res := 0

	for i := len(envelopes) - 1; i >= 0; i-- {
		curMaxEnvelopes := 1
		curEnvelop := envelopes[i]
		for j := i; j >= 0; j-- {
			if curEnvelop.width > envelopes[j].width && curEnvelop.height > envelopes[j].height {
				curMaxEnvelopes++
				curEnvelop = envelopes[j]
			}
		}
		if curMaxEnvelopes > res {
			res = curMaxEnvelopes
		}
	}

	return res
}

func TestMaxEnvelopes(envelopes []Envelop) {
	fmt.Printf("Given envelopes = [%s], the maximum number of envelopes you can Russian doll is %d\n", strings.Join(EnvelopesArray2StringArray(envelopes), ","), MaxEnvelopes(envelopes))
}

func main() {
	envelopes1 := []Envelop{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	TestMaxEnvelopes(envelopes1)

	envelopes2 := []Envelop{
		{1, 2}, {1, 3}, {1, 4}, {2, 1},
		{2, 2}, {2, 3}, {2, 4}, {2, 5},
		{3, 1}, {3, 2}, {3, 3}, {3, 4},
		{4, 5},
		{6, 1},
	}
	TestMaxEnvelopes(envelopes2)
}
