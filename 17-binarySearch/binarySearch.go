package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	// function search from package sort use binary search
	i := sort.Search(len(arr), func(i int) bool { return arr[i] >= x })
	if i < len(arr) && arr[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, arr)
	} else {
		fmt.Printf("%d not found in %v\n", x, arr)
	}
}

// my binary search
func myBinarySearch(slice []int, query int) int {
	minIndex := 0
	maxIndex := len(slice) - 1

	for minIndex <= maxIndex {
		midIndex := int((maxIndex + minIndex) / 2)
		midItem := slice[midIndex]
		if query == midItem {
			return midIndex
		}
		if midItem < query {
			minIndex = midIndex + 1
		} else if midItem > query {
			maxIndex = midIndex - 1
		}
	}

	// if we not find value, return -1
	return -1
}
