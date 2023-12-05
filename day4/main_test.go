package main

import (
	"testing"
)

func TestMainFunction(t *testing.T) {
	expected := 13

	input := readInput("testinput")
	result := parseInput(input)

	if result != float64(expected) {
		t.Errorf("got %v; want %v", result, expected)
	}
}
