package main

import (
	"fmt"
	"os"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

type Maze struct {
	Cells  [][]rune
	Width  int
	Height int
	StartX int
}

func NewMaze(input []string) *Maze {
	m := &Maze{}
	m.Height = len(input) - 1
	m.Width = len(input[0]) - 1
	m.Cells = make([][]rune, len(input))
	// Generate Array
	for i, line := range input {
		m.Cells[i] = make([]rune, len(line))
		for ii, char := range line {
			// Assign Start
			if i == 0 && char == '|' {
				m.StartX = ii
			}
			m.Cells[i][ii] = char
		}
	}
	return m
}

func (m *Maze) Walk() (string, int) {
	result := ""
	steps := 0
	x, y := m.StartX, 0
	dir := 3 // 0 right, 1 top, 2 left, 3 down
	for x >= 0 && x <= m.Width && y >= 0 && y <= m.Height {
		val := m.Cells[y][x]
		if val == ' ' {
			break
		} else if val == '+' {
			//Change direciton
			if dir == 0 || dir == 2 {
				//moving horizontally. Switch to vertical.
				//Check on both side take the one that is '|'
				if y+1 <= m.Height && m.Cells[y+1][x] == '|' {
					dir = 3
				} else if y-1 >= 0 && m.Cells[y-1][x] == '|' {
					dir = 1
				} else {
					fmt.Println("Error at Corner")
					helpers.PrintTwo(x, y)
				}
			} else {
				if x+1 <= m.Width && m.Cells[y][x+1] == '-' {
					dir = 0
				} else if x-1 >= 0 && m.Cells[y][x-1] == '-' {
					dir = 2
				} else {
					fmt.Println("Error at Corner")
					helpers.PrintTwo(x, y)
				}
			}
		} else if val >= 'A' && val <= 'Z' {
			//Collect Letter
			result += string(val)
		}

		//Move
		switch dir {
		case 0:
			x++
		case 1:
			y--
		case 2:
			x--
		case 3:
			y++
		}
		steps++
	}

	return result, steps
}

//Task19 Solution
func Task19() {
	//read input.
	pwd, _ := os.Getwd()
	file := helpers.InputFile(pwd + "/input/input19.txt")

	m := NewMaze(file)
	result, steps := m.Walk()

	fmt.Printf("[Part 1 answer] - %v\n", result)
	fmt.Printf("[Part 2 answer] - %v\n", steps)
}
