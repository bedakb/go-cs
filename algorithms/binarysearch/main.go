package main

import (
	"fmt"
)

func main() {
	s := []int{3, 5, 8, 20, 30, 66}
	fmt.Println(binarySearch(5, s))
}

func binarySearch(val int, ol []int) int {
	left := 0
	right := len(ol) - 1
	mid := (left + right) / 2

	if val == ol[mid] {
		return mid
	}

	for left <= right {
		if val < ol[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}

	if val != ol[mid] {
		return -1
	}

	return mid

}
