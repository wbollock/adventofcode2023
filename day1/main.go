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

func processInput(input []byte) int {

	inputSlice := strings.Split(string(input), "\n")
	var numberSlice []int

	for _, line := range inputSlice {
		// for each line
		stringDigitSlice := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		var digitSlice []string

		end := 1
		digitFound := false
		for i := 0; i < len(line); {

			found := false
			if end > len(line) {
				break
			}
			chunk := line[i:end]
			lastChar := chunk[len(chunk)-1]
			fmt.Println(chunk)

			fmt.Println("i is " + fmt.Sprint(i))
			fmt.Println("end is " + fmt.Sprint(end))
			for _, stringDigit := range stringDigitSlice {

				if strings.Contains(chunk, stringDigit) {
					digitSlice = append(digitSlice, stringDigit)
					found = true
					digitFound = true
					fmt.Println("Found!" + chunk)

					// subtract one for edge case of last letter of a digit being used for other digits
					// if this is a digit though don't subtract by one
					if unicode.IsDigit(rune(line[end-1])) {
						i = end
					} else {
						i = end - 1
					}

					end = i + 1
					break
				} else if unicode.IsDigit(rune(lastChar)) {
					digitSlice = append(digitSlice, string(rune(lastChar)))
					found = true
					digitFound = true
					i = end

					// fmt.Println("index is " + fmt.Sprint(index))
					end = i + 1
					// fmt.Println("end is " + fmt.Sprint(end))
					break
				}
			}

			if !found {
				end++
			}

			// break the loop if the remaining string is too short for any match
			if i >= len(line) {
				break
			}
		}
		if !digitFound {
			continue
		}

		stringToDigit := map[string]string{
			"one":   "1",
			"two":   "2",
			"three": "3",
			"four":  "4",
			"five":  "5",
			"six":   "6",
			"seven": "7",
			"eight": "8",
			"nine":  "9",
		}
		// convert any words in our digitSlice to string of the number
		for i, digit := range digitSlice {
			if num, exists := stringToDigit[digit]; exists {
				digitSlice[i] = num
			}
		}

		// number will always be first and last element of slice
		firstDigit := digitSlice[0]
		secondDigit := digitSlice[len(digitSlice)-1]

		resultDigit := firstDigit + secondDigit

		resultInt, err := strconv.Atoi(resultDigit)
		if err != nil {
			panic(err)
		}

		numberSlice = append(numberSlice, resultInt)
	}

	var sum int
	for i, number := range numberSlice {
		fmt.Printf("Line %v number %v \n", i+1, number)
		sum += number
	}

	return sum
}
