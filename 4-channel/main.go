package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// worker is a function for process values from channel. All worker have id.
func worker(ctx context.Context, wg *sync.WaitGroup, ch <-chan int, id int) {
	for {
		select {
		case val := <-ch:
			fmt.Printf("Worker %d: %d\n", id, val)
		case <-ctx.Done():
			fmt.Printf("Worker %d finished\n", id)
			wg.Done()
			return
		}
	}
}

func main() {
	var (
		countWorker int = 1
		err         error
	)
	ch := make(chan int)
	wg := sync.WaitGroup{}

	// create context and function cancel
	ctx, cansel := context.WithCancel(context.Background())

	// get and parse countWorker
	if len(os.Args) > 1 {
		countWorker, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		if countWorker < 1 {
			countWorker = 1
		}
	}

	// start workers
	for i := 0; i < countWorker; i++ {
		wg.Add(1)
		go worker(ctx, &wg, ch, i)
	}

	// function for add values in channel
	go func() {
		for {
			ch <- rand.Int()
			time.Sleep(time.Second)
		}
	}()

	// create gracefulShotDown through process signal from OS
	gracefulShotDown := make(chan os.Signal, 1)
	signal.Notify(gracefulShotDown, syscall.SIGINT, syscall.SIGTERM)
	<-gracefulShotDown

	cansel()
	fmt.Println("Program was finished")
	wg.Wait()
}
