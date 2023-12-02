package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	ID    int
	Red   int
	Green int
	Blue  int
}

func (g Game) String() string {
	return fmt.Sprintf("Game %d: %d red, %d green, %d blue", g.ID, g.Red, g.Green, g.Blue)
}

// Max number of Cubes
var GameConfig = Game{
	ID:    0,
	Red:   12,
	Green: 13,
	Blue:  14,
}

func isPossible(gameSteps []Game) bool {
	for _, game := range gameSteps {
		if game.Red > GameConfig.Red || game.Green > GameConfig.Green || game.Blue > GameConfig.Blue {
			return false
		}
	}
	return true
}

func minCubes(gameSteps []Game) Game {
	var minCube = Game{
		Red:   0,
		Green: 0,
		Blue:  0,
	}

	for _, game := range gameSteps {
		if game.Red > minCube.Red {
			minCube.Red = game.Red
		}
		if game.Green > minCube.Green {
			minCube.Green = game.Green
		}
		if game.Blue > minCube.Blue {
			minCube.Blue = game.Blue
		}
	}
	return minCube

}

func main() {
	//format: Game id: x COLOR1, y COLOR2; z COLOR1, w COLOR2
	input, _ := os.ReadFile("input.txt")

	re := regexp.MustCompile(`(\d+) (\w+)`)

	result := 0
	result2 := 0

	for i, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		if line == "" {
			continue
		}
		// fmt.Println("> ", line)
		var CurrentGame = Game{
			ID:    i + 1,
			Red:   0,
			Green: 0,
			Blue:  0,
		}

		games := []Game{}

		for _, m := range re.FindAllStringSubmatch(line, -1) {
			// print(">>", line)
			n, _ := strconv.Atoi(m[1])
			switch m[2] {
			case "red":
				CurrentGame.Red = n
			case "green":
				CurrentGame.Green = n
			case "blue":
				CurrentGame.Blue = n
			}
			games = append(games, CurrentGame)
		}

		if isPossible(games) {
			result += CurrentGame.ID
		}

		minCube := minCubes(games)

		// fmt.Println(">>", CurrentGame, "minCube", minCube)
		result2 += minCube.Red * minCube.Green * minCube.Blue

	}
	fmt.Println("Part One Result: ", result)
	fmt.Println("Part Two Result: ", result2)
}
