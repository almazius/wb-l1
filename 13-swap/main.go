package main

import "fmt"

func main() {
	a, b := 10, 20
	fmt.Printf("a = %d, b = %d\n", a, b)

	b, a = a, b
	fmt.Printf("a = %d, b = %d\n", a, b)
}
