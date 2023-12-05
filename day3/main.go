package main

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type GearMap struct {
	StartPosition []int
	Number        []int
	LineNumber    []int
}

// e.g * + $.. len is always 1
type PartMap struct {
	StartPosition []int
	LineNumber    []int
}

func main() {

	input := readInput("input")
	gearmap, partmap := parseInput(input)

	// slog.Info("Logging GearMap values")
	// for i, val := range gearmap.StartPosition {
	// 	slog.Info("GearMap StartPosition", "index", i, "value", val)
	// }
	// for i, val := range gearmap.Number {
	// 	slog.Info("GearMap Number", "index", i, "value", val)
	// }
	// for i, val := range gearmap.LineNumber {
	// 	slog.Info("GearMap  LineNumber", "index", i, "value", val)
	// }

	// slog.Info("Logging PartMap values")
	// for i, val := range partmap.StartPosition {
	// 	slog.Info("PartmapMap StartPosition", "index", i, "value", val)
	// }
	// for i, val := range partmap.LineNumber {
	// 	slog.Info("PartMap LineNumber", "index", i, "value", val)
	// }

	sum := findValidParts(gearmap, partmap)
	slog.Info("Sum", "number", fmt.Sprint(sum))
}

type Engine struct {
	Numbers       []int
	StartPosition []int
}

func readInput(file string) (input []byte) {
	input, err := os.ReadFile(file)
	if err != nil {
		slog.Error(fmt.Sprint(err))
	}

	return input
}

func parseInput(input []byte) (GearMap, PartMap) {

	var gearmap GearMap
	var partmap PartMap

	lines := strings.Split(string(input), "\n")

	numberRegex := regexp.MustCompile(`\d+`)
	charRegex := regexp.MustCompile(`[^0-9.]`)
	for i, line := range lines {

		// first find all the numbers and jot down their position
		numberMatches := numberRegex.FindAllStringIndex(line, -1)

		if numberMatches != nil {
			for _, match := range numberMatches {
				number, err := strconv.Atoi(line[match[0]:match[1]])
				if err != nil {
					slog.Error(fmt.Sprint(err))
				}
				gearmap.Number = append(gearmap.Number, number)
				gearmap.StartPosition = append(gearmap.StartPosition, match[0])
				gearmap.LineNumber = append(gearmap.LineNumber, i)
			}
		}

		charMatches := charRegex.FindAllStringIndex(line, -1)

		if charMatches != nil {
			for _, match := range charMatches {

				partmap.StartPosition = append(partmap.StartPosition, match[0])
				partmap.LineNumber = append(partmap.LineNumber, i)
			}
		}
	}

	return gearmap, partmap
}

func findValidParts(gearmap GearMap, partmap PartMap) int {

	// valid number has a part at the same line, before its start position or after
	// or on the line above or below, -1/+1 its matching position

	var sum int
	// first just analyze same line numbers

	for i := range gearmap.Number {
		if i < len(gearmap.LineNumber) && i < len(gearmap.StartPosition) {
			for j := range partmap.StartPosition {
				if j < len(partmap.LineNumber) {
					gearEndPosition := len(fmt.Sprint(gearmap.Number[i])) + gearmap.StartPosition[i] - 1
	
					// logic for whether the char is on the same line as the number
					if gearmap.LineNumber[i] == partmap.LineNumber[j] {
						if gearmap.StartPosition[i]-1 == partmap.StartPosition[j] || gearEndPosition+1 == partmap.StartPosition[j] {
							slog.Info("Match found! Same line", "number", gearmap.Number[i])
							sum += gearmap.Number[i]
						}
					}
	
					// for whether the char is BELOW
					if gearmap.LineNumber[i] == partmap.LineNumber[j]-1 {
						if (gearmap.StartPosition[i]-1 <= partmap.StartPosition[j] && gearEndPosition+1 >= partmap.StartPosition[j]) {
							slog.Info("Match found! Below number", "number", gearmap.Number[i])
							sum += gearmap.Number[i]
						}
					}
	
					// for whether the char is ABOVE
					if gearmap.LineNumber[i] == partmap.LineNumber[j]+1 {
						if gearmap.StartPosition[i]-1 <= partmap.StartPosition[j] && gearEndPosition+1 >= partmap.StartPosition[j] {
							slog.Info("Match found! Above number", "number", gearmap.Number[i])
							sum += gearmap.Number[i]
						}
					}
				}
			}
		}
	}
	

	return sum
}
