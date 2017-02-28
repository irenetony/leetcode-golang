package main

import (
	"fmt"
	"sort"
)

type IntArray []int

func (array IntArray) Len() int {
	return len(array)
}

func (array IntArray) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

func (array IntArray) Less(i, j int) bool {
	return array[i] < array[j]
}

type Interval struct {
	start int
	end   int
}

func (i *Interval) ToString() string {
	return fmt.Sprintf("[%d, %d]", i.start, i.end)
}

type SummaryRanges struct {
	elements []int
}

func (s *SummaryRanges) AddNum(val int) {
	s.elements = append(s.elements, val)
	sort.Sort(IntArray(s.elements))
}

func (s *SummaryRanges) GetIntervals() []Interval {
	if len(s.elements) == 0 {
		return nil
	} else if len(s.elements) == 1 {
		return []Interval{{s.elements[0], s.elements[0]}}
	} else {
		res := make([]Interval, 0)
		start, end := s.elements[0], -1
		isContinous := true
		for i := 0; i < len(s.elements)-1; i++ {
			if s.elements[i]+1 != s.elements[i+1] {
				if isContinous {
					end = s.elements[i]
					res = append(res, Interval{start, end})
					isContinous = false
				} else {
					res = append(res, Interval{s.elements[i], s.elements[i]})
				}
				start = s.elements[i+1]
			} else {
				if !isContinous {
					isContinous = true
				}
			}
		}
		if isContinous {
			end = s.elements[len(s.elements)-1]
			res = append(res, Interval{start, end})
		} else {
			res = append(res, Interval{s.elements[len(s.elements)-1], s.elements[len(s.elements)-1]})
		}
		return res
	}
}

func TestSummaryRanges1() {
	obj := &SummaryRanges{}
	TestSummaryRanges(obj, 1)
	TestSummaryRanges(obj, 3)
	TestSummaryRanges(obj, 7)
	TestSummaryRanges(obj, 2)
	TestSummaryRanges(obj, 6)
}

func TestSummaryRanges2() {
	obj := &SummaryRanges{}
	TestSummaryRanges(obj, 1)
	TestSummaryRanges(obj, 3)
	TestSummaryRanges(obj, 7)
	TestSummaryRanges(obj, 2)
	TestSummaryRanges(obj, 6)
	TestSummaryRanges(obj, 9)
	TestSummaryRanges(obj, 11)
}

func TestSummaryRanges(obj *SummaryRanges, val int) {
	obj.AddNum(val)
	intervals := obj.GetIntervals()
	res := ""
	for i := 0; i < len(intervals); i++ {
		if i != len(intervals)-1 {
			res += intervals[i].ToString() + ", "
		} else {
			res += intervals[i].ToString()
		}
	}
	fmt.Printf("%s\n", res)
}

func main() {
	TestSummaryRanges1()
	fmt.Println()
	TestSummaryRanges2()
}
