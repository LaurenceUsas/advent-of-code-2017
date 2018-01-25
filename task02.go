package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//Task02 Solution
func Task02() {
	//Get absolute path path
	pwd, _ := os.Getwd()
	//Open file content
	file, err := os.Open(pwd + "/input/input02.txt")

	if err != nil {
		fmt.Println("Failed opening file")
		return
	}

	defer file.Close()

	var result int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split line to []string
		stringSlice := strings.Split(line, "\t")
		// Covert to []int
		output := make([]int, len(stringSlice))
		for i := range stringSlice {
			output[i], _ = strconv.Atoi(string(stringSlice[i]))
		}

		sort.Ints(output)

		diff := output[len(output)-1] - output[0]

		result += diff
	}
	fmt.Printf("Result:\n%d", result)

}
