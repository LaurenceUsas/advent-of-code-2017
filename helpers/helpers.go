package helpers

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// InputArg displays a message and return whatever cmd line input that ended with new line. No error handling
func InputArg(displayMessage string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the digit sequence: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	return text
}

// InputFile reads file from subdirectory/filename as argument and returns list of strings (memory handicap). No error handling for input.
func InputFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed opening file")
		//Throw error
	}

	defer file.Close()

	output := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	return output
}

// StringToIntArray doeswhat it says. No error handling
func StringToIntArray(input string) []int {
	output := make([]int, len(input))

	for i := range input {
		output[i], _ = strconv.Atoi(string(input[i]))
	}
	return output
}

func StringArrayToIntArray(input []string) []int {
	output := make([]int, len(input))

	for i := range input {
		output[i], _ = strconv.Atoi(string(input[i]))
	}
	return output
}

// SplitByTabs returns array of strings split by "\t"
func SplitByTabs(input string) []string {
	return strings.Split(input, "\t")
}

// SplitBySpace returns array of strings split by " "
func SplitBySpace(input string) []string {
	return strings.Split(input, " ")
}

// CheckErr throws panic if err != nil
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

//PrintTwo values in formar [value1][value2]
func PrintTwo(a, b interface{}) {
	printTwoTypeString(a)
	printTwoTypeString(b)
	fmt.Printf("\n")
}

func printTwoTypeString(a interface{}) {
	switch reflect.TypeOf(a).Kind() {
	case reflect.String:
		fmt.Printf("[%s]", a)
	case reflect.Int:
		fmt.Printf("[%d]", a)
	}
}
