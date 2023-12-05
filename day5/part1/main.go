package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	seeds, mapper := readInput()

	lowestLocation := mapper.run(seeds[0])

	for _, seed := range seeds[1:] {
		loc := mapper.run(seed)

		if loc < lowestLocation {
			lowestLocation = loc
		}
	}

	fmt.Printf("%d\n", lowestLocation)
}

func readInput() ([]int, mapper) {
	file, err := os.ReadFile("../input.txt")

	if err != nil {
		log.Fatalf("readInput: %s\n", err)
	}

	sections := strings.Split(string(file), "\n\n")

	if len(sections) != 8 {
		log.Fatalf("readInput: wrong number of sections retrieved: got %d want %d:\n%+v\n", len(sections), 8, sections)
	}

	seeds := parseSeeds(sections[0])

	seedToSoil := parseMap(sections[1])
	soilToFertilizer := parseMap(sections[2])
	fertilizerToWater := parseMap(sections[3])
	waterToLight := parseMap(sections[4])
	lightToTemperature := parseMap(sections[5])
	temperatureToHumidity := parseMap(sections[6])
	humidityToLocation := parseMap(sections[7])
	mapper := mapper{
		seedToSoil:            seedToSoil,
		soilToFertilizer:      soilToFertilizer,
		fertilizerToWater:     fertilizerToWater,
		waterToLight:          waterToLight,
		lightToTemperature:    lightToTemperature,
		temperatureToHumidity: temperatureToHumidity,
		humidityToLocation:    humidityToLocation,
	}

	return seeds, mapper
}

func parseSeeds(section string) []int {
	str := strings.Split(section, ": ")[1]
	seedStrings := strings.Fields(str)

	seeds := []int{}

	for _, s := range seedStrings {
		i, err := strconv.Atoi(s)

		if err != nil {
			log.Fatal(err)
		}

		seeds = append(seeds, i)
	}

	return seeds
}

func parseMap(section string) mapStage {
	lines := strings.Split(section, "\n")[1:]

	ranges := []*mapRange{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		mapRange, err := parseRange(line)

		if err != nil {
			log.Fatalf("parseMap: error when parsing section:\n\n%s\n\nparsing range: %s\n", section, err)
		}

		ranges = append(ranges, mapRange)
	}

	return mapStage{ranges: ranges}
}

func parseRange(str string) (*mapRange, error) {
	numStrings := strings.Fields(str)

	nums := []int{}

	for _, s := range numStrings {
		n, err := strconv.Atoi(strings.Trim(s, " "))

		if err != nil {
			return nil, err
		}

		nums = append(nums, n)
	}

	if len(nums) != 3 {
		return nil, fmt.Errorf("wrong number of nums from line '%s': got %+v\n", str, nums)
	}

	mr := mapRange{
		destination: nums[0],
		source:      nums[1],
		coverage:    nums[2],
	}

	return &mr, nil
}

type mapRange struct {
	destination int
	source      int
	coverage    int
}

func (m *mapRange) maps(n int) (int, bool) {
	offset := n - m.source

	if 0 <= offset && offset < m.coverage {
		return m.destination + offset, true
	}

	return 0, false
}

type mapStage struct {
	ranges []*mapRange
}

func (m *mapStage) convert(n int) int {
	for _, mr := range m.ranges {
		result, ok := mr.maps(n)

		if ok {
			return result
		}
	}

	return n
}

type mapper struct {
	seedToSoil            mapStage
	soilToFertilizer      mapStage
	fertilizerToWater     mapStage
	waterToLight          mapStage
	lightToTemperature    mapStage
	temperatureToHumidity mapStage
	humidityToLocation    mapStage
}

func (m *mapper) run(seed int) int {
	soil := m.seedToSoil.convert(seed)
	fertilizer := m.soilToFertilizer.convert(soil)
	water := m.fertilizerToWater.convert(fertilizer)
	light := m.waterToLight.convert(water)
	temp := m.lightToTemperature.convert(light)
	humidity := m.temperatureToHumidity.convert(temp)
	location := m.humidityToLocation.convert(humidity)

	return location
}

type seedInfo struct {
	seed       int
	soil       int
	fertilizer int
	water      int
	light      int
	temp       int
	humidity   int
	location   int
}
