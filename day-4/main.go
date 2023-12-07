package main

import (
	"fmt"
	"math"
	"os"
	// "regexp"
	// "strconv"
	"slices"
	"strings"
)

type Card struct {
	ID             int
	WinningNumbers []string
	Numbers        []string
	Matches        int
	Points         int
	Count          int
}

func (c Card) String() string {
	return fmt.Sprintf("Card %d: %v | %v \nMatches: %d \nPoints: %d", c.ID, c.WinningNumbers, c.Numbers, c.Matches, c.Points)
}

func main() {

	input, _ := os.ReadFile("input.txt")
	//Format: Card 205: 73 65 24 75 10  8 35 83 78 67 | 89 57 38 93 70  5  9 92 29 55 54 36 37 34 21 40 71 68 33  1 18 80 42 52 72

	cards := buildCards(string(input))

	// fmt.Println("> ", cards)

	sum1 := 0
	for _, card := range cards {
		// fmt.Println("> ", card)
		sum1 += card.Points
	}

	extraCards := 0
	for i := 0; i < len(cards); i++ {
		for j := i + 1; j <= (i + cards[i].Matches); j++ {
			cards[j].Count += cards[i].Count
			extraCards += cards[i].Count
		}
	}
	sum2 := len(cards) + extraCards

	// fmt.Println("----------------------------------------")
	fmt.Println("Part one Result: ", sum1)
	fmt.Println("Part two Result: ", sum2)
}

func buildCards(input string) []Card {
	var cards []Card

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		split := strings.Split(line, ":")[1]
		splitNumbers := strings.Split(split, "|")
		winingNumbersStr := splitNumbers[0]
		numbersStr := splitNumbers[1]

		card := Card{
			ID:             0,
			WinningNumbers: strings.Fields(winingNumbersStr),
			Numbers:        strings.Fields(numbersStr),
			Points:         0,
			Count:          1,
		}

		points := 0

		for _, num := range card.Numbers {
			if slices.Contains(card.WinningNumbers, num) {
				points += 1
			}
			card.Points = score(points)
		}

		card.Matches = points

		for points > 0 {
			points -= 1
		}

		cards = append(cards, card)
	}

	return cards
}

func score(points int) int {
	// 1 if points == 0, doubles points after that
	return int(math.Pow(2, (float64(points - 1))))
}
