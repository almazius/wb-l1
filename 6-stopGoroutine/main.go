package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go writeOnUnBuffChannel()
	go writeOnBuffChannel()
	// go readFromClosedChannel() // use panic, but stop goroutine :)
	go withWaitGroup()
	go withScanFromStdin()
	time.Sleep(3 * time.Second)
}

func writeOnUnBuffChannel() {
	ch := make(chan int)
	ch <- 12

	fmt.Printf("writeOnUnBuffChannel is not stopped\n")
}

func writeOnBuffChannel() {
	ch := make(chan int, 3)
	ch <- 12
	ch <- 12
	ch <- 12
	ch <- 12

	fmt.Printf("writeOnBuffChannel is not stopped\n")
}

func readFromClosedChannel() {
	ch := make(chan int, 2)
	ch <- 1
	close(ch)

	ch <- 2

	fmt.Println("readFromClosedChannel not use panic")
}

func withWaitGroup() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
	fmt.Println("withWaitGroup is not stopped")
}

func withScanFromStdin() {
	str := ""
	fmt.Scan(&str)
	fmt.Println(" withScanFromStdin is not stopped")
}
