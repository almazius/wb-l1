package main

import "fmt"

func main() {
	firstSet := []int{1, 2, 4, 5, 6, 7, 8, 9}
	secondSet := []int{1, 4, 5, 7, 9, 12, 436, 24}

	result := intersectionSets(firstSet, secondSet)

	for _, val := range result {
		fmt.Println(val)
	}
}

func intersectionSets(firstSet, secondSet []int) []int {
	resultSet := make([]int, 0)
	checkMap := make(map[int]struct{})

	for _, val := range firstSet {
		checkMap[val] = struct{}{}
	}

	for _, val := range secondSet {
		if _, exist := checkMap[val]; exist {
			resultSet = append(resultSet, val)
		}
	}
	return resultSet
}
