package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	result := refreshString(str)
	fmt.Printf("First string: %s\nSecond string: %s\n", str, result)
}

func refreshString(str string) string {
	words := strings.Split(str, " ")
	builder := strings.Builder{}
	for i := len(words) - 1; i >= 0; i-- {
		builder.WriteString(words[i] + " ")
	}
	return builder.String()
}
