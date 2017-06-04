package main

import (
	"strings"
	"strconv"
	"sort"
	"fmt"
)

type StringArray []string

func (array StringArray) Len() int {
	return len(array)
}

func (array StringArray) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

func (array StringArray) Less(i, j int) bool {
	iArray := strings.Split(array[i], ":")
	iHour, _ := strconv.Atoi(iArray[0])
	iMinute, _ := strconv.Atoi(iArray[1])

	jArray := strings.Split(array[j], ":")
	jHour, _ := strconv.Atoi(jArray[0])
	jMinute, _ := strconv.Atoi(jArray[1])

	if iHour < jHour {
		return true
	} else if iHour == jHour {
		if iMinute < jMinute {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func findMinDifference(timePoints []string) int {
	res := 24 * 60
	sort.Sort(StringArray(timePoints))
	for i := 0; i < len(timePoints); i++ {
		leftTimeDiff := getMinuteDifference(timePoints[i], timePoints[(i-1+len(timePoints))%len(timePoints)])
		rightTimeDiff := getMinuteDifference(timePoints[i], timePoints[(i+1)%len(timePoints)])
		if leftTimeDiff < rightTimeDiff {
			if leftTimeDiff < res {
				res = leftTimeDiff
			}
		} else {
			if rightTimeDiff < res {
				res = rightTimeDiff
			}
		}
	}

	return res
}

func getMinuteDifference(time1, time2 string) int {
	halfDay := 12 * 60
	diff := halfDay

	array1 := strings.Split(time1, ":")
	hour1, _ := strconv.Atoi(array1[0])
	minute1, _ := strconv.Atoi(array1[1])

	array2 := strings.Split(time2, ":")
	hour2, _ := strconv.Atoi(array2[0])
	minute2, _ := strconv.Atoi(array2[1])

	if hour1 < hour2 {
		diff = hour2*60 + minute2 - hour1*60 - minute1
		if diff > halfDay {
			diff = (24+hour1)*60 + minute1 - hour2*60 - minute2
		}
	} else {
		if hour1 == hour2 {
			if minute1 > minute2 {
				diff = minute1 - minute2
			} else {
				diff = minute2 - minute1
			}
		} else {
			diff = hour1*60 + minute1 - hour2*60 - minute2
			if diff > halfDay {
				diff = (24+hour2)*60 + minute2 - hour1*60 - minute1
			}
		}
	}

	return diff
}

func TestFindMinDifference(timePoints []string) {
	fmt.Printf("Input: [%s]\n", strings.Join(timePoints, ","))
	fmt.Printf("Output: %d\n", findMinDifference(timePoints))
}

func main() {
	//timePoints1 := []string{"23:59", "11:58"}
	//TestFindMinDifference(timePoints1)

	timePoints2 := []string{"01:01", "02:01", "03:00", "03:01"}
	TestFindMinDifference(timePoints2)
}
