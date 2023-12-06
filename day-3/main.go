package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input, _ = os.ReadFile("input.txt")

func isNum(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return true
}

func isSymbol(s string) bool {
	return !isNum(s) && s != "."
}

func isStar(s string) bool {
	return s == "*"
}

func main() {
	fmt.Println(partOne(string(input)))
	fmt.Println(partTwo(string(input)))
}

func partOne(fileInput string) int {
	matrix := buildMatrix(fileInput)
	var adjNumbers []string

	for y, line := range matrix {
		if len(line) == 0 {
			continue
		}

		var currNum []string
		var isCurrAdj bool

		for x, char := range line {
			_, err := strconv.Atoi(char)
			if err != nil {
				if len(currNum) != 0 && isCurrAdj {
					adjNumbers = append(adjNumbers, strings.Join(currNum, ""))
				}
				isCurrAdj = false
				currNum = []string{}
				continue
			}
			currNum = append(currNum, char)

			if isAdjTo(x, y, matrix, isSymbol) {
				isCurrAdj = true
			}
		}

		if len(currNum) != 0 && isCurrAdj {
			adjNumbers = append(adjNumbers, strings.Join(currNum, ""))
		}
	}

	var sum int
	for _, n := range adjNumbers {
		num, _ := strconv.Atoi(n)
		sum += num
	}
	return sum
}

func partTwo(fileInput string) int {
	matrix := buildMatrix(fileInput)
	var validGears = make(map[string][]string)

	for y, line := range matrix {
		if len(line) == 0 {
			continue
		}

		var currNum []string
		var isCurrAdj bool
		var adjStars = make(map[string]bool)

		for x, char := range line {
			_, err := strconv.Atoi(char)
			if err != nil {
				if len(currNum) != 0 && isCurrAdj {
					for starPoint := range adjStars {
						validGears[starPoint] = append(validGears[starPoint], strings.Join(currNum, ""))
					}
				}
				adjStars = make(map[string]bool)
				isCurrAdj = false
				currNum = []string{}
				continue
			}
			currNum = append(currNum, char)

			if isAdjTo(x, y, matrix, isSymbol) {
				isCurrAdj = true
			}

			for _, starPoint := range getAdjStars(x, y, matrix) {
				adjStars[starPoint] = true
			}
		}

		if len(currNum) != 0 && isCurrAdj {
			for starPoint := range adjStars {
				validGears[starPoint] = append(validGears[starPoint], strings.Join(currNum, ""))
			}
		}
	}

	var sum int

	for _, values := range validGears {
		if len(values) == 2 {
			v1, _ := strconv.Atoi(values[0])
			v2, _ := strconv.Atoi(values[1])
			sum = sum + (v1 * v2)
		}
	}

	return sum
}

// Will use a `matcher` function to determine if the adjacent point is a symbol/number/star.
func isAdjTo(x int, y int, matrix [][]string, matcher func(string) bool) bool {
	adjChars := getAdjChars(x, y)

	for i, pointsSet := range adjChars {
		xn, yn := pointsSet[0], pointsSet[1]
		if xn < 0 || yn < 0 || yn > len(matrix)-1 || xn > len(matrix[i])-1 {
			continue
		}
		v := matrix[yn][xn]
		if matcher(v) {
			return true
		}
	}

	return false
}

func getAdjStars(x int, y int, matrix [][]string) []string {
	var stars []string
	adjChars := getAdjChars(x, y)

	for i, pointsSet := range adjChars {
		xn, yn := pointsSet[0], pointsSet[1]
		if xn < 0 || yn < 0 || yn > len(matrix)-1 || xn > len(matrix[i])-1 {
			continue
		}
		v := matrix[yn][xn]
		if isStar(v) {
			ys, xs := strconv.Itoa(yn), strconv.Itoa(xn)
			stars = append(stars, strings.Join([]string{xs, ys}, ","))
		}
	}

	return stars
}

func getAdjChars(x int, y int) [][]int {
	return [][]int{
		{x, y + 1},
		{x, y - 1},
		// --
		{x + 1, y},
		{x - 1, y},
		// --
		{x + 1, y + 1},
		{x + 1, y - 1},
		// --
		{x - 1, y + 1},
		{x - 1, y - 1},
	}
}

func buildMatrix(input string) [][]string {
	var matrix [][]string
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		matrix = append(matrix, []string{})
		for _, char := range line {
			matrix[i] = append(matrix[i], string(char))
		}
	}

	return matrix
}
