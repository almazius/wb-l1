package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// Global value - is bad :(
//var justString string
//
// Bad name on function
//func someFunc() {
// (1 << 10) why not?.. But I use 1024
//	v := createHugeString(1 << 10)
//
// len(v) == 1024, why we use only 100 symbol?
// if len(v) < 100, we have problems
//	justString = v[:100]
//}
//
// nice name on function!
//func main() {
//	someFunc()
//}

// we use int64 because len string can very big,
// we use *string because size string can very big
func CreateBigString(size int64) *strings.Builder {
	builder := &strings.Builder{}
	for i := int64(0); i < size; i++ {
		builder.WriteRune(rune('a' + rand.Int()%26))
	}
	return builder
}

func main() {
	fmt.Println(CreateBigString(1024))
}
