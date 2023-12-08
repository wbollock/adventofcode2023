package main

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {

	raceTimes []int
	raceDistances []int
}

func main() {

	input := readInput("input")
	races := parseInput(input)

	for i, time := range races.raceTimes {
		slog.Info("Race Times", "race", i, "time", time)
	}

	for i, dist := range races.raceDistances {
		slog.Info("Race Distance", "race", i, "dist", dist)
	}

	sum := calculatePoints(races)
	slog.Info("Points", "sum", sum)
}


func readInput(file string) (input []byte) {
	input, err := os.ReadFile(file)
	if err != nil {
		slog.Error(fmt.Sprint(err))
	}

	return input
}

func parseInput(input []byte) (races Race) {

	lines := strings.Split(string(input), "\n")

	timeRegex := regexp.MustCompile(`Time: `)
	sanitizedTime := timeRegex.ReplaceAllString(lines[0], "")

	


	times := strings.Split(sanitizedTime, " ")
	for _, time := range times {
		if time != "" {
			timeInt, err := strconv.Atoi(time)
			if err != nil {
				slog.Error("Error converting string", err)
			}
			races.raceTimes = append(races.raceTimes, timeInt)
		}

	}

	distRegex := regexp.MustCompile(`Distance: `)
	sanitizedDist := distRegex.ReplaceAllString(lines[1], "")

	distances := strings.Split(sanitizedDist, " ")
	for _, dist := range distances {
		if dist != ""{
			timeInt, err := strconv.Atoi(dist)
			if err != nil {
				slog.Error("Error converting string", err)
			}
			races.raceDistances = append(races.raceDistances, timeInt)
		}

	}
	
	return races
}

func calculatePoints(races Race) (points int) {
	
	// for each race
	// for every second iteration in the racetime
	// do buttonTime x (totalTime - buttonTime)
	// if that is > distance add to points total
	var totalRaceCombinations []int
	var raceCombinations int
	for i, time := range races.raceTimes {
		raceCombinations = 0
		for j := 1; j <= time; j++ {
			totalDistance := j * (time - j)
			// slog.Info("Race Distance", "race", i, "distance", totalDistance)
			if totalDistance > races.raceDistances[i] {
				slog.Info("Winning race", "race", i, "distance", totalDistance)
				raceCombinations++
			}
		}
		totalRaceCombinations = append(totalRaceCombinations, raceCombinations)
	}

	points = 1
	for _, combo := range totalRaceCombinations {
		slog.Info("Combo", "combo", combo)
		points *= combo
	}

	return points
}

