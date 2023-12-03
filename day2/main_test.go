package main

import (
	"testing"
)

func TestMainFunction(t *testing.T) {
	expected := 8

	elfRed := 12
	elfGreen := 13
	elfBlue := 14
	input := readInput("testinput")
	result := processInput(input, elfRed, elfGreen, elfBlue)
	



	// result := processInput(input)
	// fmt.Println(input)
	// if input != nil {
	// 	fmt.Println("ok")
	// }

	if result != expected {
		t.Errorf("got %d; want %d", result, expected)
	}
}

// func TestMainFunctionSpecial(t *testing.T) {
// 	expected := 8

// 	input := readInput("testinputspecial")
// 	gameData := processInput(input)

// 	elfRed := 40
// 	elfGreen := 40
// 	elfBlue := 40
// 	result := compareElfGames(gameData, elfRed, elfGreen, elfBlue)


// 	// result := processInput(input)
// 	// fmt.Println(input)
// 	// if input != nil {
// 	// 	fmt.Println("ok")
// 	// }

// 	if result != expected {
// 		t.Errorf("got %d; want %d", result, expected)
// 	}
// }
// func TestInputSpecial(t *testing.T) {
// 	expected := 499

// 	input := readInput("testinputspecial")
// 	result := processInput(input)

// 	if result != expected {
// 		t.Errorf("got %d; want %d", result, expected)
// 	}
// }
