package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	slice := make([]int, 10)
	for i := range slice {
		slice[i] = i
	}

	slice, err := dellItem(slice, 9)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d ", slice[i])
	}

}

func dellItem(slice []int, index int) ([]int, error) {
	if index >= len(slice) || index < 0 {
		return nil, errors.New("incorrect index")
	} else if len(slice) == 1 {
		return nil, nil
	} else {
		if index == 0 {
			slice = slice[1:]
		} else if index == len(slice)-1 {
			slice = slice[:len(slice)-1]
		} else {
			slice = append(slice[:index], slice[index+1:]...)

		}
	}
	return slice, nil
}
