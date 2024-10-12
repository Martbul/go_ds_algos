package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 5, 6, 7, 8, 13, 15, 18, 22, 24, 25, 29, 33, 35, 36, 37, 55, 452, 1342342}
	needle := 18

	res := binarySearch(arr, needle)
	fmt.Println(res)
}

func binarySearch(arr []int, needle int) bool {
	lo := 0
	hi := len(arr) - 1 // Change this line

	for lo <= hi { // Change the condition to lo <= hi
		half := lo + (hi-lo)/2 // Correct midpoint calculation

		if arr[half] == needle {
			return true
		}

		if arr[half] < needle {
			lo = half + 1
		} else { // Use else here to avoid unnecessary checks
			hi = half - 1
		}
	}

	return false
}
