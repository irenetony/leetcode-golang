package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	if strs == nil || len(strs) == 0 {
		return nil
	}

	res := make([][]string, 0)
	rCountsMapArr := make([]map[rune]int, 0)
	for _, str := range strs {
		matchExistedCountsMap := false
		matchIdx := -1
		curRCountsMap := make(map[rune]int)
		for _, r := range str {
			if _, ok := curRCountsMap[r]; !ok {
				curRCountsMap[r] = 1
			} else {
				curRCountsMap[r]++
			}
		}
		for i, rCountsMap := range rCountsMapArr {
			twoMapsMatch := true
			if len(curRCountsMap) == len(rCountsMap) {
				for r, count := range curRCountsMap {
					if _, ok := rCountsMap[r]; !ok || count != rCountsMap[r] {
						twoMapsMatch = false
						break
					}
				}
				if twoMapsMatch {
					matchExistedCountsMap = true
					matchIdx = i
					break
				}
			}
		}
		if !matchExistedCountsMap {
			res = append(res, []string{str})
			rCountsMapArr = append(rCountsMapArr, curRCountsMap)
		} else {
			res[matchIdx] = append(res[matchIdx], str)
		}
	}

	return res
}

func main() {
	fmt.Printf("res=%+v\n", groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
