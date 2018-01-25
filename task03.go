package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//Task03 Solution
func Task03() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	number, _ := strconv.Atoi(text)

	// Find length of the number box edge
	edgeLength := int(math.Sqrt(float64(number))) + 1
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
	lv := number
	for _, v := range middleValues {
		diff := int(math.Abs(float64(v - number)))
		if diff < lv {
			lv = diff
		}
	}

	result := sp + lv
	fmt.Printf("Result:\n%d\n", result)
}

func task03PartTwo() {
	/*




	 */
}
