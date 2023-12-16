package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	// input, _ := os.ReadFile("input_test.txt")

	resP1, resP2 := 0, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		rows, cols := []string{}, make([]string, len(strings.Fields(s)[0]))
		for _, s := range strings.Fields(s) {
			rows = append(rows, s)
			for i, r := range s {
				cols[i] += string(r)
			}
		}

		resP1 += mirror(cols, slices.Equal) + 100*mirror(rows, slices.Equal)
		resP2 += mirror(cols, smudge) + 100*mirror(rows, smudge)
	}
	fmt.Println("Result P1: ", resP1)
	fmt.Println("Result P2: ", resP2)
}

func mirror(s []string, equal func([]string, []string) bool) int {
	for i := 1; i < len(s); i++ {
		l := slices.Min([]int{i, len(s) - i})
		a, b := slices.Clone(s[i-l:i]), s[i:i+l]
		slices.Reverse(a)
		//fmt.Println(">>", a, b)
		if equal(a, b) {
			return i
		}
	}
	return 0
}

func smudge(a, b []string) bool {
	diffs := 0
	for i := range a {
		for j := range a[i] {
			// fmt.Println(a[i][j], b[i][j])
			if a[i][j] != b[i][j] {
				diffs++
				// fmt.Println(a[i], b[i])
			}
		}
	}
	return diffs == 1
}
