package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	// "sync"
	// "sort"
)

type Round struct {
	ID                int
	Hand              []string
	HandRanked        []int
	HandJokered       []int //After J is replaced with highest card
	HandJokeredRanked []int //After J is replaced with highest card but its rank is set to 13, instad of 4
	Rank              int
	Score             int
	HandType          string
	HandTypeJokered   string
}

func (r Round) String() string {
	// return fmt.Sprintf("Round %d: %s %v %v Hand Type: %s", r.ID, r.Hand, r.Score, r.Rank, r.HandType)
	return fmt.Sprintf("{ \n Round %d: \n Hand: %s \n Hand Ranked: %v \n Hand Jokered: %v \n Hand Jokered Ranked: %v \n Score: %d \n Rank: %d \n Hand Type: %s \n Hand Type Jokered: %s \n }", r.ID, r.Hand, r.HandRanked, r.HandJokered, r.HandJokeredRanked, r.Score, r.Rank, r.HandType, r.HandTypeJokered)
}

var CardRankings = map[string]int{
	"A": 1,
	"K": 2,
	"Q": 3,
	"J": 4,
	"T": 5,
	"9": 6,
	"8": 7,
	"7": 8,
	"6": 9,
	"5": 10,
	"4": 11,
	"3": 12,
	"2": 13,
}

var CardRankingsJ = map[string]int{
	"A": 1,
	"K": 2,
	"Q": 3,
	"T": 4,
	"9": 5,
	"8": 6,
	"7": 7,
	"6": 8,
	"5": 9,
	"4": 10,
	"3": 11,
	"2": 12,
	"J": 13,
}

func (r *Round) getCardRankings() []int {
	cardRankings := []int{}
	for _, card := range r.Hand {
		cardRankings = append(cardRankings, CardRankings[card])
	}
	return cardRankings
}

func getCardRankings(hand []string) []int {
	cardRankings := []int{}
	for _, card := range hand {
		cardRankings = append(cardRankings, CardRankings[card])
	}
	return cardRankings
}

func getCardRankingsJ(hand []string) []int {
	cardRankings := []int{}
	for _, card := range hand {
		cardRankings = append(cardRankings, CardRankingsJ[card])
	}
	return cardRankings
}

func getCardRankingsJokered(hand []string) []int {

	cardCounts := make(map[string]int)

	for _, card := range hand {
		if card == "J" {
			continue
		}
		cardCounts[card]++
	}
	fmt.Println("Card Counts: ", cardCounts, hand)

	highestCard := 0
	highestCardCount := 0
	for card, count := range cardCounts {
		fmt.Println("Card: ", card, "Count: ", count)
		if count > highestCardCount {
			highestCardCount = count
			highestCard = CardRankingsJ[card]
		}
	}

	fmt.Println("Highest Card: ", highestCard)

	//replace all jokers with the highest card
	//while the highest card is not joker itself

	cardRankings := []int{}
	for _, card := range hand {
		if card == "J" {
			if highestCard == 0 {
				cardRankings = append(cardRankings, 13)
			} else {
				fmt.Println("Replacing J with: ", highestCard)
				cardRankings = append(cardRankings, highestCard)
			}
		} else {
			cardRankings = append(cardRankings, CardRankingsJ[card])
		}
	}

	fmt.Println("Card Rankings: ", cardRankings)

	return cardRankings
}

var HandTypes = map[string]int{
	"Five of a Kind":  1, //All 5 cards have same label
	"Four of a Kind":  2, //Four cards have the same label
	"Full House":      3, //Three cards have the same label, the remaining two have the same label
	"Three of a Kind": 4, //Three cards have the same label
	"Two Pair":        5, //Two pairs of cards have the same label EX: 2,2,5,5,7
	"One Pair":        6, //Two cards have the same label
	"High Card":       7, //All cards have different (unique) labels Ex 2,3,4,5,7
}

// for given hand, return the type of hand it is
func getHandType(hand []string) string {
	// handType := ""
	// allFive := false

	cardCounts := make(map[string]int)

	for _, card := range hand {
		cardCounts[card]++
	}

	fmt.Println(">> ", cardCounts)

	if len(cardCounts) == 1 {
		return "Five of a Kind"
	}

	if len(cardCounts) == 2 {
		for _, count := range cardCounts {
			if count == 4 {
				return "Four of a Kind"
			}
			if count == 3 {
				return "Full House"
			}
		}
	}

	if len(cardCounts) == 3 {
		for _, count := range cardCounts {
			if count == 3 {
				return "Three of a Kind"
			}
			if count == 2 {
				return "Two Pair"
			}
		}
	}

	if len(cardCounts) == 4 {
		return "One Pair"
	}

	if len(cardCounts) == 5 {
		return "High Card"
	}

	return "Unknown"
}

func rebuildHand(hand []int) []string {
	rebuiltHand := []string{}
	for _, card := range hand {
		for k, v := range CardRankingsJ {
			if v == card {
				rebuiltHand = append(rebuiltHand, k)
			}
		}
	}
	return rebuiltHand
}

func parseEntries(entry string, index int) Round {

	// hand := handToArr(entry.split(" ")[0])
	// fmt.Println(strings.Split(entry, " ")[0])
	// fmt.Println(strings.Split(entry, " ")[1])

	handArr := strings.Split(strings.Split(entry, " ")[0], "")
	// sortedHandArr := quicksort(strings.Split(strings.Split(entry, " ")[0], "")) //INFO: Seems to not work right unless I do this
	scoreStr := strings.Split(entry, " ")[1]
	score, _ := strconv.Atoi(scoreStr)
	// fmt.Println(entry, handArr)

	// fmt.Println("Jokered Hand: ", getCardRankingsJokered(handArr), getCardRankings(handArr))

	r := Round{
		ID:   index,
		Hand: handArr,
		// HandRanked: sortedHandArr,
		HandRanked:        getCardRankings(handArr),
		HandJokered:       getCardRankingsJokered(handArr),
		HandJokeredRanked: getCardRankingsJ(handArr),
		Score:             score,
		Rank:              index,
		HandType:          getHandType(handArr),
		HandTypeJokered:   getHandType(rebuildHand(getCardRankingsJokered(handArr))),
	}

	return r
}

func main() {

	// file, _ := os.Open("input_test.txt")
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// var wg sync.WaitGroup

	rounds := []Round{}

	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		// fmt.Println("> ", line)

		// wg.Add(1)
		//
		// go func(line string, lineNumber int) {
		// 	defer wg.Done()

		e := parseEntries(line, lineNumber)
		// fmt.Println("> ", e)
		// }(line, lineNumber)
		rounds = append(rounds, e)
	}

	// fmt.Println(">> ", rounds)

	rankedRounds := generateRankings(rounds)
	rankedRoundsJokered := generateRankingsJokered(rounds)

	// fmt.Println("Ranked Rounds: ", len(rankedRounds))

	score := calcScore(rankedRounds)
	fmt.Println("--------------------")
	scoreJokered := calcScore(rankedRoundsJokered)

	fmt.Println("Score: ", score)
	fmt.Println("Score Jokered: ", scoreJokered)
}

func generateRankings(rounds []Round) []Round {

	var rankings []Round // rank -> id

	//groups rounds by hand type
	handGroups := make(map[string][]Round) // handType -> []Round
	//Add Hand Types to map
	for handType := range HandTypes {
		handGroups[handType] = []Round{}
	}

	//sort all by hand type
	for i, round := range rounds {
		handGroups[round.HandType] = append(handGroups[round.HandType], rounds[i])
	}

	for handType, rounds := range handGroups {

		handTypeRanks := rankByCardValue(rounds)
		handGroups[handType] = handTypeRanks
		// rankings[HandTypes[handType]] = handTypeRanks
		// rankings = append(rankings, handTypeRanks...)
	}

	rankings = append(rankings, handGroups["Five of a Kind"]...)
	rankings = append(rankings, handGroups["Four of a Kind"]...)
	rankings = append(rankings, handGroups["Full House"]...)
	rankings = append(rankings, handGroups["Three of a Kind"]...)
	rankings = append(rankings, handGroups["Two Pair"]...)
	rankings = append(rankings, handGroups["One Pair"]...)
	rankings = append(rankings, handGroups["High Card"]...)

	// fmt.Println("-------------------- PRINTING RANKINGS -------------------")
	// for i, round := range rankings {
	// 	fmt.Println(">> ", i, "->", round)
	// }

	return rankings
}

func generateRankingsJokered(rounds []Round) []Round {

	var rankings []Round // rank -> id

	//groups rounds by hand type
	handGroups := make(map[string][]Round) // handType -> []Round
	//Add Hand Types to map
	for handType := range HandTypes {
		handGroups[handType] = []Round{}
	}

	//sort all by hand type
	for i, round := range rounds {
		handGroups[round.HandTypeJokered] = append(handGroups[round.HandTypeJokered], rounds[i])
	}

	for handType, rounds := range handGroups {

		handTypeRanks := rankByCardValueJokered(rounds)
		handGroups[handType] = handTypeRanks
		// rankings[HandTypes[handType]] = handTypeRanks
		// rankings = append(rankings, handTypeRanks...)
	}

	rankings = append(rankings, handGroups["Five of a Kind"]...)
	rankings = append(rankings, handGroups["Four of a Kind"]...)
	rankings = append(rankings, handGroups["Full House"]...)
	rankings = append(rankings, handGroups["Three of a Kind"]...)
	rankings = append(rankings, handGroups["Two Pair"]...)
	rankings = append(rankings, handGroups["One Pair"]...)
	rankings = append(rankings, handGroups["High Card"]...)

	// fmt.Println("-------------------- PRINTING RANKINGS -------------------")
	// for i, round := range rankings {
	// fmt.Println(">> ", i, "->", round)
	// }

	return rankings
}

// Ranks rounds by card value
func rankByCardValue(rounds []Round) []Round {
	// rankings := make(map[int]int) // rank -> id
	// cardRankings := make(map[int]Round)
	// rankedCards := []Round{}
	// fmt.Println("-------------------- PRINTING CARDs BEFORE RANKINGS -------------------")
	// fmt.Println(">> ", rounds)

	sort.Slice(rounds, func(i, j int) bool {
		for x := range rounds[i].HandRanked {
			if rounds[i].HandRanked[x] == rounds[j].HandRanked[x] {
				continue
			}
			return rounds[i].HandRanked[x] < rounds[j].HandRanked[x]
		}
		return false
	})

	// fmt.Println("-------------------- PRINTING CARD RANKINGS -------------------")
	// fmt.Println(">> Rankings By Card Value", rankings)
	// fmt.Println(">> Rankings By Card Value", cardRankings)
	// fmt.Println(">> Rankings By Card Value", rounds)

	return rounds
}

func rankByCardValueJokered(rounds []Round) []Round {
	// rankings := make(map[int]int) // rank -> id
	// cardRankings := make(map[int]Round)
	// rankedCards := []Round{}
	// fmt.Println("-------------------- PRINTING CARDs BEFORE RANKINGS -------------------")
	// fmt.Println(">> ", rounds)

	sort.Slice(rounds, func(i, j int) bool {
		for x := range rounds[i].HandJokeredRanked {
			if rounds[i].HandJokeredRanked[x] == rounds[j].HandJokeredRanked[x] {
				continue
			}
			return rounds[i].HandJokeredRanked[x] < rounds[j].HandJokeredRanked[x]
		}
		return false
	})

	return rounds
}

func calcScore(rounds []Round) int {
	score := 0
	total := len(rounds)
	for i, round := range rounds {
		// for i, round := range rounds {
		// fmt.Println("Score: ", (total - i), "*", round.Score, " --- ", round.Hand, round.HandType, round.HandTypeJokered)
		score += (total - i) * round.Score
	}
	return score
}
