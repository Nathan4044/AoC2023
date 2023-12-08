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
	tickets := parseInput()
	var ticketCount int

	for _, ticket := range tickets[1:] {
		ticket.process(tickets)
		ticketCount += ticket.count
	}

	fmt.Printf("%d\n", ticketCount)
}

type ticket struct {
	id             int
	winningNumbers []int
	ticketNumbers  []int
	count          int
}

func (t *ticket) score() int {
	winners := make(map[int]bool)
	var matches int

	for _, n := range t.winningNumbers {
		winners[n] = true
	}

	for _, n := range t.ticketNumbers {
		_, ok := winners[n]

		if ok {
			matches++
		}
	}

	return matches
}

func (t *ticket) process(tickets []ticket) {
	score := t.score()

	for i := score; i > 0; i-- {
		tickets[t.id+i].count += t.count
	}
}

func parseInput() []ticket {
	file, err := os.Open("../input.txt")
    defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	tickets := []ticket{
		// dummy ticket to shift tickets down 1 place and make maths easier
		{},
	}

	for scanner.Scan() {
		ticket := parseTicket(scanner.Text())
		tickets = append(tickets, ticket)
	}

	return tickets
}

func parseTicket(str string) ticket {
	line := str[5:]
	parts := strings.Split(line, ":")

	index, err := strconv.Atoi(strings.Trim(parts[0], " "))
	if err != nil {
		log.Fatal(err)
	}

	listsOfNums := strings.Split(parts[1], "|")
	winningNumbers := numbersFrom(listsOfNums[0])
	ticketNumbers := numbersFrom(listsOfNums[1])

	return ticket{
		id:             index,
		winningNumbers: winningNumbers,
		ticketNumbers:  ticketNumbers,
		count:          1,
	}
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
