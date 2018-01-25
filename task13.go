package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

//Task13 Solution
func Task13() {
	pwd, _ := os.Getwd()
	file := helpers.InputFile(pwd + "/input/input13.txt")
	input := make(map[int]int)

	for i := range file {
		ab := strings.Split(file[i], ": ")
		a, _ := strconv.Atoi(ab[0])
		b, _ := strconv.Atoi(ab[1])
		input[a] = b
	}

	fmt.Printf("[Part 1 answer] - %v\n", task13PartOne(input))
	//read input.
	//solve part 1
	//solve part 2
}

func task13PartOne(input map[int]int) int {
	//penalty
	p := 0
	for k, v := range input {
		if ok := willCatch(k, v); ok {
			p += k * v
		}
	}
	return p
}

func willCatch(id, d int) bool {
	i := 0
	v := 0
	for v < id {
		//TODO optimize, check if already counted. return.
		v = 2 * (d - 1) * i
		fmt.Printf("At iteration [%d] will be [%d]\n", i, v)
		i++
	}
	if v == id {
		return true
	}
	return false
}

func task13PartTwo() {
	// Find offset of start so it goes without a catch.
}
