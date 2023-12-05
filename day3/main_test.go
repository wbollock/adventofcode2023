package main

import (
	"testing"
)

func TestMainFunction(t *testing.T) {
	expected := 4361

	input := readInput("testinput")
	gearmap, partmap := parseInput(input)
	result := findValidParts(gearmap, partmap)

	if result != expected {
		t.Errorf("got %d; want %d", result, expected)
	}
}
