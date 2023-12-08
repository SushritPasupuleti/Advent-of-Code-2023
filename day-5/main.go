package main

import (
	"fmt"
	"math"
	"os"
	debug "runtime/debug"
	"strconv"
	"strings"
	"sync"
)

var seeds = []int{}

// mappings := [7]string{"seed-to-soil map:", "soil-to-fertilizer map:", "fertilizer-to-water map:", "water-to-light map:", "light-to-temperature map:", "temperature-to-humidity map:", "humidity-to-location map:"}

type XToSoil struct {
	Name                  string
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
	Mappings              map[int]int
}

func (s XToSoil) String() string {
	return fmt.Sprintf("DestinationRangeStart: %d | SourceRangeStart: %d | RangeLength: %d | Mappings: %v", s.DestinationRangeStart, s.SourceRangeStart, s.RangeLength, s.Mappings)
}

func main() {
	// --- Limit OS Resource Usage ---
	debug.SetMemoryLimit(8096 * 1 << 20) // 8GB
	debug.SetMaxThreads(12)
	// --- Limit OS Resource Usage ---

	input, _ := os.ReadFile("input_test.txt")
	// input, _ := os.ReadFile("input.txt")
	//Format: seeds: 0 212 3445 454
	//				 434 24
	//		  seed-to-soil map:
	//		  10 23 44
	//		  23 34 45
	//
	//		  soil-to-fertilizer map:
	seeds = getSeeds(string(input))

	Run(seeds, string(input))

	rangeSeeds := generateSeedRanges(seeds)

	//concat all arrays in rangeSeeds
	var allSeeds []int
	fmt.Println("> rangeSeeds: ", rangeSeeds)
	for _, seed := range rangeSeeds {
		if len(seed) == 0 {
			continue
		}
		allSeeds = append(allSeeds, seed...)
	}

	fmt.Println("> allSeeds: ", allSeeds)

	Run(allSeeds, string(input))
}

func Run(seeds []int, input string) {

	var wg sync.WaitGroup
	seedToSoilMappings := []XToSoil{}
	soilToFertilizerMappings := []XToSoil{}
	fertilizerToWaterMappings := []XToSoil{}
	waterToLightMappings := []XToSoil{}
	lightToTemperatureMappings := []XToSoil{}
	temperatureToHumidityMappings := []XToSoil{}
	humidityToLocationMappings := []XToSoil{}

	fmt.Println("> seeds: ", seeds)

	for i, line := range strings.Split(string(input), "\n") {
		wg.Add(1)
		if line == "" {
			continue
		}
		go func(i int, line string) {

			if strings.Contains(line, "seed-to-soil map:") {
				seedToSoilMappings = getSoilMappings("seed-to-soil", i, string(input))
			}

			if strings.Contains(line, "soil-to-fertilizer map:") {
				soilToFertilizerMappings = getSoilMappings("soil-to-fertilizer", i, string(input))
			}

			if strings.Contains(line, "fertilizer-to-water map:") {
				fertilizerToWaterMappings = getSoilMappings("fertilizer-to-water", i, string(input))
			}

			if strings.Contains(line, "water-to-light map:") {
				waterToLightMappings = getSoilMappings("water-to-light", i, string(input))
			}

			if strings.Contains(line, "light-to-temperature map:") {
				lightToTemperatureMappings = getSoilMappings("light-to-temperature", i, string(input))
			}

			if strings.Contains(line, "temperature-to-humidity map:") {
				temperatureToHumidityMappings = getSoilMappings("temperature-to-humidity", i, string(input))
			}

			if strings.Contains(line, "humidity-to-location map:") {
				humidityToLocationMappings = getSoilMappings("humidity-to-location", i, string(input))
			}
		}(i, line)
	}

	seedToSoilMappingsList := mergeMappings(seedToSoilMappings)
	soilToFertilizerMappingsList := mergeMappings(soilToFertilizerMappings)
	fertilizerToWaterMappingsList := mergeMappings(fertilizerToWaterMappings)
	waterToLightMappingsList := mergeMappings(waterToLightMappings)
	lightToTemperatureMappingsList := mergeMappings(lightToTemperatureMappings)
	temperatureToHumidityMappingsList := mergeMappings(temperatureToHumidityMappings)
	humidityToLocationMappingsList := mergeMappings(humidityToLocationMappings)

	seedToSoilMappingsGenerated := generateMappings(seeds, seedToSoilMappingsList)

	soilList := []int{}

	for _, soil := range soilToFertilizerMappingsList {
		soilList = append(soilList, soil)
	}

	soilToFertilizerMappingsGenerated := generateMappings(soilList, soilToFertilizerMappingsList)

	ferilizerList := []int{}

	for _, fertilizer := range fertilizerToWaterMappingsList {
		ferilizerList = append(ferilizerList, fertilizer)
	}

	fertilizerToWaterMappingsGenerated := generateMappings(ferilizerList, fertilizerToWaterMappingsList)

	waterList := []int{}

	for _, water := range waterToLightMappingsList {
		waterList = append(waterList, water)
	}

	waterToLightMappingsGenerated := generateMappings(waterList, waterToLightMappingsList)

	lightList := []int{}

	for _, light := range lightToTemperatureMappingsList {
		lightList = append(lightList, light)
	}

	lightToTemperatureMappingsGenerated := generateMappings(lightList, lightToTemperatureMappingsList)

	temperatureList := []int{}

	for _, temperature := range temperatureToHumidityMappingsList {
		temperatureList = append(temperatureList, temperature)
	}

	temperatureToHumidityMappingsGenerated := generateMappings(temperatureList, temperatureToHumidityMappingsList)

	humidityList := []int{}

	for _, humidity := range humidityToLocationMappingsList {
		humidityList = append(humidityList, humidity)
	}

	humidityToLocationMappingsGenerated := generateMappings(humidityList, humidityToLocationMappingsList)

	fmt.Println("> Seed -> Soil -> Fertilizer -> Water -> Light -> Temperature -> Humidity -> Location")

	lowestLocation := int(math.Pow(10, 10))

	for _, seed := range seeds {
		fmt.Println("> ", seed, " -> ", getMappingValue(seed, seedToSoilMappingsGenerated), " -> ", getMappingValue(getMappingValue(seed, seedToSoilMappingsGenerated), soilToFertilizerMappingsGenerated), " -> ", getMappingValue(getMappingValue(getMappingValue(seed, seedToSoilMappingsGenerated), soilToFertilizerMappingsGenerated), fertilizerToWaterMappingsGenerated), " -> ", getMappingValue(getMappingValue(getMappingValue(getMappingValue(seed, seedToSoilMappingsGenerated), soilToFertilizerMappingsGenerated), fertilizerToWaterMappingsGenerated), waterToLightMappingsGenerated), " -> ", getMappingValue(getMappingValue(getMappingValue(getMappingValue(getMappingValue(seed, seedToSoilMappingsGenerated), soilToFertilizerMappingsGenerated), fertilizerToWaterMappingsGenerated), waterToLightMappingsGenerated), lightToTemperatureMappingsGenerated), " -> ", getMappingValue(getMappingValue(getMappingValue(getMappingValue(getMappingValue(getMappingValue(seed, seedToSoilMappingsGenerated), soilToFertilizerMappingsGenerated), fertilizerToWaterMappingsGenerated), waterToLightMappingsGenerated), lightToTemperatureMappingsGenerated), temperatureToHumidityMappingsGenerated), " -> ", getMappingValue(getMappingValue(getMappingValue(getMappingValue(getMappingValue(getMappingValue(getMappingValue(seed, seedToSoilMappingsGenerated), soilToFertilizerMappingsGenerated), fertilizerToWaterMappingsGenerated), waterToLightMappingsGenerated), lightToTemperatureMappingsGenerated), temperatureToHumidityMappingsGenerated), humidityToLocationMappingsGenerated))

		if getMappingValue(getMappingValue(getMappingValue(getMappingValue(getMappingValue(getMappingValue(getMappingValue(seed, seedToSoilMappingsGenerated), soilToFertilizerMappingsGenerated), fertilizerToWaterMappingsGenerated), waterToLightMappingsGenerated), lightToTemperatureMappingsGenerated), temperatureToHumidityMappingsGenerated), humidityToLocationMappingsGenerated) < lowestLocation {
			lowestLocation = getMappingValue(getMappingValue(getMappingValue(getMappingValue(getMappingValue(getMappingValue(getMappingValue(seed, seedToSoilMappingsGenerated), soilToFertilizerMappingsGenerated), fertilizerToWaterMappingsGenerated), waterToLightMappingsGenerated), lightToTemperatureMappingsGenerated), temperatureToHumidityMappingsGenerated), humidityToLocationMappingsGenerated)
		}
	}

	fmt.Println("> Lowest location: ", lowestLocation)

	// fmt.Println("> seed-to-soil", seedToSoilMap)
	// fmt.Println("> soil-to-fertilizer", soilToFertilizerMap)
}

func getSeeds(input string) []int {
	var seeds []int

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		if strings.Contains(line, "seeds:") {
			for _, seed := range strings.Split(line, " ") {
				if seed == "" {
					continue
				}

				seedInt, _ := strconv.Atoi(seed)
				// fmt.Println("> ", seed, " -> ", seedInt)
				seeds = append(seeds, seedInt)
			}
		}
	}

	//remove first element
	seeds = seeds[1:]

	return seeds
}

func getSoilMappings(mode string, index int, input string) []XToSoil {
	var xToSoilMappings []XToSoil

	fmt.Println("> ", fmt.Sprint(mode, " map:"))

	//start from index + 1
	for _, line := range strings.Split(input, "\n")[index+1:] {
		// for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		// // next time map is found, break
		// if strings.Contains(line, "map:") {
		// 	break
		// }

		xToSoilInts := []int{}

		for _, seed := range strings.Split(line, " ") {
			if seed == "" {
				continue
			}

			seedInt, _ := strconv.Atoi(seed)
			xToSoilInts = append(xToSoilInts, seedInt)
		}

		fmt.Println(">> ", xToSoilInts)
		xToSoilMapping := XToSoil{
			Name:                  mode,
			DestinationRangeStart: xToSoilInts[0],
			SourceRangeStart:      xToSoilInts[1],
			RangeLength:           xToSoilInts[2],
		}

		xToSoilMapping.Mappings = createMappings(xToSoilMapping)
		// fmt.Println("> ", createMappings(xToSoilMapping))

		xToSoilMappings = append(xToSoilMappings, xToSoilMapping)
	}

	return xToSoilMappings
}

func createMappings(xToSoil XToSoil) map[int]int {

	mappings := make(map[int]int)

	for i := 0; i < xToSoil.RangeLength; i++ {
		mappings[xToSoil.SourceRangeStart+i] = xToSoil.DestinationRangeStart + i
	}

	return mappings
}

func mergeMappings(mappings []XToSoil) map[int]int {
	mergedMappings := make(map[int]int)

	for _, mapping := range mappings {
		// fmt.Println("> ", mapping.Mappings)
		for k, v := range mapping.Mappings {
			mergedMappings[k] = v
		}
	}

	// fmt.Println("> merged: ", mergedMappings)

	return mergedMappings
}

func generateMappings(source []int, toMap map[int]int) map[int]int {
	mappings := make(map[int]int)

	for _, humidity := range source {
		if toMap[humidity] != 0 {
			// fmt.Println("13> ", humidity, " -> ", humidityToLocationMappings[humidity])
			mappings[humidity] = toMap[humidity]
		} else {
			// fmt.Println("13> ", humidity, " -> ", humidity, " (no mapping found)")
			mappings[humidity] = humidity
		}
	}

	return mappings
}

func getMappingValue(key int, mappings map[int]int) int {
	if mappings[key] != 0 {
		// fmt.Println("14> ", humidity, " -> ", humidityToLocationMappings[humidity])
		return mappings[key]
	}

	// fmt.Println("14> ", humidity, " -> ", humidity, " (no mapping found)")
	return key
}

func generateSeedRanges(seeds []int) [][]int {
	var ranges [][]int
	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		length := seeds[i+1]
		rangeVals := make([]int, length)
		for j := 0; j < length; j++ {
			rangeVals[j] = start + j
		}
		ranges = append(ranges, rangeVals)
	}
	return ranges
}
