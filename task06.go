package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

//Task06 Solution
func Task06() {
	//Read input.
	pwd, _ := os.Getwd()
	input := helpers.InputFile(pwd + "/input/input06.txt")[0]
	banksString := strings.Split(input, "\t")
	banks := make([]int, len(banksString))
	for i, v := range banksString {
		banks[i], _ = strconv.Atoi(v)
	}

	count := 0
	history := map[string]int{}
	oldSnap := ""
	repeat := 0
	for {
		snapshot := hashString(banks)
		if _, ok := history[snapshot]; ok {
			// Found
			if oldSnap == "" {
				oldSnap = snapshot
			}
			if snapshot == oldSnap {
				repeat++
				if repeat == 1 { // Answer 1
					fmt.Printf("[Part 1 answer] - %v\n", count)
				} else if repeat == 2 { // Answer 2
					fmt.Printf("[Part 2 answer] - %v\n", count-history[snapshot])
					return
				}
			}
		}
		history[snapshot] = count
		banks = redistributeMemory(banks)
		count++
	}
}

func hashString(values []int) string {
	var output string

	for i := 0; i < len(values); i++ {
		output += strconv.Itoa(values[i])
		output += " "
	}
	return output
}

func redistributeMemory(input []int) []int {
	//find index of highest bank.
	hb := 0
	for i, v := range input {
		if v > input[hb] {
			hb = i
		}
	}

	//take everything from that bank and redistribute.
	memory := input[hb]
	input[hb] = 0
	for i := 1; i <= memory; i++ {
		input[(hb+i)%len(input)]++
	}

	return input
}
