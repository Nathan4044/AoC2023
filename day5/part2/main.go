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
	mapper := parseMap(sections[1:])

	lowest := math.MaxInt

	for _, sr := range seedRanges {
		ranges := mapper.run(sr)

		for _, r := range ranges {
			if r.start < lowest {
				lowest = r.start
			}
		}
	}

	fmt.Printf("%d\n", lowest)
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

func (m *mapRange) lastSource() int {
	return m.source + m.coverage - 1
}

func (m *mapRange) offset() int {
	return m.destination - m.source
}

func (m *mapRange) maps(ranges []*seedRange) ([]*seedRange, []*seedRange) {
	changed := []*seedRange{}
	unchanged := []*seedRange{}

	for _, sr := range ranges {
		switch {
		case sr.start >= m.source && sr.end() <= m.lastSource():
			changed = append(changed, &seedRange{
				start:    sr.start + m.offset(),
				coverage: sr.coverage,
			})
		case sr.start >= m.source && sr.start <= m.lastSource():
			insideCount := m.lastSource() - sr.start

			innerRange := &seedRange{
				start:    sr.start + m.offset(),
				coverage: insideCount + 1,
			}

			outerRange := &seedRange{
				start:    m.lastSource() + 1,
				coverage: sr.coverage - insideCount - 1,
			}

			changed = append(changed, innerRange)
			unchanged = append(unchanged, outerRange)
		case sr.end() <= m.lastSource() && sr.end() >= m.source:
			insideCount := sr.end() - m.source + 1

			innerRange := &seedRange{
				start:    sr.end() - insideCount + 1 + m.offset(),
				coverage: insideCount,
			}

			outerRange := &seedRange{
				start:    sr.start,
				coverage: sr.coverage - insideCount,
			}

			changed = append(changed, innerRange)
			unchanged = append(unchanged, outerRange)
		default:
			unchanged = append(unchanged, sr)
		}
	}

	return changed, unchanged
}

type mapStage struct {
	ranges []*mapRange
}

func (m *mapStage) convert(sr []*seedRange) []*seedRange {
	changed := []*seedRange{}
	unchanged := sr

	initialCoverage := 0

	for _, r := range sr {
		initialCoverage += r.coverage
	}

	for _, mr := range m.ranges {
		newChanged, newUnchanged := mr.maps(unchanged)
		unchanged = newUnchanged
		changed = append(changed, newChanged...)
	}

	changed = append(changed, unchanged...)

	changedCoverage := 0

	for _, r := range changed {
		changedCoverage += r.coverage
	}

	if changedCoverage != initialCoverage {
		log.Fatal("missing seeds!")
	}

	return changed
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

func (m *mapper) run(sr *seedRange) []*seedRange {
	soil := m.seedToSoil.convert([]*seedRange{sr})
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

func (s *seedRange) end() int {
	return s.start + s.coverage - 1
}
