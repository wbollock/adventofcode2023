package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	input := readInput("input")
	answer := processInput(input)
	fmt.Println(answer)
}

func readInput(file string) (input []byte) {
	input, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return input
}

func processInput(input []byte) (int) {

	inputSlice := strings.Split(string(input), "\n")
	var numberSlice []int

	// then add up all the numbers
	for _, line := range inputSlice {
		// for each line
		// combine the first and last digit into a number
		var digitSlice []rune
		for _, char := range line {
			if unicode.IsDigit(char) {
			   digitSlice = append(digitSlice, char)
			}
		 }
		 // number will always be first and last element of slice
		 firstDigit := digitSlice[0]
		 secondDigit := digitSlice[len(digitSlice) - 1]
		 resultDigit := string(firstDigit) + string(secondDigit)

		 resultInt, err := strconv.Atoi(resultDigit)
		 if err != nil {
			panic(err)
		}

		 numberSlice = append(numberSlice, resultInt)
	}

	var sum int
	for _, number := range numberSlice {
		sum += number
	}

	return sum
}