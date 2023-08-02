package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {
	var (
		duration time.Duration
		err      error
		wg       sync.WaitGroup
	)
	ch := make(chan int)

	// parse duration
	if len(os.Args) > 1 {
		duration, err = time.ParseDuration(os.Args[1])
		if err != nil {
			duration = 5 * time.Second
			log.Println(err, "\nDuration = 5")
		}
		if duration < 0 {
			log.Printf("Duration must be biggest 0!\n Duration = 5")
			duration = 5 * time.Second
		}
	}

	//create context and cancel function
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	// Add goroutine on wait group
	wg.Add(2)

	go publisher(ctx, ch, &wg)
	go consumer(ctx, ch, &wg)

	// wait goroutine
	wg.Wait()
}

func consumer(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	for {
		select {
		// when ctx use cancel function, consumer finished
		case <-ctx.Done():
			fmt.Printf("consumer finished\n")
			wg.Done()
			return
		default:
			ch <- rand.Int()
			time.Sleep(time.Second)
		}
	}
}

func publisher(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	for {
		select {
		case val := <-ch:
			fmt.Printf("You send: %d\n", val)

		// when ctx use cancel function, consumer finished
		case <-ctx.Done():
			fmt.Printf("publisher finished\n")
			wg.Done()
			return
		}
	}
}
