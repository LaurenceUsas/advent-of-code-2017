package main

import (
	"fmt"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

// Task01 Solution
func Task01() {
	input := helpers.InputArg("Enter the digits")
	solution := task01PartOne(input)
	fmt.Printf("Solution: %d \n", solution)
}

func task01PartOne(input string) int {
	values := helpers.StringToIntArray(input)

	sum := 0
	// previous value
	pv := values[len(values)-1:][0]

	for _, v := range values {
		if pv == v {
			sum += v
		}
		pv = v
	}
	return sum
}
