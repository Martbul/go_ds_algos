package main

import "fmt"

func main() {
	arr := []int{342, 234, 675, 45234, 453, 234, 345, 74536, 63446, 23, 2345436, 64, 234, 12, 7563, 345345345345696, 06356, 34534534534, 534532342, 3423424234}
	res := bubbleSort(arr)
	fmt.Println(res)
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}
