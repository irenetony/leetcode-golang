package main

import (
	"fmt"
	"strconv"
	"os"
)

const FORMAT_STRING = "s = \"%s\",retrun \"%s\"\n"

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{} // All types satisfy the empty interface, so we can store anything here.
	next  *Element
}

// Return the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func DecodeString(s string) string {
	stack := &Stack{}
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
