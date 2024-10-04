package main

import "fmt"

func main() {
	arr := []int{5, 3, 67, 8, 533, 66, 33, 776, 8894, 4759, 5475, 35, 4, 55, 5442, 123, 5432, 12342}
	lo := 0
	hi := len(arr) - 1

	sortArr := qs(arr, lo, hi)
	fmt.Println(sortArr)
}

func qs(arr []int, lo int, hi int) []int {
	if lo >= hi {
		return arr
	}
	pivotIndex := partition(arr, lo, hi)

	_ = qs(arr, lo, pivotIndex-1)
	_ = qs(arr, pivotIndex+1, hi)
	return arr
}

func partition(arr []int, lo int, hi int) int {
	pivot := arr[hi]

	index := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			index++
			tmp := arr[i]
			arr[i] = arr[index]
			arr[index] = tmp
		}
	}

	index++
	arr[hi] = arr[index]
	arr[index] = pivot
	return index

}
