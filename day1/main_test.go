package main

import "testing"

func TestMainFunction(t *testing.T) {
    // expected behavior or output
    expected := 142

	input := readInput("testinput")
    result := processInput(input)

    if result != expected {
        t.Errorf("got %d; want %d", result, expected)
    }
}