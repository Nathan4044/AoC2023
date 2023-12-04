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
	file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var result int

	for scanner.Scan() {
		_, winningNumbers, ticketNumbers := parseTicket(scanner.Text())

		result += scoreTicket(winningNumbers, ticketNumbers)
	}

	fmt.Printf("%d\n", result)
}

func scoreTicket(winningNumbers, ticketNumbers []int) int {
	winners := make(map[int]bool)
	var matches int

	for _, n := range winningNumbers {
		winners[n] = true
	}

	for _, n := range ticketNumbers {
		_, ok := winners[n]

		if ok {
			matches++
		}
	}

	if matches == 0 {
		return 0
	}

	score := 1

	for i := 0; i < matches-1; i++ {
		score *= 2
	}

	return score
}

func parseTicket(str string) (int, []int, []int) {
	line := str[5:]
	parts := strings.Split(line, ":")

	index, err := strconv.Atoi(strings.Trim(parts[0], " "))
	if err != nil {
		log.Fatal(err)
	}

	listsOfNums := strings.Split(parts[1], "|")
	winningNumbers := numbersFrom(listsOfNums[0])
	ticketNumbers := numbersFrom(listsOfNums[1])

	return index, winningNumbers, ticketNumbers
}

func numbersFrom(str string) []int {
	numbers := []int{}

	for _, n := range strings.Fields(str) {
		num, err := strconv.Atoi(n)

		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, num)
	}

	return numbers
}
