package main

import "testing"

func TestMainFunction(t *testing.T) {
	expected := 281

	input := readInput("testinput")
	result := processInput(input)

	if result != expected {
		t.Errorf("got %d; want %d", result, expected)
	}
}

func TestInputSpecial(t *testing.T) {
	expected := 499

	input := readInput("testinputspecial")
	result := processInput(input)

	if result != expected {
		t.Errorf("got %d; want %d", result, expected)
	}
}
