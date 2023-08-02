package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	nums := make(map[int]int)
	mutex := sync.Mutex{}
	var val int32
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeOnMap(ctx, &wg, nums, &mutex, &val)
	}

	wg.Wait()
	fmt.Println(val)

}

func writeOnMap(ctx context.Context, wg *sync.WaitGroup, nums map[int]int, mutex *sync.Mutex, val *int32) {
	for {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			wg.Done()
			return
		default:
			mutex.Lock()
			nums[rand.Int()] = rand.Int()
			atomic.AddInt32(val, 1)
			mutex.Unlock()
		}
	}
}
