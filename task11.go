package main

import (
	"fmt"
	"strings"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

//Task11 Solution
func Task11() {
	input := helpers.InputFile("/input11.txt")[0]
	instructions := strings.Split(input, ",")
	fmt.Printf("[Part 1 answer] - %v", task11PartOne(instructions))
	//fmt.Printf("[Part 2 answer] - %v", part2(instructions))
}

func task11PartOne(instructions []string) int {
	var d = [6]int{}
	for _, v := range instructions {
		switch v {
		case "n":
			d[0]++
		case "ne":
			d[1]++
		case "se":
			d[2]++
		case "s":
			d[3]++
		case "sw":
			d[4]++
		case "nw":
			d[5]++
		}
	}

	//Clean up opposite
	for i := 0; i < 3; i++ {
		switch {
		case d[i] > d[i+3]:
			d[i] -= d[i+3]
			d[i+3] = 0
		case d[i] < d[i+3]:
			d[i+3] -= d[i]
			d[i] = 0
		case d[i] == d[i+3]:
			d[i] = 0
			d[i+3] = 0
		}
	}

	//Reduce
	for i := 0; i < 5; i++ {
		switch {
		case d[i] == 0 || d[(i+2)%6] == 0:
		case d[i] > d[(i+2)%6]:
			d[i] -= d[(i+2)%6]
			d[i+1] += d[(i+2)%6]
			d[(i+2)%6] = 0
		case d[i] < d[(i+2)%6]:
			d[(i+2)%6] -= d[i]
			d[i+1] += d[i]
			d[i] = 0
		case d[i] == d[(i+2)%6]:
			d[i] = 0
			d[i+1] += d[(i+2)%6]
			d[(i+2)%6] = 0
		}
	}

	distance := 0
	for _, v := range d {
		distance += v
	}

	return distance
}

func task11PartTwo(instructions []string) int {
	//Build path
	//Check distance.
	return 0
}
