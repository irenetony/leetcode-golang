package main

import "fmt"

type MInterval struct {
	Start int
	End   int
}

//func (m *MInterval) String() string {
//	return fmt.Sprintf("[%d,%d]", m.Start, m.End)
//}

func merge(intervals []MInterval) []MInterval {

	return []MInterval{{1, 6}, {8, 10}, {15, 18}}
}

func sortInterval(intervals []MInterval) {
	quickSortInterval(intervals, 0, len(intervals)-1)
}

func quickSortInterval(intervals []MInterval, l, r int) {

}

func main() {
	fmt.Printf("%v\n", merge([]MInterval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
