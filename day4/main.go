package main

import (
	"fmt"
	"log/slog"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {

	input := readInput("input")
	sum := parseInput(input)
	slog.Info("Points", "sum", sum)
}


func readInput(file string) (input []byte) {
	input, err := os.ReadFile(file)
	if err != nil {
		slog.Error(fmt.Sprint(err))
	}

	return input
}

func parseInput(input []byte) (points float64) {


	cardRegex := regexp.MustCompile(`Card \d+: `)

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		// remove Card X: lines we don't need
		sanitizedLine := cardRegex.ReplaceAllString(line, "")
		cardHalves := strings.Split(sanitizedLine, "|")
		
		// cardHalve[0] == legend for winning
		// count matches from there
		winningNums := strings.Split(cardHalves[0], " ")

		potentialNums := strings.Split(cardHalves[1], " ")

		var matchingNums []string
		for _, num := range winningNums {
			for _, potNum := range potentialNums {
				if num != "" && potNum != "" {
					if num == potNum {
						matchingNums = append(matchingNums, num)
					}
				}
			}
		}

		// points is just 2^(amount of winning nums - 1)
		var totalPoints float64
		if len(matchingNums) > 1 {
			totalPoints = math.Pow(2, float64(len(matchingNums) - 1))
		} else if len(matchingNums) == 1 {
			totalPoints = 1
		}

		points += totalPoints

	}

	return points
	
}


