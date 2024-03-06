package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AllScratchCards struct {
	ScratchCards []ScratchCard
}

type ScratchCard struct {
	Game          Game
	numCardBoards int64
}

type Game struct {
	WinningNumbers []int
	CurrentNumbers []int
}

func New(input []string) (*AllScratchCards, error) {
	games, err := ParseInput(input)

	allCards := &AllScratchCards{
		ScratchCards: make([]ScratchCard, 0, len(games)),
	}

	if err != nil {
		return nil, err
	}

	for _, game := range games {
		s := ScratchCard{
			Game:          game,
			numCardBoards: 1,
		}
		allCards.ScratchCards = append(allCards.ScratchCards, s)
	}
	return allCards, nil
}

func (a *AllScratchCards) TotalScratchCards() (int64, error) {
	totalCards := int64(0)
	for i, card := range a.ScratchCards {
		gameResult := CalculateGameResult(card.Game)
		for j := int64(0); j < a.ScratchCards[i].numCardBoards; j++ {
			for k := 1; k <= gameResult; k++ {
				finalIndex := k + i
				if finalIndex < len(a.ScratchCards) {
					a.ScratchCards[finalIndex].numCardBoards += 1
				}
			}
		}
		totalCards += card.numCardBoards
	}
	return totalCards, nil
}

func CalculateGameResult(game Game) int {
	gameResult := 0
	for _, curNum := range game.CurrentNumbers {
		for _, winNum := range game.WinningNumbers {
			if curNum == winNum {
				gameResult += 1
			}
		}
	}
	return gameResult
}

func ParseInput(input []string) ([]Game, error) {
	games := make([]Game, 0, len(input))

	for _, el := range input {
		splitInput := strings.Split(el, ":")
		// Now we get the part after : to get the relevant info
		gamesString := strings.Split(splitInput[1], "|")
		// Position 0 of gamesString are winning numbers and 1 are the numbers i got
		winningString := strings.Fields(gamesString[0])
		currentString := strings.Fields(gamesString[1])
		gameToAppend := Game{
			WinningNumbers: make([]int, 0, len(winningString)),
			CurrentNumbers: make([]int, 0, len(currentString)),
		}
		for _, str := range winningString {
			num, err := strconv.Atoi(str)
			if err != nil {
				return []Game{}, err
			}
			gameToAppend.WinningNumbers = append(gameToAppend.WinningNumbers, num)
		}
		for _, str := range currentString {
			num, err := strconv.Atoi(str)
			if err != nil {
				return []Game{}, err
			}
			gameToAppend.CurrentNumbers = append(gameToAppend.CurrentNumbers, num)
		}
		games = append(games, gameToAppend)
	}
	return games, nil
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
	t, err := New(lines)

	if err != nil {
		fmt.Println(err)
	}

	res, err := t.TotalScratchCards()

	if err != nil {
		fmt.Println(err)
	}

	println(res)
}
