package main

import (
	"fmt"
	"os"
	"strings"
)

type Coordinates struct {
	x, y int
}

type Direction Coordinates

var (
	dirNorth = Direction{-1, 0}
	dirSouth = Direction{1, 0}
	dirWest  = Direction{0, -1}
	dirEast  = Direction{0, 1}
)

// for source direction, the value is the destination direction
var pipes = map[byte]map[Direction]Direction{
	'|': {
		dirNorth: dirNorth,
		dirSouth: dirSouth,
	},
	'-': {
		dirEast: dirEast,
		dirWest: dirWest,
	},
	'L': {
		dirSouth: dirEast,
		dirWest:  dirNorth,
	},
	'J': {
		dirEast:  dirNorth,
		dirSouth: dirWest,
	},
	'7': {
		dirEast:  dirSouth,
		dirNorth: dirWest,
	},
	'F': {
		dirNorth: dirEast,
		dirWest:  dirSouth,
	},
}

func main() {
	file, _ := os.ReadFile("input.txt")
	// file, _ := os.ReadFile("input_test.txt")
	input := string(file)

	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

func part1(input string) any {
	grid := parseInput(input)
	s := findStart(grid)
	loop := findPath(s, grid)

	return len(loop) / 2
}

func part2(input string) any {
	grid := parseInput(input)
	s := findStart(grid)
	loop := findPath(s, grid)

	polyArea := 0
	for i := 0; i < len(loop); i++ {
		curr := loop[i]
		next := loop[(i+1)%len(loop)]

		polyArea += curr.x*next.y - curr.y*next.x
	}

	if polyArea < 0 {
		polyArea = -polyArea
	}
	polyArea /= 2

	return polyArea - len(loop)/2 + 1
}

func parseInput(input string) [][]byte {

	lines := strings.Split(input, "\n")
	grid := make([][]byte, len(lines))

	for i := range lines {
		grid[i] = []byte(lines[i])
	}

	return grid
}

func findStart(grid [][]byte) Coordinates {
	var s Coordinates

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'S' {
				s.x = i
				s.y = j
				return s
			}
		}
	}

	fmt.Println("No start found")
	return s
}

func findPath(s Coordinates, grid [][]byte) []Coordinates {
	for _, pipe := range "|-LJ7F" {

		grid[s.x][s.y] = byte(pipe)
		found := traverse(s, grid)
		if found != nil {
			return found
		}
	}

	fmt.Println("No path found")
	return nil
}

func traverse(s Coordinates, grid [][]byte) []Coordinates {
	cur := s
	dir := getKey(pipes[grid[s.x][s.y]])

	res := []Coordinates{}

	for {
		res = append(res, cur)
		newDir, ok := pipes[grid[cur.x][cur.y]][dir]
		if !ok {
			return nil
		}

		newCoord := Coordinates{cur.x + newDir.x, cur.y + newDir.y}

		if newCoord.x < 0 || newCoord.x >= len(grid) || newCoord.y < 0 || newCoord.y >= len(grid[newCoord.x]) {
			return nil
		}
		if newCoord == s {
			if _, ok := pipes[grid[s.x][s.y]][newDir]; !ok {
				return nil
			}
			break
		}
		cur = newCoord
		dir = newDir
	}

	return res
}

func getKey(m map[Direction]Direction) Direction {
	for k := range m {
		return k
	}

	fmt.Println("No key found")
	return Direction{}
}
