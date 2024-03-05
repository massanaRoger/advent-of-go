package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type SymbolInfo struct {
	TimesEncountered int
	Value            int64
}

type TextToCheckSecondPart struct {
	FullText      []string
	CurrentRow    int
	CurrentColumn int
	AboveChars    []rune
	CurrentChars  []rune
	BelowChars    []rune
	Symbols       map[int]SymbolInfo
}

func (t *TextToCheckSecondPart) GearRatios() (int64, error) {
	totalNumber := int64(0)
	for len(t.CurrentChars) != 0 {
		err := t.GetNextNumber()
		if err != nil {
			return 0, err
		}
	}

	for _, value := range t.Symbols {
		if value.TimesEncountered == 2 {
			println(value.Value)
			totalNumber += value.Value
		}
	}

	return totalNumber, nil
}

func New(text []string) *TextToCheckSecondPart {
	if len(text) < 2 {
		return nil
	}
	return &TextToCheckSecondPart{
		FullText:      text,
		CurrentRow:    0,
		CurrentColumn: 0,
		AboveChars:    []rune{},
		CurrentChars:  []rune(text[0]),
		BelowChars:    []rune(text[1]),
		Symbols:       make(map[int]SymbolInfo),
	}
}

func (t *TextToCheckSecondPart) GetNextNumber() error {
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
			num, err := strconv.ParseInt(finalNumber, 0, 64)
			if err != nil {
				return err
			}
			t.AdjacentSymbols(num)
			t.CurrentColumn = in
			return nil
		}
		i += 1
	}
	t.IterateNext()
	return nil
}

func (t *TextToCheckSecondPart) IterateNext() {
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

func (t *TextToCheckSecondPart) AdjacentSymbols(value int64) {

	indexToCheck := t.CurrentColumn - 1
	if indexToCheck > 0 {
		whichRow := t.CheckIndex(indexToCheck)
		if whichRow != -2 {
			symbolIndex := (t.CurrentRow+whichRow)*len(t.CurrentChars) + indexToCheck
			t.addToMap(symbolIndex, value)
			return
		}
	}
	i := 0
	indexToCheck = t.CurrentColumn
	for indexToCheck < len(t.CurrentChars) {
		whichRow := t.CheckIndex(t.CurrentColumn + i)
		if whichRow != -2 {
			symbolIndex := (t.CurrentRow+whichRow)*len(t.CurrentChars) + t.CurrentColumn + i
			t.addToMap(symbolIndex, value)
			return
		}
		if t.CurrentChars[t.CurrentColumn+i] == '.' {
			return
		}
		indexToCheck += 1
		i += 1
	}
}

func (t *TextToCheckSecondPart) CheckIndex(index int) int {
	if len(t.AboveChars) != 0 {
		if !isNumeric(t.AboveChars[index]) {
			return -1
		}
	}
	if !isNumeric(t.CurrentChars[index]) {
		return 0
	}
	if len(t.BelowChars) != 0 {
		if !isNumeric(t.BelowChars[index]) {
			return 1
		}
	}
	return -2
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

func (t *TextToCheckSecondPart) addToMap(key int, value int64) {
	if symbol, ok := t.Symbols[key]; ok {
		symbol.TimesEncountered = symbol.TimesEncountered + 1
		symbol.Value = symbol.Value * value
		t.Symbols[key] = symbol
	} else {
		t.Symbols[key] = SymbolInfo{TimesEncountered: 1, Value: value}
	}
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
