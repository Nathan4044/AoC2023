package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sections := readInput()
	seedRanges := parseSeeds(sections[0])

	lowestLocation := math.MaxInt

	ch := make(chan int)

	for _, sr := range seedRanges {
		go processSeed(sr, sections[1:], ch)
	}

	for range seedRanges {
		loc := <-ch

		if loc < lowestLocation {
			lowestLocation = loc
		}
	}

	fmt.Printf("%d\n", lowestLocation)
}

func processSeed(r *seedRange, m []string, out chan int) {
	mapper := parseMap(m)
	lowest := mapper.run(r.start)

	for i := 1; i < r.coverage; i++ {
		loc := mapper.run(r.start + i)

		if loc < lowest {
			lowest = loc
		}
	}

	out <- lowest
}

func readInput() []string {
	file, err := os.ReadFile("../input.txt")

	if err != nil {
		log.Fatalf("readInput: %s\n", err)
	}

	sections := strings.Split(string(file), "\n\n")

	if len(sections) != 8 {
		log.Fatalf("readInput: wrong number of sections retrieved: got %d want %d:\n%+v\n", len(sections), 8, sections)
	}

	return sections
}

func parseSeeds(section string) []*seedRange {
	str := strings.Split(section, ": ")[1]
	seedStrings := strings.Fields(str)

	ranges := []*seedRange{}

	for i := 0; i < len(seedStrings)-1; i += 2 {
		start, err := strconv.Atoi(seedStrings[i])

		if err != nil {
			log.Fatal(err)
		}

		rangeOfSeeds, err := strconv.Atoi(seedStrings[i+1])

		if err != nil {
			log.Fatal(err)
		}

		ranges = append(ranges, &seedRange{
			start:    start,
			coverage: rangeOfSeeds,
		})
	}

	return ranges
}

func parseMap(sections []string) *mapper {
	seedToSoil := parseMapStage(sections[0])
	soilToFertilizer := parseMapStage(sections[1])
	fertilizerToWater := parseMapStage(sections[2])
	waterToLight := parseMapStage(sections[3])
	lightToTemperature := parseMapStage(sections[4])
	temperatureToHumidity := parseMapStage(sections[5])
	humidityToLocation := parseMapStage(sections[6])

	mapper := mapper{
		seedToSoil:            seedToSoil,
		soilToFertilizer:      soilToFertilizer,
		fertilizerToWater:     fertilizerToWater,
		waterToLight:          waterToLight,
		lightToTemperature:    lightToTemperature,
		temperatureToHumidity: temperatureToHumidity,
		humidityToLocation:    humidityToLocation,
	}

	return &mapper
}

func parseMapStage(section string) *mapStage {
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

	return &mapStage{ranges: ranges}
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
	seedToSoil            *mapStage
	soilToFertilizer      *mapStage
	fertilizerToWater     *mapStage
	waterToLight          *mapStage
	lightToTemperature    *mapStage
	temperatureToHumidity *mapStage
	humidityToLocation    *mapStage
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

type seedRange struct {
	start    int
	coverage int
}
