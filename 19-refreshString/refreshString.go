package main

import (
	"fmt"
	"strings"
)

func main() {
	str := ""

	newStr := refresh(str)
	fmt.Println(newStr)
}

func refresh(s string) string {
	b := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteRune(rune(s[i]))
	}
	return b.String()
}
