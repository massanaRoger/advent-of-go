package main

import (
	"testing"
)

func TestScratchCards(t *testing.T) {
	input := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	scratchCard, err := New(input)

	if err != nil {
		t.Errorf("Result was expected, got error: %s", err.Error())
	}

	expectedResult := int64(30)

	result, _ := scratchCard.TotalScratchCards()
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, expectedResult)
	}

}

func TestParseInput(t *testing.T) {
	input := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	}

	expectedResult := []Game{{
		WinningNumbers: []int{41, 48, 83, 86, 17},
		CurrentNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53},
	}, {
		WinningNumbers: []int{13, 32, 20, 16, 61},
		CurrentNumbers: []int{61, 30, 68, 82, 17, 32, 24, 19},
	},
	}

	result, err := ParseInput(input)
	if err != nil {
		t.Errorf("Result was expected, got error: %s", err.Error())
	}

	for j, el := range result {
		for i, num := range el.CurrentNumbers {
			if num != expectedResult[j].CurrentNumbers[i] {
				t.Errorf("Results are not equal")
			}
		}
		for i, num := range el.WinningNumbers {
			if num != expectedResult[j].WinningNumbers[i] {
				t.Errorf("Results are not equal")
			}
		}
	}
}
