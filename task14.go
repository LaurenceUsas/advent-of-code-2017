package main

import (
	"fmt"
	"strconv"
)

//Task14 Solution
func Task14() {
	input := "ljoxqyyw-"

	fmt.Printf("[Part 1 answer] - %v\n", task14PartOne(input))
	fmt.Printf("[Part 2 answer] - %v\n", task14PartTwo(input))
}

func task14PartOne(input string) int {
	var instructions []string
	for i := 0; i < 128; i++ {
		s := input + strconv.Itoa(i)
		instructions = append(instructions, s)
	}
	fmt.Println(instructions)
	kh := NewKnotHash()
	sum := 0
	for _, v := range instructions {
		s := kh.Hash(v)
		sum += HexSum(s)
	}
	return sum
}

// HexSum returns 4 bit 1 sum
func HexSum(input string) int {
	sum := 0
	for _, v := range input {
		switch v {
		case '1', '2', '4', '8':
			sum++
		case '3', '5', '6', '9', 'a', 'c':
			sum += 2
		case '7', 'b', 'd', 'e':
			sum += 3
		case 'f':
			sum += 4
		}
	}
	return sum
}

func task14PartTwo(input string) int {

	return 0
}
