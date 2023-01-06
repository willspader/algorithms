package main

import "fmt"

type queue struct {
	beginning *node
	end       *node
	length    int
}

type node struct {
	elem interface{}
	next *node
}

func main() {
	// create new queue. This is how users would use in a project. The structs are private.
	var queue queue = New()

	queue.enqueue(5)
	queue.enqueue(1)
	queue.enqueue(-4)
	queue.enqueue(3)

	fmt.Println(queue.beginning.elem)
	fmt.Println(queue.beginning.next.elem)
	fmt.Println(queue.beginning.next.next.elem)
	fmt.Println(queue.beginning.next.next.next.elem)

	fmt.Println(queue.end.elem)

	queue.dequeue()
	queue.dequeue()

	fmt.Println(queue.beginning.elem)
	fmt.Println(queue.beginning.next.elem)

	queue.dequeue()
	queue.dequeue()

	fmt.Println(queue.beginning)
	fmt.Println(queue.end)
	fmt.Println(queue.length)
}

func (queue *queue) enqueue(elem interface{}) {
	if queue.beginning == nil {
		queue.beginning = &node{elem: elem, next: nil}
		queue.end = &node{elem: elem, next: nil}
		queue.length++
		return
	}

	var newEnd *node = &node{elem: elem, next: nil}

	queue.end.next = newEnd
	queue.end = newEnd

	queue.length++

	if queue.length == 2 {
		queue.beginning.next = newEnd
	}
}

func (queue *queue) dequeue() interface{} {
	if queue.beginning == nil {
		return nil
	}

	var dequeued interface{} = queue.beginning.elem

	if queue.length >= 2 {
		queue.beginning = queue.beginning.next
	} else {
		queue.beginning = nil
		queue.end = nil
	}

	queue.length--

	return dequeued
}

func New() queue {
	return queue{beginning: nil, end: nil, length: 0}
}
