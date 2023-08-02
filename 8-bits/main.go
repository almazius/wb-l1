package main

import "fmt"

func main() {
	var val int64 = 32
	changeBit(&val, 5, 1)
	fmt.Println(val)
}

func changeBit(value *int64, pos, bit int64) {
	if bit == 1 {
		*value = *value | (bit << pos)
	} else if bit == 0 {
		*value = *value & (bit << pos)
	}
}
