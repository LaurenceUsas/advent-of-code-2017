package main

import (
	"fmt"
	"os"

	"github.com/LaurenceUsas/advent-of-code-2017/helpers"
)

/*
If the current node is infected, it turns to its right. Otherwise, it turns to its left. (Turning is done in-place; the current node does not change.)

If the current node is clean, it becomes infected. Otherwise, it becomes cleaned. (This is done after the node is considered for the purposes of changing direction.)
The virus carrier moves forward one node in the direction it is facing.

....#.#...##..######.#...
.#####.#.#..#..#.##.#.##.
.#...#####.##....#.###.##
*/

type Vector2D struct {
	x, y int
}

//Task22 Solution
func Task22() {
	//read input.
	pwd, _ := os.Getwd()
	file := helpers.InputFile(pwd + "/input/input22.txt")

	fmt.Printf("[Part 1 answer] - %v\n", task22PartOne(file))
	fmt.Printf("[Part 2 answer] - %v\n", task22PartTwo(file))
}

func task22PartOne(file []string) int {
	grid := map[Vector2D]rune{}
	for y, line := range file {
		for x, v := range line {
			grid[Vector2D{x: x, y: y}] = v
		}
	}

	pos := Vector2D{x: 12, y: 12}
	dir := Vector2D{x: 0, y: -1}

	infected := 0
	for i := 0; i < 10000; i++ {
		switch grid[pos] {
		case '#':
			dir.x, dir.y = -dir.y, dir.x //Turn
			grid[pos] = '.'              //Affect
		default:
			dir.x, dir.y = dir.y, -dir.x
			grid[pos] = '#'
			infected++
		}
		//Move
		pos.x += dir.x
		pos.y += dir.y
	}
	return infected
}

func task22PartTwo(file []string) int {
	grid := map[Vector2D]rune{}
	for y, line := range file {
		for x, v := range line {
			grid[Vector2D{x: x, y: y}] = v
		}
	}

	pos := Vector2D{x: 12, y: 12}
	dir := Vector2D{x: 0, y: -1}

	infected := 0
	for i := 0; i < 10000000; i++ {
		switch grid[pos] {
		case '#':
			dir.x, dir.y = -dir.y, dir.x //Turn
			grid[pos] = 'F'              //Affect
		case 'W':
			grid[pos] = '#'
			infected++
		case 'F':
			dir.x, dir.y = -dir.x, -dir.y //Turn
			grid[pos] = '.'
		default:
			dir.x, dir.y = dir.y, -dir.x
			grid[pos] = 'W'
		}
		//Move
		pos.x += dir.x
		pos.y += dir.y
	}
	return infected
}
