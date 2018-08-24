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
	if intervals == nil || len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}

	sortInterval(intervals)
	res := make([]MInterval, 0)
	var curInterval MInterval
	for i := 0; i < len(intervals)-1; i++ {
		if i == 0 {
			curInterval = intervals[0]
		}

		if curInterval.End >= intervals[i+1].Start {
			if curInterval.End < intervals[i+1].End {
				curInterval.End = intervals[i+1].End
			}
			if i == len(intervals)-2 {
				res = append(res, curInterval)
			}
		} else {
			res = append(res, curInterval)
			if i == len(intervals)-2 {
				res = append(res, intervals[i+1])
			}
			curInterval = intervals[i+1]
		}
	}

	return res
}

func sortInterval(intervals []MInterval) {
	quickSortInterval(intervals, 0, len(intervals)-1)
}

func quickSortInterval(intervals []MInterval, l, r int) {
	if l >= r {
		return
	}

	pivot := intervals[l].Start
	left := l
	right := r
	for left < right {
		for right > left {
			if intervals[right].Start >= pivot {
				right--
			} else {
				intervals[right], intervals[left] = intervals[left], intervals[right]
				break
			}
		}

		if right == left {
			break
		}

		for right > left {
			if intervals[left].Start <= pivot {
				left++
			} else {
				intervals[right], intervals[left] = intervals[left], intervals[right]
				break
			}
		}
	}

	quickSortInterval(intervals, l, left-1)
	quickSortInterval(intervals, left+1, r)
}

func main() {
	fmt.Printf("%v\n", merge([]MInterval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Printf("%v\n", merge([]MInterval{{1, 4}, {4, 5}}))
	fmt.Printf("%v\n", merge([]MInterval{{1, 4}, {1, 4}}))
	fmt.Printf("%v\n", merge([]MInterval{{1, 4}, {5, 6}}))
	fmt.Printf("%v\n", merge([]MInterval{{2, 3}, {5, 5}, {2, 2}, {3, 4}, {3, 4}}))
	fmt.Printf("%v\n", merge([]MInterval{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}}))
}