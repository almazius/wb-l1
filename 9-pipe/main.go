package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	out := readValue(arr)
	writeValue(out)
}

func readValue(arr []int) <-chan int {
	out := make(chan int, len(arr))
	for _, val := range arr {
		out <- val
	}
	close(out)
	return out
}

func writeValue(in <-chan int) {
	for n := range in {
		fmt.Println(n * n)
	}
}
