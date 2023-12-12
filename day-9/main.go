package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 9")

	// input, _ := os.Open("input_test.txt")
	input, _ := os.Open("input.txt")

	output, output2 := readInput(input)

	fmt.Println("Result P1: ", output)
	fmt.Println("Result P2: ", output2)
}

func readInput(input *os.File) (int, int) {

	scanner := bufio.NewScanner(input)
	output := 0
	output2 := 0

	for scanner.Scan() {
		line := scanner.Text()

		digitsStr := strings.Split(line, " ")
		digits := toIntArr(digitsStr)

		// fmt.Println("Line: ", digits)
		diffedDigitsMatrix := diffLoop(digits)
		extrapolated := extrapolate(diffedDigitsMatrix)

		output += extrapolated[0]

		//reverse input to get part 2
		revDigits := reverseArr(digits)
		diffedDigitsMatrix2 := diffLoop(revDigits)
		extrapolated2 := extrapolate(diffedDigitsMatrix2)
		output2 += extrapolated2[0]
	}

	return output, output2
}

func toIntArr(input []string) []int {
	output := make([]int, len(input))

	for i, v := range input {
		output[i], _ = strconv.Atoi(v)
	}

	return output
}

func diffArrSeq(input []int) []int {
	output := make([]int, len(input)-1)

	for i := 0; i < len(input)-1; i++ {
		output[i] = diffNext(input[i], input[i+1])
	}

	return output
}

func diffNext(a int, b int) int {
	return b - a
}

func checkIfAllZero(input []int) bool {
	for _, v := range input {
		if v != 0 {
			return false
		}
	}

	return true
}

func diffLoop(input []int) map[int][]int {
	output := make(map[int][]int)

	i := 1
	output[0] = input
	for !checkIfAllZero(input) {
		input = diffArrSeq(input)
		// fmt.Println(i)
		output[i] = input
		// fmt.Println("Diffed: ", input)
		i++
	}

	return output
}

func extrapolate(input map[int][]int) map[int]int {
	lasts := make(map[int]int)

	for k, v := range input {
		// fmt.Println("Key: ", k, " Value: ", v, "len: ", len(v), "Last: ", v[len(v)-1])
		lasts[k] = v[len(v)-1]
	}

	extrapolatedLasts := make(map[int]int)

	for k, v := range lasts {
		//if k = 0, add k[1] + k[2] + k[3] ... k[len -1]
		for i := k + 1; i < len(lasts); i++ {
			extrapolatedLasts[k] += lasts[i]
		}
		extrapolatedLasts[k] += v
	}

	return extrapolatedLasts
}

func reverseArr(input []int) []int {
	output := make([]int, len(input))

	for i, v := range input {
		output[len(input)-i-1] = v
	}

	return output
}
