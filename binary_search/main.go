package main

import "fmt"

func main() {
	fmt.Println(binarySearchItetrative([]int{11, 12, 13, 14, 15}, 13))
	binarySearchRecursive([]int{11, 12, 13, 14, 15}, 13)
	fmt.Println(searchInRoatedSortArray([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

func binarySearchItetrative(a []int, target int) int {
	min := 0
	max := len(a) - 1

	for min <= max {
		mid := (min + max) / 2

		if a[mid] == target {
			return mid
		}

		if mid < max {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	return -1
}

func binarySearchRecursive(a []int, target int) {
	fmt.Println(search(a, 0, len(a)-1, target))
}

func search(a []int, min, max, target int) int {
	var mid int
	if max >= min {
		mid = (min + max) / 2
		if a[mid] == target {
			return mid
		} else if a[mid] < target {
			return search(a, mid+1, max, target)
		} else {
			return search(a, min, mid-1, target)
		}
	}

	return -1
}

// You are given an integer array nums sorted in ascending order (with distinct values), and an integer target.

// Suppose that nums is rotated at some pivot unknown to you beforehand (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

// If target is found in the array return its index, otherwise, return -1.
func searchInRoatedSortArray(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2

		switch {
		case target == nums[mid]:
			return mid
		case target == nums[low]:
			return low
		case target == nums[high]:
			return high
		}

		if nums[low] > nums[high] {
			low++
			high--
		} else {
			if nums[mid] > target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}
