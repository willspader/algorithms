package main

import (
	"fmt"
)

func main() {
	var arr []int8 = []int8{1, 5, -1, 0, 7, 15, 7, 10}

	insertionSort(arr)

	fmt.Println(arr)
}

// O(n^2)
func insertionSort(arr []int8) {
	var slice_size int = len(arr) // O(1)
	for i := 1; i < slice_size; i++ {
		var key int8 = arr[i]
		var j int = i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
