package main

import (
	"math/rand"
	"sort"
)

func main() {
	slice := make([]int, 99999)
	for i := range slice {
		slice[i] = rand.Int()
	}

	// package "sort" always use intro sort, who select quick sort for big array
	sort.Ints(slice)
}

// my quick sort
func myQuickSort(slice []int, l, r int) {
	if l < r {
		medium := partion(slice, l, r)
		myQuickSort(slice, l, medium-1)
		myQuickSort(slice, medium+1, r)
	}
}

func partion(slice []int, l int, r int) int {
	pivot := slice[r]
	i := l - 1

	for j := l; j < r; j++ {
		if slice[j] < pivot {
			i++

			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[r] = slice[r], slice[i+1]
	return i + 1
}
