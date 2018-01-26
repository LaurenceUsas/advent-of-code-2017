package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

// Task01 Solution
func Task01() {
	pwd, _ := os.Getwd()
	input := helpers.InputFile(pwd + "/input/input01.txt")[0]

	fmt.Printf("[Part 1 answer] - %v\n", task01PartOne(input, 1))
	fmt.Printf("[Part 2 answer] - %v\n", task01PartOne(input, len(input)/2))

}

func task01PartOne(input string, offset int) int {
	l := len(input)
	sum := 0

	for i, v := range input {
		if byte(v) == input[(i+offset)%l] {
			d, _ := strconv.Atoi(string(v))
			sum += d
		}
	}
	return sum
}
