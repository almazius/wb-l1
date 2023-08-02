package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	sum := new(int64)
	arr := []int64{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	for i := range arr {
		wg.Add(1)
		sumSquares(arr[i], sum, &wg)
	}

	// the delay is necessary so that the program does not end before the function is executed
	wg.Wait()
	fmt.Println(*sum)
}

func sumSquares(val int64, sum *int64, wg *sync.WaitGroup) {
	// use atomic for avoid race condition
	atomic.AddInt64(sum, val*val)
	wg.Done()
}

// создаем канал чисел, и конкуренто, в другой горутине вычисляем значения квадратов, далее просто их кидаем в канал
