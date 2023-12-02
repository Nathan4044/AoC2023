package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type game struct {
	id    int64
	red   int64
	blue  int64
	green int64
}

func main() {
	games := readInput()

	validGameIds := []int64{}
	var total int64

	for _, game := range games {
		if game.red <= 12 && game.green <= 13 && game.blue <= 14 {
			validGameIds = append(validGameIds, game.id)
			total += game.id
		}
	}

	fmt.Printf("%+v\n", validGameIds)
	fmt.Printf("%d\n", total)
}

func readInput() []game {
	file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	games := []game{}

	for scanner.Scan() {
		game := parseGame(scanner.Text())
		games = append(games, game)
	}

	return games
}

func parseGame(line string) game {
	game := game{}
	firstSpace := strings.Index(line, " ")
	colon := strings.Index(line, ":")

	gameId, err := strconv.ParseInt(line[firstSpace+1:colon], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	game.id = gameId

	gameInfo := strings.Replace(line[colon+1:], ";", ",", -1)
	rounds := strings.Split(gameInfo, ",")

	for _, round := range rounds {
		info := strings.Split(strings.Trim(round, " "), " ")

		colour := info[1]
		count, err := strconv.ParseInt(info[0], 10, 64)

		if err != nil {
			log.Fatalf("error parsing number from %+v: invalid number %s", info, info[0])
		}

		switch colour {
		case "red":
			if count > game.red {
				game.red = count
			}
		case "blue":
			if count > game.blue {
				game.blue = count
			}
		case "green":
			if count > game.green {
				game.green = count
			}
		default:
			log.Fatalf("expected red, green, or blue. got=%s", colour)
		}
	}

	return game
}
