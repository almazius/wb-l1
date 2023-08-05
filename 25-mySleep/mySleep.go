package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("1")
	mySleep(3 * time.Second)
	fmt.Println("3")
}

// mySleep time.After block goroutine on duration time and unlock after time
func mySleep(duration time.Duration) {
	<-time.After(duration)
}
