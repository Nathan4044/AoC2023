package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	races := getRaces()

	result := 1

	for _, r := range races {
		result *= r.winningMargin()
	}

	fmt.Printf("%d\n", result)
}

func getRaces() []race {
	file, err := os.Open("../input.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	lines := [2]string{}
	index := 0

	for scanner.Scan() {
		lines[index] = scanner.Text()
		index++
	}

	timeStrings := strings.Fields(strings.Split(lines[0], ": ")[1])
	distanceStrings := strings.Fields(strings.Split(lines[1], ": ")[1])

	races := []race{}

	for i := 0; i < len(timeStrings); i++ {
		var r race

		time, err := strconv.Atoi(timeStrings[i])

		if err != nil {
			log.Fatal(err)
		}

		r.time = time

		dist, err := strconv.Atoi(distanceStrings[i])

		if err != nil {
			log.Fatal(err)
		}

		r.record = dist

		races = append(races, r)
	}

	return races
}

type race struct {
	time   int
	record int
}

func (r *race) winningMargin() int {
	var firstVictory int

	for i := 0; i < r.time; i++ {
		distance := i * (r.time - i)

		if distance > r.record {
			firstVictory = i
			break
		}
	}

	failureStates := 2 * (firstVictory)

	return r.time - failureStates + 1
}
