package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coordinates struct {
	x int
	y int
}

func calcDistances(galaxies map[int]Coordinates) int {
	sum := 0
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			sum += abs(galaxies[i].x-galaxies[j].x) + abs(galaxies[i].y-galaxies[j].y)
		}
	}
	return sum
}

func part2(lines []string) {
	// part1(lines, 1000000)
}

func main() {
	file, _ := os.Open("input.txt")
	// file, _ := os.Open("input_test.txt")

	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	universe := parseMatrix(lines)

	// fmt.Println(universe)

	expandedUniverse := expandMatrix(universe, 2)

	// fmt.Println(expandedUniverse)

	res1 := calcDistances(expandedUniverse)
	fmt.Println("Result part 1: ", res1)

	expandedUniverseP2 := expandMatrix(universe, 1000000)

	res2 := calcDistances(expandedUniverseP2)
	fmt.Println("Result part 2: ", res2)

}

func parseMatrix(lines []string) [][]string {
	matrix := [][]string{}
	for _, line := range lines {
		row := []string{}
		row = append(row, strings.Split(line, "")...)
		matrix = append(matrix, row)
	}
	return matrix
}

func expandMatrix(matrix [][]string, expansionFactor int) map[int]Coordinates {
	toExpandRows := map[int]bool{}
	toExpandCols := map[int]bool{}

	//loop through rows
	for i, row := range matrix {

		//check if row has no galaxies
		if !strings.Contains(strings.Join(row, ""), "#") {
			// fmt.Println("Row: ", row, i)
			// fmt.Println("Row has no galaxies", i)
			toExpandRows[i] = true
		}
	}

	//loop through cols
	length := len(matrix)

	for i := 0; i < length; i++ {
		col := []string{}
		for j := 0; j < length; j++ {
			col = append(col, matrix[j][i])
		}
		if !strings.Contains(strings.Join(col, ""), "#") {
			// fmt.Println("Col: ", col, i)
			// fmt.Println("Col has no galaxies", i)
			toExpandCols[i] = true
		}
	}

	//build map of new coordinates for galaxies
	galaxies := map[int]Coordinates{}
	n := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == "#" {
				nx := 0
				ny := 0
				for k := 0; k < i; k++ {
					if toExpandRows[k] {
						ny++
					}
				}
				for k := 0; k < j; k++ {
					if toExpandCols[k] {
						nx++
					}
				}
				// fmt.Println(nx, ny, i, j)
				galaxies[n] = Coordinates{i + ny*(expansionFactor-1), j + nx*(expansionFactor-1)}
				n++
			}
		}
	}

	// fmt.Println(galaxies)

	return galaxies
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
