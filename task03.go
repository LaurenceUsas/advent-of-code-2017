package main

import (
	"fmt"
	"math"
)

//Task03 Solution
func Task03() {
	input := 325489

	fmt.Printf("[Part 1 answer] - %v\n", task03PartOne(input))
	fmt.Printf("[Part 2 answer] - %v\n", task03PartTwo(uint(input)))
}

func task03PartOne(input int) int {
	// Find length of the number box edge
	edgeLength := int(math.Sqrt(float64(input))) + 1
	// Make sure its odd
	if edgeLength%2 == 0 {
		edgeLength++
	}

	// Find path 1 - Distance from middle member on outer layer to Centre
	sp := (edgeLength - 1) / 2

	// Find path 2 - Distance to middle member on Outer Layer.
	// calculate middle member values.
	middleValues := []int{}
	first := (edgeLength * edgeLength) - ((edgeLength - 1) / 2)
	middleValues = append(middleValues, first)
	for i := 1; i <= 3; i++ {
		middleValues = append(middleValues, first-i*(edgeLength-1))
	}

	// Find distance to closest member
	lv := input
	for _, v := range middleValues {
		diff := int(math.Abs(float64(v - input)))
		if diff < lv {
			lv = diff
		}
	}

	result := sp + lv
	return result
}

func task03PartTwo(input uint) uint {
	sv := [8][2]int{ //Surrounding values
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	a := [31][31]uint{}
	cp := &Vector2D{ //Current Pos
		x: 16,
		y: 16,
	}
	a[16][16] = 1

	d := &Vector2D{ // Direction Move
		x: 1,
		y: 0,
	}

	st := 1 // straight walk
	w := 0  //walked
	r := 0
	var cr uint //current result
	for i := 1; ; i++ {

		// if done 2 rotate, else, walk straight.
		if w == st {
			r++
			if r == 2 {
				r = 0
				st++
			}
			//Change direction
			d.x, d.y = d.y, -d.x //turn right.
			w = 0
		}
		//Walk
		cp.x, cp.y = cp.x+d.x, cp.y+d.y
		//Calculate new value
		cr = 0
		for _, v := range sv {
			cr += a[cp.x+v[0]][cp.y+v[1]]
		}
		a[cp.x][cp.y] = cr
		w++
		if cr > input {
			return cr
		}
	}
}
