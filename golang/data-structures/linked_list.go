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
	// testing purposes
	// expect [5, 9, 7, -3, 2]

	var linkedList linkedList = linkedList{head: nil, length: 0}

	linkedList.add(5, 0)
	linkedList.add(7, 1)
	linkedList.add(2, 2)
	linkedList.add(9, 1)
	linkedList.add(-3, 3)

	fmt.Println(linkedList)
	fmt.Println(linkedList.head.data)
	fmt.Println(linkedList.head.next.data)
	fmt.Println(linkedList.head.next.next.data)
	fmt.Println(linkedList.head.next.next.next.data)
	fmt.Println(linkedList.head.next.next.next.next.data)
}

func (list *linkedList) add(data int, pos int) {
	if list.length == 0 {
		var newNode node = node{next: nil, data: data}
		list.head = &newNode
		list.length++
		return
	}

	if pos > list.length {
		return
	}

	var head *node = list.head
	var idx int = 0

	for head.next != nil && idx < pos-1 {
		head = head.next
		idx++
	}

	var newNode node
	if head.next != nil {
		newNode = node{next: head.next, data: data}
	} else {
		newNode = node{next: nil, data: data}
	}

	head.next = &newNode
	list.length++
	return
}
