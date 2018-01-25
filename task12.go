package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

/*
0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5
*/

var input []string
var idsInGroup = map[int]bool{}

//Task12 Solution
func Task12() {
	input = helpers.InputFile("./Task12/input.txt")
	fmt.Printf("[Part 1 answer] - %v\n", task12PartOne(input, 0))
}

func task12PartOne(input []string, id int) int {
	generateGroup(id)
	return len(idsInGroup)
}

func generateGroup(id int) {
	idsInGroup[id] = true

	l := input[id]             //line
	e := strings.Split(l, " ") //elements

	for _, v := range e[2:] {
		str := strings.Trim(v, ",")
		val, _ := strconv.Atoi(str)
		//add to ID to check.
		if _, ok := idsInGroup[val]; !ok {
			generateGroup(val)
		}
	}
}

func part2(input []string, id int) int {
	for len(input) != 0 {

	}
	generateGroup(id)
	return len(idsInGroup)
}
