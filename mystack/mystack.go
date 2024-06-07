package mystack

import (
	"errors"
	"fmt"
)

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Cap() int {
	return 0
}

func (stack Stack) IsEmpty() bool {
	return stack.Len() == 0
}

func (stack *Stack) Push(e interface{}) {
	*stack = append(*stack, e)
}

func (stack Stack) Top() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("empty stack")
	}
	return stack[stack.Len()-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("empty stack")
	}
	e := (*stack)[stack.Len()-1]
	fmt.Println("before stack", *stack)
	*stack = (*stack)[:stack.Len()-1]
	fmt.Println("after stack", *stack)
	return e, nil
}
