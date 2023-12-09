package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"regexp"
	"strings"
)

func main() {

	// input, _ := os.ReadFile("input.txt")
	input, _ := os.ReadFile("input_test.txt")
	//Format: Time:        53     83     72     88
	//		  Distance:   333   1635   1289   1532

	fmt.Println("Day 6 of Advent of Code 2020")

	timeDistanceMappings := getTimeDistanceMappings(string(input))

	totalWays := 1
	totalWaysPart2 := 0
	for time, distance := range timeDistanceMappings {
		fmt.Println("Time: ", time, "Distance: ", distance)

		x := calcQuadratic(1, time, distance)
		// fmt.Println("x1: ", x1, "x2: ", x2)
		// fmt.Println("Time: ", time, "Distance: ", distance, "Speed: ", x)
		fmt.Println("> ", x)
		totalWays *= x
	}

	// fmt.Println("> ", timeDistanceMappings)

	// for time, distance := range timeDistanceMappings {
	// 	x := calculateTimeTarget(distance, time)
	// 	// fmt.Println("Time: ", time, "Distance: ", distance, "Speed: ", x)
	//
	// 	totalWays *= len(x)
	// }

	timeDistanceMappingsPart2 := getTimeDistanceMappingsPart2(string(input))
	for time, distance := range timeDistanceMappingsPart2 {
		fmt.Println("Time: ", time, "Distance: ", distance)

		x := calcQuadratic(1, time, distance)
		fmt.Println("> ", x)
		totalWaysPart2 += x
	}

	// fmt.Println("> ", timeDistanceMappingsPart2)

	// for time, distance := range timeDistanceMappingsPart2 {
	// 	x := calculateTimeTarget(distance, time)
	// 	// fmt.Println("Time: ", time, "Distance: ", distance, "Speed: ", x)
	//
	// 	totalWaysPart2 += len(x)
	// }

	fmt.Println("Total ways: ", totalWays)
	fmt.Println("Total ways part 2: ", totalWaysPart2)
}

func getTimeDistanceMappingsPart2(input string) map[int]int {
	timeDistance := map[int]int{}

	reTime := regexp.MustCompile(`Time:\s+`)
	reDistance := regexp.MustCompile(`Distance:\s+`)

	timings := []int{}
	distances := []int{}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		if reTime.MatchString(line) {
			line = reTime.ReplaceAllString(line, "")

			line = strings.ReplaceAll(line, " ", "")
			num, _ := strconv.Atoi(line)

			timings = append(timings, num)
			continue
		}

		if reDistance.MatchString(line) {
			line = reDistance.ReplaceAllString(line, "")

			line = strings.ReplaceAll(line, " ", "")
			num, _ := strconv.Atoi(line)

			distances = append(distances, num)
			continue
		}
	}

	for i, time := range timings {
		timeDistance[time] = distances[i]
	}

	return timeDistance
}

func getTimeDistanceMappings(input string) map[int]int {
	timeDistance := map[int]int{}

	reTime := regexp.MustCompile(`Time:\s+`)
	reDistance := regexp.MustCompile(`Distance:\s+`)
	reNumbers := regexp.MustCompile(`(\d+)`)

	timings := []int{}
	distances := []int{}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		if reTime.MatchString(line) {
			line = reTime.ReplaceAllString(line, "")

			numbers := reNumbers.FindAllString(line, -1)

			for _, number := range numbers {
				timing, _ := strconv.Atoi(number)
				timings = append(timings, timing)
			}

			fmt.Println(numbers)
			continue
		}

		if reDistance.MatchString(line) {
			line = reDistance.ReplaceAllString(line, "")

			numbers := reNumbers.FindAllString(line, -1)

			for _, number := range numbers {
				distance, _ := strconv.Atoi(number)
				distances = append(distances, distance)
			}

			fmt.Println(numbers)
			continue
		}
	}

	for i, time := range timings {
		timeDistance[time] = distances[i]
	}

	return timeDistance
}

func getSpeed(distance int, time int) int {
	return distance / time
}

func getDistance(speed int, time int) int {
	return speed * time
}

//Brute force and hence not used
func calculateTimeTarget(minDistance int, minTime int) []int {

	// for given distance, calculate the hold down time such that the distance is reached in the remaining time
	holdDuration := 1
	remainingDuration := 0

	winningHoldDurations := []int{}

	for minTime > holdDuration {
		remainingDuration = minTime - holdDuration

		speed := holdDuration

		distanceCovered := getDistance(speed, remainingDuration)

		fmt.Println("For time: ", minTime, "Hold duration: ", holdDuration, "Remaining duration: ", remainingDuration, "Speed: ", speed, "Distance covered: ", distanceCovered, "Distance Target: ", minDistance)

		if distanceCovered > minDistance {
			winningHoldDurations = append(winningHoldDurations, holdDuration)
		}

		holdDuration += 1
	}

	return winningHoldDurations
}

func calcQuadratic(a int, b int, c int) int {

	// Equation for current problem: x(t-x) -d = 0
	// As distance should be greater than d we assume d = d + 1

	// x = (-b +- sqrt(b^2 - 4ac)) / 2a
	// x = (-b +- sqrt(b^2 - 4ac)) / 2a

	a1 := float64(-a)
	b1 := float64(b)
	c1 := float64(-(c + 1))

	x1 := (-b1 + math.Sqrt(b1*b1-4*a1*c1)) / (2 * a1)
	x2 := (-b1 - math.Sqrt(b1*b1-4*a1*c1)) / (2 * a1)

	res := int(math.Floor(x2) - math.Ceil(x1) + 1)
	fmt.Println("x1: ", x1, "x2: ", x2, "res: ", res)
	return res
}
