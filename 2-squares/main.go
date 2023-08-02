package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wg := sync.WaitGroup{}
	for i := range arr {
		wg.Add(1)
		go squares(arr[i], &wg)
	}

	// the delay is necessary so that the program does not end before the function is executed
	wg.Wait()
}

func squares(val int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(val * val)
}
