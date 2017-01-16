package main

import (
	"fmt"
	"strings"
	"strconv"
)

func ReadBinaryWatch(num int) []string {
	var res []string

	if num == 0 {
		res = append(res, "0:00")
		return []string{"0:00"}
	}

	for hour := 0; hour <= num; hour++ {
		minute := num - hour
		hourTimes := HourTime(hour)
		minuteTimes := MinuteTime(minute)
		for _, hourTime := range hourTimes {
			for _, minuteTime := range minuteTimes {
				res = append(res, strconv.Itoa(hourTime)+":"+MinuteTimeFormat(minuteTime))
			}
		}
	}

	return res
}

func MinuteTimeFormat(num int) string {
	if num < 10 {
		return "0" + strconv.Itoa(num)
	} else {
		return strconv.Itoa(num)
	}
}

func HourTime(num int) []int {
	return Combination([]int{8, 4, 2, 1}, num)
}

func MinuteTime(num int) [] int {
	return Combination([]int{32, 16, 8, 4, 2, 1}, num)
}

func Combination(array []int, count int) []int {
	if count == 0 {
		return []int{0}
	}

	var res []int

	if count == 1 {
		return array
	}

	for i := 0; i < len(array); i++ {
		nextRes := Combination(DeleteElement(array, i), count-1)
		for _, val := range nextRes {
			res = append(res, array[i]+val)
		}
	}

	return res
}

func DeleteElement(array []int, idx int) []int {
	var res []int
	for i, val := range array {
		if i != idx {
			res = append(res, val)
		}
	}
	return res
}

func TestReadBinaryWatch(n int) {
	fmt.Printf("Input: n = %d\n", n)

	var res []string = ReadBinaryWatch(n)

	for idx, val := range res {
		res[idx] = "\"" + val + "\""
	}

	fmt.Printf("Return: [ %s ]\n", strings.Join(res, ","))
}

func main() {
	TestReadBinaryWatch(1)
	TestReadBinaryWatch(2)
}
