package main

import (
	"fmt"
	"math/big"
)

func main() {
	// we can use bigInt, bigFloat, rat
	firstValue := big.NewInt(99999)
	secondValue := big.NewInt(100000)
	result := big.NewInt(0)

	fmt.Println("v1 + v2", result.Add(firstValue, secondValue))
	result = big.NewInt(0)
	fmt.Println("v1 - v2", result.Add(firstValue, result.Neg(secondValue)))
	fmt.Println("v1 * v2", result.Mul(firstValue, secondValue))
	fmt.Println("v1 / v2", result.Div(firstValue, secondValue))

}

// Any try
// We can use BIG array
// first value - this sign, any values - number.
// For operations, we use fistValue[i] and secondValue[i]

// Any try
// We can use struct decimal
// type struct decimal {
// nums [4]uint32
// }
//
// all numbers - bits, when bits[0] - last bit from nums[0]
// bits[31] - first bit from nums[0] and etc
// more detailed in my repository from school21
// https://github.com/almazius/s21-dicimal
