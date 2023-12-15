package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache = make(map[string]int)

func hitCache(key string) bool {
	if _, ok := cache[key]; ok {
		return true
	}
	return false
}

func parseInput(file *os.File) []string {
	lines := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func recursiveArrangements(springs string, groups []int) int {
	key := springs
	for _, group := range groups {
		key += strconv.Itoa(group) + ","
	}
	if hitCache(key) {
		return cache[key]
	}
	if len(springs) == 0 {
		if len(groups) == 0 {
			return 1
		} else {
			return 0
		}
	}
	if strings.HasPrefix(springs, "?") {
		return recursiveArrangements(strings.Replace(springs, "?", ".", 1), groups) +
			recursiveArrangements(strings.Replace(springs, "?", "#", 1), groups)
	}
	if strings.HasPrefix(springs, ".") {
		res := recursiveArrangements(strings.TrimPrefix(springs, "."), groups)
		cache[key] = res
		return res
	}

	if strings.HasPrefix(springs, "#") {
		if len(groups) == 0 {
			cache[key] = 0
			return 0
		}
		if len(springs) < groups[0] {
			cache[key] = 0
			return 0
		}
		if strings.Contains(springs[0:groups[0]], ".") {
			cache[key] = 0
			return 0
		}
		if len(groups) > 1 {
			if len(springs) < groups[0]+1 || string(springs[groups[0]]) == "#" {
				cache[key] = 0
				return 0
			}
			res := recursiveArrangements(springs[groups[0]+1:], groups[1:])
			cache[key] = res
			return res
		} else {
			res := recursiveArrangements(springs[groups[0]:], groups[1:])
			cache[key] = res
			return res
		}
	}

	return 0
}

func main() {
	file, _ := os.Open("input.txt")
	// file, _ := os.Open("input_test.txt")

	parsedInput := parseInput(file)
	sumP1 := 0
	sumP2 := 0

	for _, line := range parsedInput {
		var comb []int
		lineContent := strings.Split(line, " ")
		for _, Scomb := range strings.Split(lineContent[1], ",") {
			conv, _ := strconv.Atoi(Scomb)
			comb = append(comb, conv)
		}
		sumP1 += recursiveArrangements(lineContent[0], comb)

	}
	fmt.Println("Part1:", sumP1)

	//unfold by adding 5 copies
	for _, line := range parsedInput {
		var comb []int
		lineContent := strings.Split(line, " ")
		springs := ""
		for i := 0; i < 5; i++ {
			springs += lineContent[0]
			if i < 4 {
				springs += "?"
			}
		}
		for i := 0; i < 5; i++ {
			for _, sComb := range strings.Split(lineContent[1], ",") {
				conv, _ := strconv.Atoi(sComb)
				comb = append(comb, conv)
			}
		}
		sumP2 += recursiveArrangements(springs, comb)
	}

	fmt.Println("Part2:", sumP2)
}
