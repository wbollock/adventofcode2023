package main

import (
	"testing"
)

func TestMainFunction(t *testing.T) {
	expected := 288

	input := readInput("testinput")
	races := parseInput(input)

	result := calculatePoints(races)
	
	if result != expected {
		t.Errorf("got %v; want %v", result, expected)
	}
}
