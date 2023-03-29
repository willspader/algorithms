package main

import (
	"fmt"
)

func main() {
	fmt.Println(binarySearch([]int{0, 3, 5, 6, 7}, 0, 4, 15))
}

func binarySearch(arr []int, low int, high int, data int) int {
	if low > high {
		return -1
	}

	var mid int = (low + high) / 2

	if arr[mid] == data {
		return mid
	}

	if arr[mid] > data {
		return binarySearch(arr, low, mid-1, data)
	} else {
		return binarySearch(arr, mid+1, high, data)
	}
}
