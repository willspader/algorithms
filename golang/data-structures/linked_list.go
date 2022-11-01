package main

import (
	"fmt"
)

type LinkedList struct {
	next *LinkedList
	data int
}

func main() {
	var third LinkedList = LinkedList{next: nil, data: 7}

	var second LinkedList = LinkedList{next: &third, data: 5}

	var first LinkedList = LinkedList{next: &second, data: 10}

	fmt.Println(first.next.data)      // 5
	fmt.Println(first.data)           // 10
	fmt.Println(first.next.next.data) // 7
}
