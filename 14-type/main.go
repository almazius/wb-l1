package main

import (
	"fmt"
	"reflect"
)

func main() {
	withFlag(int(8))
	withSwitch(byte(8))
	withReflect(int32(8))

}

func withFlag(val interface{}) {
	fmt.Printf("%T\n", val)
}

func withSwitch(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Println("int")
	case int64:
		fmt.Println("int64")
	case int32:
		fmt.Println("int32")
	case float32:
		fmt.Println("float32")
	case float64:
		fmt.Println("float64")
	case byte:
		fmt.Println("byte")
	default:
		fmt.Println("any type")
	}
}

func withReflect(val interface{}) {
	fmt.Println(reflect.TypeOf(val))
}
