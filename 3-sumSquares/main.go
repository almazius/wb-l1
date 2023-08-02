package main

import (
	"fmt"
	"sync"
)

func main() {
	sum := new(int)
	arr := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	for i := range arr {
		wg.Add(1)
		sumSquares(arr[i], sum, &wg)
	}

	// the delay is necessary so that the program does not end before the function is executed
	wg.Wait()
	fmt.Println(*sum)
}

func sumSquares(val int, sum *int, wg *sync.WaitGroup) {
	*sum += val * val
	wg.Done()
}
