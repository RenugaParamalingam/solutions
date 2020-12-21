package main

import "fmt"

func main() {
	fmt.Println(sortBinaryArrayBySwapping([]int{1, 0, 1, 0, 1, 0}))
	fmt.Println(sortBinaryArray([]int{1, 0, 1, 0, 1, 0}))
}

func sortBinaryArrayBySwapping(arr []int) []int {
	j := -1

	for i := range arr {
		if arr[i] < 1 {
			j++

			temp := arr[j]
			arr[j] = arr[i]
			arr[i] = temp
		}
	}

	return arr
}

func sortBinaryArray(arr []int) []int {
	zeroCount := 0

	for _, a := range arr {
		if a == 0 {
			zeroCount++
		}
	}

	for i := range arr {
		if i < zeroCount {
			arr[i] = 0
			continue
		}
		arr[i] = 1
	}

	return arr
}
