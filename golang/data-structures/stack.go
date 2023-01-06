package main

import "fmt"

type stack struct {
	head   *node
	length int
}

type node struct {
	data int
	next *node
}

func main() {
	var stack stack = stack{}

	stack.push(5)
	stack.push(7)
	stack.push(3)
	stack.push(-2)
	stack.push(10)

	fmt.Println(stack.head.data)
	fmt.Println(stack.head.next.data)
	fmt.Println(stack.head.next.next.data)
	fmt.Println(stack.head.next.next.next.data)
	fmt.Println(stack.head.next.next.next.next.data)

	fmt.Println(*stack.pop())
	fmt.Println(*stack.pop())

	fmt.Println(stack.head.data)
	fmt.Println(stack.head.next.data)
	fmt.Println(stack.head.next.next.data)
}

func (stack *stack) push(data int) {
	if stack == nil {
		return
	}

	if stack.head == nil {
		stack.head = &node{data: data, next: nil}
		stack.length = 0
		return
	}

	stack.head = &node{data: data, next: stack.head}
	stack.length++
}

func (stack *stack) pop() *int {
	if stack.head == nil {
		return nil
	}

	var removed int = stack.head.data

	if stack.head.next == nil {
		stack.head = nil
	}

	stack.head = stack.head.next
	stack.length--

	return &removed
}
