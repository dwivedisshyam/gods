package main

import "fmt"

type Stack[T comparable] struct {
	Items []T
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

func (stack *Stack[T]) Push(item T) {
	stack.Items = append(stack.Items, item)
}

func (stack *Stack[T]) Pop() T {
	if !stack.isEmpty() {
		item := stack.Items[len(stack.Items)-1]
		stack.Items = stack.Items[:len(stack.Items)-1]
		return item
	}
	return *new(T)
}

func (stack *Stack[T]) Top() T {
	if !stack.isEmpty() {
		return stack.Items[len(stack.Items)-1]
	}

	return *new(T)
}

func (stack *Stack[T]) isEmpty() bool {
	return len(stack.Items) == 0
}

func (stack *Stack[T]) PrintStack() {
	fmt.Printf("stack is %v ", stack.Items)
}
