package main

/*
You receive a signal directly from the CPU. Because of your recent assistance with jump instructions, it would like you to compute the result of a series of unusual register instructions.

Each instruction consists of several parts: the register to modify, whether to increase or decrease that register's value, the amount by which to increase or decrease it, and a condition. If the condition fails, skip the instruction without modifying the register. The registers all start at 0. The instructions look like this:

b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10
These instructions would be processed as follows:

Because a starts at 0, it is not greater than 1, and so b is not modified.
a is increased by 1 (to 1) because b is less than 5 (it is 0).
c is decreased by -10 (to 10) because a is now greater than or equal to 1 (it is 1).
c is increased by -20 (to -10) because c is equal to 10.
After this process, the largest value in any register is 1.

You might also encounter <= (less than or equal to) or != (not equal to). However, the CPU doesn't have the bandwidth to tell you what all the registers are named, and leaves that to you to determine.

What is the largest value in any register after completing the instructions in your puzzle input?


*/

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

/*
a dec -511 if x >= -4
pq inc -45 if cfa == 7
vby dec 69 if tl < 1
*/
var reg = map[string]int{}

//Task08 Solution
func Task08() {
	var max int

	input := helpers.InputFile("/input/input08.txt")

	for _, line := range input {
		elements := strings.Split(line, " ")

		if checkCondition(elements[4], elements[5], elements[6]) {
			value := 0
			change, _ := strconv.Atoi(elements[2])
			if v, ok := reg[elements[0]]; ok {
				value = v
			}

			if elements[1] == "inc" {
				reg[elements[0]] = value + change
			} else {
				reg[elements[0]] = value - change
			}
		}
		// Part 2
		if max < reg[elements[0]] {
			max = reg[elements[0]]
		}
	}

	//find largerst val - Part 1
	var h int
	for _, v := range reg {
		if v > h {
			h = v
		}
	}
	fmt.Printf("Part 1 answer: %d\n", h)
	fmt.Printf("Part 2 answer: %d\n", max)
}

func checkCondition(dest, cond, op string) bool {
	value := 0
	if v, ok := reg[dest]; ok {
		value = v
	}

	expected, _ := strconv.Atoi(op)

	switch {
	case cond == "<=" && value <= expected:
		return true
	case cond == ">=" && value >= expected:
		return true
	case cond == "==" && value == expected:
		return true
	case cond == "!=" && value != expected:
		return true
	case cond == "<" && value < expected:
		return true
	case cond == ">" && value > expected:
		return true
	default:
		return false
	}
}
