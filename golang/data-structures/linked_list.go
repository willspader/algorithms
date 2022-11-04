package main

import (
	"fmt"
)

type linkedList struct {
	head   *node
	length int
}

type node struct {
	next *node
	data int
}

func main() {
	var linkedList linkedList = linkedList{head: nil, length: 0}

	linkedList.insertAt(5, 0)
	linkedList.insertAt(7, 1)
	linkedList.insertAt(2, 2)
	linkedList.insertAt(9, 1)
	linkedList.insertAt(-3, 3)

	// expect [5, 9, 7, -3, 2]
	fmt.Println(linkedList)
	fmt.Println(linkedList.head.data)
	fmt.Println(linkedList.head.next.data)
	fmt.Println(linkedList.head.next.next.data)
	fmt.Println(linkedList.head.next.next.next.data)
	fmt.Println(linkedList.head.next.next.next.next.data)

	linkedList.removeAt(0)
	linkedList.removeAt(3)
	linkedList.removeAt(1)

	// expect [9, -3]
	fmt.Println(linkedList)
	fmt.Println(linkedList.head.data)
	fmt.Println(linkedList.head.next.data)

	linkedList.insertAt(0, 1)
	linkedList.insertAt(7, 3)
	linkedList.insertAt(4, 4)
	linkedList.insertAt(2, 0)

	// expect [2, 9, 0, -3, 7, 4]
	fmt.Println(linkedList)
	fmt.Println(linkedList.head.data)
	fmt.Println(linkedList.head.next.data)
	fmt.Println(linkedList.head.next.next.data)
	fmt.Println(linkedList.head.next.next.next.data)
	fmt.Println(linkedList.head.next.next.next.next.data)
	fmt.Println(linkedList.head.next.next.next.next.next.data)
}

func (list *linkedList) insertAt(data int, pos int) {
	if list.length == 0 {
		list.head = &node{next: nil, data: data}
		list.length++
		return
	}

	if pos == 0 {
		list.head = &node{next: list.head, data: data}
		list.length++
		return
	}

	var prev *node = list.getAt(pos)

	if prev != nil {
		prev.next = &node{next: prev.next, data: data}
		list.length++
	}
}

func (list *linkedList) removeAt(pos int) {
	if list.length == 0 {
		return
	}

	if pos == 0 {
		if list.length == 1 {
			list.head = nil
		} else {
			list.head = list.head.next
		}
		list.length--
		return
	}

	var prev *node = list.getAt(pos)

	if prev == nil {
		return
	}

	if pos == list.length-1 {
		prev.next = nil
	} else {
		prev.next = prev.next.next
	}
	list.length--
}

func (list *linkedList) getAt(pos int) *node {
	if pos > list.length {
		return nil
	}

	var head *node = list.head
	var idx int = 0

	for head.next != nil && idx < pos-1 {
		head = head.next
		idx++
	}

	return head
}
