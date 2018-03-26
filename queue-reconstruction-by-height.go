package main

import (
        "fmt"
        "github.com/berryjam/leetcode/util"
        "strings"
)

func reconstructQueue(people [][]int) [][]int {
        // FIXME
        return [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}}
}

func quickSortEntry(array [][]int) [][]int {

        return array
}


func TestReconstructQueue(people [][]int) {
        fmt.Println("Input:")
        fmt.Printf("%s", IntPair2String(people))

        fmt.Println("Output:")
        fmt.Printf("%s", IntPair2String(reconstructQueue(people)))
}

func IntPair2String(pair [][]int) string {
        inputArray := make([]string, 0)
        for _, hk := range pair {
                inputArray = append(inputArray, fmt.Sprintf("[%s]", strings.Join(util.IntArray2StringArray(hk), ",")))
        }
        return fmt.Sprintf("[%s]\n\n", strings.Join(inputArray, ","))
}

func main() {
        TestReconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}})
}
