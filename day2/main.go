package main

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type GameData struct {
	GameNumber []int
	Reds []int
	Blues []int
	Greens []int
}

func main() {

	elfRed := 12
	elfGreen := 13
	elfBlue := 14
	input := readInput("input")
	sumGames := processInput(input, elfRed, elfBlue, elfGreen)
	


	
	// sum := compareElfGames(gameData)
	slog.Info("Answer", "sum", fmt.Sprint(sumGames))
    // for _, number := range gameData.GameNumber {
    //     slog.Info("Game Numbers", "game", fmt.Sprint(number))
    // }
    // for _, red := range gameData.Reds {
    //     slog.Info("Reds", "red", fmt.Sprint(red))
    // }
    // for _, blue := range gameData.Blues {
	// 	slog.Info("Blues", "blue", fmt.Sprint(blue))
    // }
    // for _, green := range gameData.Greens {
    //     slog.Info("Greens", "green", fmt.Sprint(green))
    // }
}


func readInput(file string) (input []byte) {
	input, err := os.ReadFile(file)
	if err != nil {
		slog.Error(fmt.Sprint(err))
	}

	return input
}

func processInput(input []byte, elfRed int, elfBlue int, elfGreen int) (int) {
	// want a data structure of game, colors
	inputs := strings.Split(string(input), "\n")

	var gamedata GameData
	var sumGames int

	for i, line := range inputs {
		// don't parse Game X just use the loop number
		gamedata.GameNumber = append(gamedata.GameNumber, i + 1)
		gameNum := i + 1

		games := strings.Split(line, ";")

		var redSlice []int
		var blueSlice []int
		var greenSlice []int

		redsGamePossible := false
		bluesGamePossible := false
		greensGamePossible := false
		for _, game := range games {
			// e.g 3 blue, 2 green
			colors := strings.Split(game, ",")
			redRegex := regexp.MustCompile(`(\d+)\sred`)
			blueRegex := regexp.MustCompile(`(\d+)\sblue`)
			greenRegex := regexp.MustCompile(`(\d+)\sgreen`)
			for _, color := range colors {
				redMatch := redRegex.FindStringSubmatch(color)
				blueMatch := blueRegex.FindStringSubmatch(color)
				greenMatch := greenRegex.FindStringSubmatch(color)

				if redMatch != nil {
					reds, err := strconv.Atoi(redMatch[1])
					if err != nil {
						slog.Error(fmt.Sprint(err))
					}
					redSlice = append(redSlice, reds)
				} else {
					redSlice = append(redSlice, 0)
				}

				if blueMatch != nil {
					blues, err := strconv.Atoi(blueMatch[1])
					if err != nil {
						slog.Error(fmt.Sprint(err))
					}
					blueSlice = append(blueSlice, blues)
				} else {
					blueSlice = append(blueSlice, 0)
				}

				if greenMatch != nil {
					greens, err := strconv.Atoi(greenMatch[1])
					if err != nil {
						slog.Error(fmt.Sprint(err))
					}
					greenSlice = append(greenSlice, greens)
				} else {
					greenSlice = append(greenSlice, 0)
				}
			}
		}
		// var rSum int
		// var bSum int
		// var gSum int
		
		for j := range redSlice {
			if redSlice[j] <= elfRed {
				redsGamePossible = true
			} else {
				redsGamePossible = false
				slog.Info("Error! no match!", "game", gameNum, "match", j, "reds", redSlice[j])
				break
			}
			if blueSlice[j] <= elfBlue {
				bluesGamePossible = true
			} else {
				bluesGamePossible = false
				slog.Info("Error! no match!", "game", gameNum, "match", j,  "blues", blueSlice[j], "elfBlue", elfBlue)
				break
			}
			if greenSlice[j] <= elfGreen {
				greensGamePossible = true
			} else {
				greensGamePossible = false
				slog.Info("Error! no match!", "game", gameNum, "match", j,  "greens", greenSlice[j])
				break
			}
		}

		if redsGamePossible && bluesGamePossible && greensGamePossible {
			sumGames = sumGames + i + 1
			slog.Info("Match Found!", "game", i + 1, "sumGames", sumGames)
		}

		
		// gamedata.Reds = append(gamedata.Reds, rSum)
		// gamedata.Blues = append(gamedata.Blues, bSum)
		// gamedata.Greens = append(gamedata.Greens, gSum)
	}

	return sumGames
}

// func compareElfGames(gamedata GameData) (sumGames int) {


// 	for i, number := range gamedata.GameNumber {
// 		redsGamePossible := false
// 		bluesGamePossible := false
// 		greensGamePossible := false
// 		if gamedata.Reds[i] <= elfRed {
// 			redsGamePossible = true
// 		}
// 		if gamedata.Blues[i] <= elfBlue {
// 			bluesGamePossible = true
// 		}
// 		if gamedata.Greens[i] <= elfGreen {
// 			greensGamePossible = true
// 		}

// 		if redsGamePossible && bluesGamePossible && greensGamePossible {
// 			slog.Info("Match Found!", "game", number, "reds", gamedata.Reds[i], "blues", gamedata.Blues[i], "greens", gamedata.Greens[i])
// 			sumGames+=number
// 		}
//     }


// 	return sumGames
// }