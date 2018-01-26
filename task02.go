package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

//Task02 Solution
func Task02() {
	pwd, _ := os.Getwd()
	input := helpers.InputFile(pwd + "/input/input02.txt")

	fmt.Printf("[Part 1 answer] - %v\n", task02PartOne(input))
	fmt.Printf("[Part 2 answer] - %v\n", task02PartTwo(input))
}

func task02PartOne(input []string) int {
	result := 0
	for _, v := range input {
		stringSlice := strings.Split(v, "\t")
		ints := make([]int, len(stringSlice))

		for i := range stringSlice {
			ints[i], _ = strconv.Atoi(string(stringSlice[i]))
		}

		sort.Ints(ints)
		diff := ints[len(ints)-1] - ints[0]
		result += diff
	}
	return result
}

// Brute Force
func task02PartTwo(input []string) int {
	result := 0
	for _, v := range input {
		stringSlice := strings.Split(v, "\t")
		ints := make([]int, len(stringSlice))

		for i := range stringSlice {
			ints[i], _ = strconv.Atoi(string(stringSlice[i]))
		}
		sort.Ints(ints)

		for a := 0; a < len(ints)-1; a++ {
			for aa := a + 1; aa < len(ints); aa++ {
				if ints[aa]%ints[a] == 0 {
					result += ints[aa] / ints[a]
				}
			}
		}
	}
	return result
}
