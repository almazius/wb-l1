package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Counter struct {
	i int32
}

func main() {
	var counter Counter
	done := make(chan struct{}, 1)

	for i := 0; i < 10; i++ {
		go increment(&counter, done)
	}

	time.Sleep(5 * time.Second)
	close(done)
	fmt.Println(counter.i)
}

func increment(counter *Counter, done chan struct{}) {
	for {
		select {
		case <-done:
			return
		default:
			atomic.AddInt32(&counter.i, 1)
		}
	}
}
