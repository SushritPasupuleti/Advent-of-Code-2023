package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var mappings = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func getCharAtIndex(s string, i int) string {
	return string([]rune(s)[i])
}

func calcResult(input string, isPartTwo bool) int {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		if isPartTwo {
			for key, value := range mappings {
				line = strings.ReplaceAll(line, key, value)
			}
		}
		firstDigitChar := getCharAtIndex(line, strings.IndexAny(line, "123456789"))
		lastDigitChar := getCharAtIndex(line, strings.LastIndexAny(line, "123456789"))

		num := fmt.Sprintf("%s%s", firstDigitChar, lastDigitChar)
		// fmt.Println(num)
		numStr, _ := strconv.Atoi(num)
		sum = sum + numStr
	}

	return sum
}

func main() {

	println(mappings["one"])

	input, _ := os.ReadFile("input.txt")

	sum := calcResult(string(input), false)
	fmt.Println("Part one Result: ", sum)
	sum = calcResult(string(input), true)
	fmt.Println("Part two Result: ", sum)
}
