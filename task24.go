package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

/*
0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10
*/

type BridgePart struct {
	PortA int
	PortB int
	Value int
}

//Task24 Solution
func Task24() {
	//read input.
	pwd, _ := os.Getwd()
	file := helpers.InputFile(pwd + "/input/input24.txt")
	parts := make([]BridgePart, len(file))
	for i, v := range file {
		split := strings.Split(v, "/")
		a, _ := strconv.Atoi(split[0])
		b, _ := strconv.Atoi(split[1])

		part := BridgePart{
			PortA: a,
			PortB: b,
			Value: a + b,
		}

		parts[i] = part
	}

	fmt.Printf("[Part 1 answer] - %v\n", ConnectPort(0, 0, parts))
	fmt.Printf("[Part 2 answer] - %v\n", 0)
}

var level int

func ConnectPort(port, partValue int, options []BridgePart) int {
	returned := 0
	returnedHighest := 0
	sum := partValue

	found := false
	for i, part := range options {
		switch port {
		case part.PortA:
			found = true
			a := RemoveAndCopy(options, i)
			returnedHighest = ConnectPort(part.PortB, part.Value, a)
		case part.PortB:
			found = true
			a := RemoveAndCopy(options, i)
			returnedHighest = ConnectPort(part.PortA, part.Value, a)
		default:
			continue
		}

		if returned > returnedHighest {
			returnedHighest = returned
		}
	}
	if found == false {
		level++
	}
	// if level%1000000 == 0 {
	// 	fmt.Println(level)
	// }

	return sum + returnedHighest
}

func RemoveAndCopy(s []BridgePart, i int) []BridgePart {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	r := make([]BridgePart, len(s)-1)
	copy(r, s)
	return r[:len(s)-1]
}
