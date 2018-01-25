package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

type EnhancmentBook struct {
	RuleList map[string]*Rule
}

func NewEnhancmentBook(input []string) *EnhancmentBook {
	eb := &EnhancmentBook{}

	t2 := [][]int{
		[]int{0, 1, 2, 3}, // unchanged
		[]int{2, 3, 0, 1}, // flip vertical
		[]int{1, 0, 3, 2}, // flip horizontal
		[]int{2, 0, 3, 1}, // rotate 90°
		[]int{1, 3, 0, 2}, // rotate -90°
		[]int{3, 2, 1, 0}, // rotate 180°
	}

	t3 := [][]int{
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, // unchanged
		[]int{6, 7, 8, 3, 4, 5, 0, 1, 2}, // flip vertical
		[]int{2, 1, 0, 5, 4, 3, 8, 7, 6}, // flip horizontal
		[]int{8, 5, 2, 7, 4, 1, 6, 3, 0}, // flip horizontal and rotate 90°
		[]int{6, 3, 0, 7, 4, 1, 8, 5, 2}, // rotate 90°
		[]int{2, 5, 8, 1, 4, 7, 0, 3, 6}, // rotate -90°
		[]int{8, 7, 6, 5, 4, 3, 2, 1, 0}, // rotate 180°
	}

	rl := make(map[string]*Rule)
	for _, v := range input {
		split := strings.Split(v, " => ")

		// Array of New Pattern
		r := &Rule{}
		b := strings.Split(split[1], "/")
		rr := make([][]string, len(b))
		for i, line := range b {
			c := make([]string, len(b))
			for ii, char := range line {
				c[ii] = string(char)
			}
			rr[i] = c
		}
		r.r = rr

		// Generate Flip/Rotate Patterns
		s := strings.Replace(split[0], "/", "", -1)
		switch len(s) {
		case 4:
			for _, v := range t2 {
				hash := Hash(s, v)
				_, has := rl[hash]
				if !has {
					rl[hash] = r
				}
			}
		case 9:
			for _, v := range t3 {
				hash := Hash(s, v)
				_, has := rl[hash]
				if !has {
					rl[hash] = r
				}
			}
		}
		eb.RuleList = rl
	}
	return eb
}

func Hash(s string, p []int) string {
	r := make([]rune, len(s))
	for i, char := range s {
		r[p[i]] = char
	}
	return string(r)
}

type Rule struct {
	r [][]string
}

func Grow(i int, p [][]string, eb *EnhancmentBook) [][]string {
	for i > 0 {
		s := 0            // Pattern Size
		nps := 0          // New Pattern block size
		var np [][]string // New Pattern Array
		switch {
		case len(p)%2 == 0:
			s = len(p) / 2
			nps = 3
		case len(p)%3 == 0:
			s = len(p) / 3
			nps = 4
		}
		np = Make2DArray(s * nps)
		for i := 0; i < s; i++ { // Pattern Row ID
			for ii := 0; ii < s; ii++ { // Pattern Collumn ID
				ptf := "" //Pattern to find
				ptfs := nps - 1
				for a := 0; a < ptfs; a++ {
					for aa := 0; aa < ptfs; aa++ {
						ptf += p[i*ptfs+a][ii*ptfs+aa]
					}
				}
				tt := eb.RuleList[ptf] // Matched Pattern

				for z := 0; z < nps; z++ {
					for zz := 0; zz < nps; zz++ {
						np[i*nps+z][ii*nps+zz] = tt.r[z][zz]
					}
				}
			}
		}
		return np
	}
	return [][]string{}
}

func Make2DArray(s int) [][]string {
	a := make([][]string, s)
	for i := range a {
		a[i] = make([]string, s)
	}
	return a
}

//Task21 Solution
func Task21() {

	pwd, _ := os.Getwd()
	file := helpers.InputFile(pwd + "/input/input21.txt")

	inputPattern := make([][]string, 3)
	inputPattern = [][]string{
		{".", "#", "."},
		{".", ".", "#"},
		{"#", "#", "#"},
	}

	//Generate ehancment book
	eb := NewEnhancmentBook(file)

	p := Grow(1, inputPattern, eb)

	for i := 0; i < 4; i++ {
		fmt.Println(len(p))
		p = Grow(1, p, eb)
	}

	s := len(p)
	c := 0
	for a := 0; a < s; a++ {
		for aa := 0; aa < s; aa++ {
			if p[a][aa] == "#" {
				c++
			}
		}
	}

	fmt.Printf("[Part 1 answer] - %v\n", c)
}
