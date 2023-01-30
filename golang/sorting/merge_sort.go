package main

import "fmt"

func main() {
	var arr []int = []int{1, 5, -1, 0, 7, 15, 7, 10}

	var sorted []int = mergeSort(arr)

	fmt.Println(sorted)
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var beg []int = mergeSort(arr[:len(arr)/2])
	var end []int = mergeSort(arr[len(arr)/2:])

	return merge(beg, end)
}

func merge(beg []int, end []int) []int {
	var final []int
	var i int = 0
	var j int = 0

	for i < len(beg) && j < len(end) {
		if beg[i] < end[j] {
			final = append(final, beg[i])
			i++
		} else {
			final = append(final, end[j])
			j++
		}
	}

	for i < len(beg) {
		final = append(final, beg[i])
		i++
	}

	for j < len(end) {
		final = append(final, end[j])
		j++
	}

	return final
}
