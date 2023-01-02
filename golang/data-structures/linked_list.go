package main

import "fmt"

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

	linkedList.insert(5)
	linkedList.insert(5)
	linkedList.insert(7)
	linkedList.insert(5)
	linkedList.insert(10)
	linkedList.insert(5)

	linkedList.remove(5)

	fmt.Println(linkedList)
	fmt.Println(linkedList.head.data)
	fmt.Println(linkedList.head.next.data)
}

// It inserts a node at the end of the list.
func (list *linkedList) insert(data int) {
	if list.length == 0 {
		list.head = &node{next: nil, data: data}
		list.length++
		return
	}

	var head *node = list.head

	for head.next != nil {
		head = head.next
	}

	head.next = &node{next: nil, data: data}
	list.length++
}

// It removes all node occurrences of `data`. If the list is empty, then it does nothing.
func (list *linkedList) remove(data int) {
	if list.length == 0 {
		return
	}

	if list.length == 1 && list.head.data == data {
		list.head = nil
	}

	var head *node = list.head
	for head.next != nil {
		var prev *node = head

		if head.next.data == data {
			if head.next.next != nil {
				prev.next = head.next.next
				head = head.next
			} else {
				prev.next = nil
			}
			list.length--
		} else {
			head = head.next
		}
	}

	if list.length > 1 && list.head.data == data {
		list.head = list.head.next
		list.length--
	}
}

// It inserts a node at the specified position. If the position is bigger than list.length, then it does nothing.
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

	var prev *node = list.getPrev(pos)

	if prev != nil {
		prev.next = &node{next: prev.next, data: data}
		list.length++
	}
}

// It removes a node from the specified position. If the position is >= list.length, then it does nothing.
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

	var prev *node = list.getPrev(pos)

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

// It gets a previous node from the specified position.
func (list *linkedList) getPrev(pos int) *node {
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
