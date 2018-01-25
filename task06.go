package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

/*
A debugger program here is having an issue: it is trying to repair a memory reallocation routine,
but it keeps getting stuck in an infinite loop.

In this area, there are sixteen memory banks; each memory bank can hold any number of blocks. The goal
of the reallocation routine is to balance the blocks between the memory banks.

The reallocation routine operates in cycles. In each cycle, it finds the memory bank with the most blocks
(ties won by the lowest-numbered memory bank) and redistributes those blocks among the banks. To do this,
it removes all of the blocks from the selected bank, then moves to the next (by index) memory bank and
inserts one of the blocks. It continues doing this until it runs out of blocks; if it reaches the last
memory bank, it wraps around to the first one.

The debugger would like to know how many redistributions can be done before a blocks-in-banks configuration
is produced that has been seen before.

For example, imagine a scenario with only four memory banks:

The banks start with 0, 2, 7, and 0 blocks. The third bank has the most blocks, so it is chosen for redistribution.
Starting with the next bank (the fourth bank) and then continuing to the first bank, the second bank, and so on,
the 7 blocks are spread out over the memory banks. The fourth, first, and second banks get two blocks each, and
the third bank gets one back. The final result looks like this: 2 4 1 2.

Next, the second bank is chosen because it contains the most blocks (four). Because there are four memory banks,
each gets one block. The result is: 3 1 2 3.

Now, there is a tie between the first and fourth memory banks, both of which have three blocks. The first bank wins
the tie, and its three blocks are distributed evenly over the other three banks, leaving it with none: 0 2 3 4.

The fourth bank is chosen, and its four blocks are distributed such that each of the four banks receives one: 1 3 4 1.

The third bank is chosen, and the same thing happens: 2 4 1 2.

At this point, we've reached a state we've seen before: 2 4 1 2 was already seen. The infinite loop is detected
after the fifth block redistribution cycle, and so the answer in this example is 5.

Given the initial block counts in your puzzle input, how many redistribution cycles must be completed before a
configuration is produced that has been seen before?


*/

//Task06 Solution
func Task06() {
	//Read input.
	input := helpers.InputFile("/input/input06.txt")[0]
	banksString := strings.Split(input, "\t")
	banks := make([]int, len(banksString))
	for i, v := range banksString {
		banks[i], _ = strconv.Atoi(v)
	}

	count := 0
	history := map[string]int{}

	for {
		snapshot := hashString(banks)

		if _, ok := history[snapshot]; ok {
			fmt.Println(count)
			return
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
