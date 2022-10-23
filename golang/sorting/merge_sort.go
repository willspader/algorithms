package main

func main() {

}

func mergeSort(arr []int, p int, r int) {
	if p < r {
		var q int = (p + r) / 2
		mergeSort(arr, p, q)
		mergeSort(arr, q+1, r)
		merge(arr, p, q, r)
	}
}

func merge(arr []int, p int, q int, r int) {

}
