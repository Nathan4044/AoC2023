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
	r := getRace()

	fmt.Printf("%d\n", r.winningMargin())
}

func getRace() race {
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

	timeString := strings.ReplaceAll(strings.Split(lines[0], ": ")[1], " ", "")
	distanceString := strings.ReplaceAll(strings.Split(lines[1], ": ")[1], " ", "")

	time, err := strconv.Atoi(timeString)

	if err != nil {
		log.Fatal(err)
	}

	dist, err := strconv.Atoi(distanceString)

	if err != nil {
		log.Fatal(err)
	}

	return race{
		time:   time,
		record: dist,
	}
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
