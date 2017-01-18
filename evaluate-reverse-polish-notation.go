package main

import (
	"fmt"
	"strings"
	"strconv"
)

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

func (stack *Stack) Push(value interface{}) {
	stack.top = &Element{value, stack.top}
	stack.size++
}

func (stack *Stack) Len() int {
	return stack.size
}

func (stack *Stack) Pop() (value interface{}) {
	if stack.size > 0 {
		value, stack.top = stack.top.value, stack.top.next
		stack.size--
		return
	}
	return nil
}

func EvalRPN(tokens []string) int {
	var numStack Stack
	for _, token := range tokens {
		switch token {
		case "+":
			postOperand := numStack.Pop().(int)
			preOperand := numStack.Pop().(int)
			numStack.Push(preOperand + postOperand)
		case "-":
			postOperand := numStack.Pop().(int)
			preOperand := numStack.Pop().(int)
			numStack.Push(preOperand - postOperand)
		case "*":
			postOperand := numStack.Pop().(int)
			preOperand := numStack.Pop().(int)
			numStack.Push(preOperand * postOperand)
		case "/":
			postOperand := numStack.Pop().(int)
			preOperand := numStack.Pop().(int)
			numStack.Push(preOperand / postOperand)
		default:
			operand, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			numStack.Push(operand)
		}
	}

	return numStack.Pop().(int)
}

func TestEvalRPN() {
	tokens1 := []string{"2", "1", "+", "3", "*"}
	tokens2 := []string{"4", "13", "5", "/", "+"}

	fmt.Printf("[ %s ] ->  %d\n", strings.Join(tokens1, ","), EvalRPN(tokens1))
	fmt.Printf("[ %s ] -> %d\n", strings.Join(tokens2, ","), EvalRPN(tokens2))
}

func main() {
	TestEvalRPN()
}
