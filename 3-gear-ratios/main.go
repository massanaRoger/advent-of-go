package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TextToCheck struct {
	FullText      []string
	CurrentRow    int
	CurrentColumn int
	AboveChars    []rune
	CurrentChars  []rune
	BelowChars    []rune
}

func (t *TextToCheck) GearRatios() (int64, error) {
	totalNumber := int64(0)
	for len(t.CurrentChars) != 0 {
		num, err := t.GetNextNumber()
		if err != nil {
			return 0, err
		}
		totalNumber += num
	}
	return totalNumber, nil
}

func New(text []string) *TextToCheck {
	if len(text) < 2 {
		return nil
	}
	return &TextToCheck{
		FullText:      text,
		CurrentRow:    0,
		CurrentColumn: 0,
		AboveChars:    []rune{},
		CurrentChars:  []rune(text[0]),
		BelowChars:    []rune(text[1]),
	}
}

func (t *TextToCheck) GetNextNumber() (int64, error) {
	finalNumber := ""
	i := t.CurrentColumn
	for i < len(t.CurrentChars) {
		char := t.CurrentChars[i]
		if char >= '0' && char <= '9' {
			in := i
			for in < len(t.CurrentChars) && t.CurrentChars[in] >= '0' && t.CurrentChars[in] <= '9' {
				finalNumber += string(t.CurrentChars[in])
				in += 1
			}
			t.CurrentColumn = i
			if !t.IsSymbolAdjacent() {
				t.CurrentColumn = in
				return t.GetNextNumber()
			} else {
				t.CurrentColumn = in
				num, err := strconv.ParseInt(finalNumber, 0, 64)
				if err != nil {
					return 0, err
				}
				return num, nil
			}
		}
		i += 1
	}
	t.IterateNext()
	if len(t.CurrentChars) == 0 {
		return 0, nil
	}
	return t.GetNextNumber()
}

func (t *TextToCheck) IterateNext() {
	t.CurrentColumn = 0
	if t.CurrentRow == len(t.FullText)-1 {
		t.CurrentChars = []rune{}
	} else if t.CurrentRow == len(t.FullText)-2 {
		t.CurrentRow += 1
		t.AboveChars = []rune(t.FullText[t.CurrentRow-1])
		t.CurrentChars = []rune(t.FullText[t.CurrentRow])
		t.BelowChars = []rune{}
	} else {
		t.CurrentRow += 1
		t.AboveChars = []rune(t.FullText[t.CurrentRow-1])
		t.CurrentChars = []rune(t.FullText[t.CurrentRow])
		t.BelowChars = []rune(t.FullText[t.CurrentRow+1])
	}
}

func (t *TextToCheck) IsSymbolAdjacent() bool {

	indexToCheck := t.CurrentColumn - 1
	if indexToCheck > 0 {
		if t.CheckIndex(indexToCheck) {
			return true
		}
	}
	i := 0
	indexToCheck = t.CurrentColumn
	for indexToCheck < len(t.CurrentChars) {
		if t.CheckIndex(t.CurrentColumn + i) {
			return true
		}
		if t.CurrentChars[t.CurrentColumn+i] == '.' {
			return false
		}
		indexToCheck += 1
		i += 1
	}
	return false
}

func (t *TextToCheck) CheckIndex(index int) bool {
	if len(t.AboveChars) != 0 {
		if !isNumeric(t.AboveChars[index]) {
			return true
		}
	}
	if !isNumeric(t.CurrentChars[index]) {
		return true
	}
	if len(t.BelowChars) != 0 {
		if !isNumeric(t.BelowChars[index]) {
			return true
		}
	}
	return false
}

func isNumeric(char rune) bool {
	if (char >= '0' && char <= '9') || char == '.' {
		return true
	}
	return false
}

func ReadFileLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func main() {
	lines, err := ReadFileLines("input-4.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	t := New(lines)
	res, err := t.GearRatios()

	if err != nil {
		println(err.Error())
	}

	println(res)

}
