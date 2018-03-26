package main

import (
	"fmt"
	ds "github.com/berryjam/leetcode-golang/datastructure"
	"os"
	"strconv"
)

const FORMAT_STRING = "s = \"%s\",retrun \"%s\"\n"

func DecodeString(s string) string {
	stack := &ds.Stack{}
	res := ""
	in_stack_state := false
	in_num_state := false
	tmp_str := ""
	tmp_num := ""
	for _, char := range s {
		if char >= '0' && char <= '9' {
			if in_stack_state {
				tmp_str += string(byte(char))
			} else {
				if in_num_state {
					tmp_num += string(byte(char))
				} else {
					tmp_num = string(byte(char))
					in_num_state = true
				}
			}
		} else if char == '[' || char == ']' {
			if char == '[' {
				in_num_state = false
				if in_stack_state {
					tmp_str += string(byte(char))
					stack.Push(char)
				} else {
					in_stack_state = true
					stack.Push(char)
				}
			} else {
				stack.Pop()
				if stack.Len() == 0 {
					count, err := strconv.Atoi(tmp_num)
					if err != nil {
						fmt.Printf("Error input format:%s", s)
						os.Exit(1)
					}
					for i := 0; i < count; i++ {
						res += DecodeString(tmp_str)
					}
					in_stack_state = false
					tmp_str = ""
				} else {
					tmp_str += string(byte(char))
				}
			}
		} else { // alaphabet
			if in_stack_state {
				tmp_str += string(byte(char))
			} else {
				res += string(byte(char))
			}
		}
	}

	return res
}

func main() {
	s1 := "3[a]2[bc]"
	s2 := "3[a2[c]]"
	s3 := "2[abc]3[cd]ef"
	fmt.Printf(FORMAT_STRING, s1, DecodeString(s1))
	fmt.Printf(FORMAT_STRING, s2, DecodeString(s2))
	fmt.Printf(FORMAT_STRING, s3, DecodeString(s3))
}
