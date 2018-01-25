package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

/*
Suppose we instead only had a circular list containing five elements, 0, 1, 2, 3, 4, and were
given input lengths of 3, 4, 1, 5.

The list begins as [0] 1 2 3 4 (where square brackets indicate the current position).
The first length, 3, selects ([0] 1 2) 3 4 (where parentheses indicate the sublist to be reversed).
After reversing that section (0 1 2 into 2 1 0), we get ([2] 1 0) 3 4.
Then, the current position moves forward by the length, 3, plus the skip size, 0: 2 1 0 [3] 4.
Finally, the skip size increases to 1.

The second length, 4, selects a section which wraps: 2 1) 0 ([3] 4.
The sublist 3 4 2 1 is reversed to form 1 2 4 3: 4 3) 0 ([1] 2.
The current position moves forward by the length plus the skip size, a total of 5, causing it not to
move because it wraps around: 4 3 0 [1] 2. The skip size increases to 2.

The third length, 1, selects a sublist of a single element, and so reversing it has no effect.
The current position moves forward by the length (1) plus the skip size (2): 4 [3] 0 1 2.
The skip size increases to 3.

The fourth length, 5, selects every element starting with the second: 4) ([3] 0 1 2.
Reversing this sublist (3 0 1 2 4 into 4 2 1 0 3) produces: 3) ([4] 2 1 0.

Finally, the current position moves forward by 8: 3 4 2 1 [0]. The skip size increases to 4.
In this example, the first two numbers in the list end up being 3 and 4; to check the process,
you can multiply them together to produce 12.

*/

//Task10 Solution
func Task10() {

}

func task10PartOne() {
	input := helpers.InputFile("/input/input10.txt")[0]
	strInstructions := strings.Split(input, ",")

	var instructions []int
	for _, v := range strInstructions {
		i, _ := strconv.Atoi(v)
		instructions = append(instructions, i)
	}

	var array [256]int
	for i := 0; i < 256; i++ {
		array[i] = i
	}

	skip := 0
	pos := 0

	for _, length := range instructions {
		//copy to array to slice
		low := pos % 256
		high := (pos + length) % 256
		fmt.Printf("[%v][%v]\n", low, high)

		if length != 0 {
			sublist := getSlice(low, high, array)
			for i := len(sublist); i == 0; i-- {
				//array[]
			}
		}
		pos = high + skip
		skip++
	}

	//fmt.Println(array[0]+array[1])
	//fmt.Println(array)
}

func getSlice(min, max int, a [256]int) []int {
	if min < max {
		println("returning")
		fmt.Println(a[min:max])
		r := make([]int, len(a[min:max]))
		copy(r, a[min:max])
		return r
	}
	s := append(a[min:], a[:max]...)
	r := make([]int, len(s))
	copy(r, s)
	fmt.Printf("LENGTH [%v]", len(r))
	return r

}
