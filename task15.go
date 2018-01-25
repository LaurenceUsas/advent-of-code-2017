package main

import (
	"fmt"
)

const (
	INPUT_A  = 783
	INPUT_B  = 325
	FACTOR_A = 16807
	FACTOR_B = 48271
	DIVIDER  = 2147483647
)

//Task15 Solution
func Task15() {
	fmt.Printf("[Part 1 answer] - %v\n", task15PartOne())
	fmt.Printf("[Part 2 answer] - %v\n", task15PartTwo())
}

func task15PartOne() int {
	score := 0
	genA := INPUT_A
	genB := INPUT_B

	for i := 0; i < 40000000; i++ {
		genA = generateValue(genA, FACTOR_A)
		genB = generateValue(genB, FACTOR_B)

		if uint16(genA) == uint16(genB) {
			score++
		}
	}
	return score
}

func generateValue(input, f int) int {
	return (input * f) % DIVIDER
}

func task15PartTwo() int {
	score := 0

	ca := generateDivisibleValue(INPUT_A, FACTOR_A, 4, 5000000)
	cb := generateDivisibleValue(INPUT_B, FACTOR_B, 8, 5000000)

	for {
		valA, openA := <-ca
		valB, openB := <-cb

		if !openA || !openB {
			return score
		} else if uint16(valA) == uint16(valB) {
			score++
		}
	}
}

func generateDivisibleValue(input, f, d, i int) <-chan int {
	out := make(chan int, 50)
	go func() {
		v := input
		for ii := 0; ii < i; ii++ {
			v = (v * f) % DIVIDER
			if v%d == 0 {
				out <- v
			}
		}

		close(out)
	}()
	return out
}
