package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(checkString("abcdef"))
}

func checkString(s string) bool {
	checkMap := make(map[rune]struct{})
	s = strings.ToLower(s)

	for _, r := range s {
		if _, exist := checkMap[r]; !exist {
			checkMap[r] = struct{}{}
		} else {
			return false
		}
	}
	return true
}
