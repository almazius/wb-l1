package main

import "fmt"

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.9, 15.5, 24.5, -21.0, 32.5}
	result := make(map[int][]float64)

	parse(arr, result)

	for key, val := range result {
		fmt.Println(key * 10)
		for i := range val {
			fmt.Printf("\t%f\n", val[i])
		}
	}
}

func parse(arr []float64, result map[int][]float64) {
	for _, val := range arr {
		if _, ok := result[int(val/10)]; ok {
			result[int(val/10)] = append(result[int(val/10)], val)
		} else {
			result[int(val/10)] = make([]float64, 0)
			result[int(val/10)] = append(result[int(val/10)], val)
		}
	}
}
