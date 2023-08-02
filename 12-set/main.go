package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	result := createSet(arr)

	for _, val := range result {
		fmt.Println(val)
	}
}

func createSet(arr []string) []string {
	result := make([]string, 0)
	checkMap := make(map[string]struct{})

	for _, val := range arr {
		checkMap[val] = struct{}{}
	}

	for val := range checkMap {
		result = append(result, val)
	}
	return result
}
