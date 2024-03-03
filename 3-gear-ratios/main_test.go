package main

import (
	"testing"
)

func TestGear(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	textToCheck := New(input)

	result, err := textToCheck.GearRatios()

	if err != nil {
		t.Errorf("Result was expected, got error: %s", err.Error())
	}

	expectedResult := int64(4361)

	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, expectedResult)
	}
}

func TestNextNumber(t *testing.T) {
	text := []string{
		"...*......",
		"..35..633.",
		"..........",
	}
	above := "...*......"
	input := "..35..633."
	below := ".........."

	textToCheck := TextToCheck{
		FullText:      text,
		AboveChars:    []rune(above),
		CurrentChars:  []rune(input),
		BelowChars:    []rune(below),
		CurrentRow:    1,
		CurrentColumn: 2,
	}

	result, _ := textToCheck.GetNextNumber()

	if result != 35 {
		t.Errorf("Result was incorrect, got: %d, want: %d", result, 35)
	}

	text = []string{
		"........*.",
		"..35..633.",
		"..........",
	}
	above = "........*."
	input = "..35..633."
	below = ".........."

	textToCheck = TextToCheck{
		FullText:      text,
		AboveChars:    []rune(above),
		CurrentChars:  []rune(input),
		BelowChars:    []rune(below),
		CurrentRow:    1,
		CurrentColumn: 2,
	}

	result, _ = textToCheck.GetNextNumber()

	if result != 633 {
		t.Errorf("Result was incorrect, got: %d, want: %d", result, 633)
	}
}
func TestSymbolAdjacent(t *testing.T) {
	text := []string{
		"...*......",
		"..35..633.",
		"..........",
	}
	above := "...*......"
	input := "..35..633."
	below := ".........."

	textToCheck := TextToCheck{
		FullText:      text,
		AboveChars:    []rune(above),
		CurrentChars:  []rune(input),
		BelowChars:    []rune(below),
		CurrentRow:    1,
		CurrentColumn: 2,
	}

	result := textToCheck.IsSymbolAdjacent()

	if result != true {
		t.Errorf("Result was incorrect, got: %t, want: %t", result, true)
	}

	textToCheck.CurrentColumn = 6

	result = textToCheck.IsSymbolAdjacent()

	if result != false {
		t.Errorf("Result was incorrect, got: %t, want: %t", result, false)
	}

	text = []string{
		"........*.",
		"..35..633.",
		"..........",
	}
	above = "........*."
	input = "..35..633."
	below = ".........."

	textToCheck = TextToCheck{
		FullText:      text,
		AboveChars:    []rune(above),
		CurrentChars:  []rune(input),
		BelowChars:    []rune(below),
		CurrentRow:    1,
		CurrentColumn: 6,
	}

	result = textToCheck.IsSymbolAdjacent()

	if result != true {
		t.Errorf("Result was incorrect, got: %t, want: %t", result, true)
	}

	text = []string{
		"...*......",
		"..35..633.",
		"......#...",
	}

	above = "...*......"
	input = "..35..633."
	below = "......#..."
	textToCheck = TextToCheck{
		FullText:      text,
		AboveChars:    []rune(above),
		CurrentChars:  []rune(input),
		BelowChars:    []rune(below),
		CurrentRow:    6,
		CurrentColumn: 6,
	}

	result = textToCheck.IsSymbolAdjacent()

	if result != true {
		t.Errorf("Result was incorrect, got: %t, want: %t", result, true)
	}
}

func TestCheckIndex(t *testing.T) {
	above := "...*......"
	input := "..35..633."
	below := ".........."

	textToCheck := TextToCheck{
		AboveChars:   []rune(above),
		CurrentChars: []rune(input),
		BelowChars:   []rune(below),
	}

	result := textToCheck.CheckIndex(3)

	if result != true {
		t.Errorf("Result was incorrect, got: %t, want: %t", result, true)
	}

	result = textToCheck.CheckIndex(4)

	if result != false {
		t.Errorf("Result was incorrect, got: %t, want: %t", result, false)
	}

	result = textToCheck.CheckIndex(7)

	if result != false {
		t.Errorf("Result was incorrect, got: %t, want: %t", result, false)
	}

}
