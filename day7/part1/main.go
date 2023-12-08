package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	highCard = iota
	pair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func main() {
	hands := parse_input()

	var handsByRank [7][]*hand

	for _, h := range hands {
		r := h.rank

		handsByRank[r] = append(handsByRank[r], h)
	}

	currentScore := 1
	total := 0

	for _, hands := range handsByRank {
		sortEqualHands(hands)

		for _, h := range hands {
			total += h.bet * currentScore
			currentScore++
		}
	}

	fmt.Printf("%d\n", total)
}

func parse_input() []*hand {
	file, err := os.Open("../input.txt")
    defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	hands := []*hand{}

	for scanner.Scan() {
		h := parseHand(scanner.Text())
		hands = append(hands, h)
	}

	return hands
}

type hand struct {
	cards [5]int
	rank  int
	bet   int
}

func (h *hand) lessThanHandOfSameRank(h2 *hand) bool {
	for i, v := range h.cards {
		if v < h2.cards[i] {
			return true
		}

		if v > h2.cards[i] {
			return false
		}
	}

	log.Fatalf("hands are exactly equal: %+v %+v\n", h, h2)
	return false
}

func parseHand(line string) *hand {
	cardStr := line[:5]
	betStr := line[6:]

	cards := parseCards(cardStr)

	bet, err := strconv.Atoi(betStr)

	if err != nil {
		log.Fatal(err)
	}

	rank := rankHand(cards)

	return &hand{
		cards: cards,
		rank:  rank,
		bet:   bet,
	}
}

func parseCards(s string) [5]int {
	var cards [5]int

	for i, r := range s {
		val, err := strconv.Atoi(string(r))

		if err == nil {
			cards[i] = val
			continue
		}

		switch r {
		case 'A':
			cards[i] = 14
		case 'K':
			cards[i] = 13
		case 'Q':
			cards[i] = 12
		case 'J':
			cards[i] = 11
		case 'T':
			cards[i] = 10
		}
	}

	return cards
}

func rankHand(cards [5]int) int {
	cardCounts := make(map[int]int)

	for _, c := range cards {
		cardCounts[c]++
	}

	switch len(cardCounts) {
	case 1:
		return fiveOfAKind
	case 2:
		for _, v := range cardCounts {
			if v == 1 || v == 4 {
				return fourOfAKind
			}

			return fullHouse
		}
	case 3:
		for _, v := range cardCounts {
			if v == 3 {
				return threeOfAKind
			} else if v == 2 {
				return twoPair
			}
		}
	case 4:
		return pair
	default:
		return highCard
	}

	log.Fatalf("Failed to rank hand: %+v\n", cards)
	return -1
}

func sortEqualHands(hands []*hand) {
	for i := 1; i < len(hands); i++ {
		for i > 0 && hands[i].lessThanHandOfSameRank(hands[i-1]) {
			h := hands[i]
			hands[i] = hands[i-1]
			hands[i-1] = h

			i--
		}
	}
}
