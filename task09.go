package main

import (
	"fmt"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

// Task09 Solution
func Task09() {
	input := helpers.InputFile("/input/input09.txt")[0]

	skip, garbage := false, false
	cval := 0
	sum := 0
	gsum := 0
	for _, c := range input {
		switch {
		case skip:
			skip = false
		case c == '!':
			skip = true
		case garbage && c == '>':
			garbage = false
		case garbage:
			gsum++
		case c == '<':
			garbage = true
		case !garbage && c == '{':
			cval++
		case !garbage && c == '}':
			sum += cval
			cval--
		}
	}
	fmt.Printf("Part 1 answer: %d", sum)
	fmt.Printf("Part 2 answer: %d", gsum)
}
