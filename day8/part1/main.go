package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	instructions, nodes := getInput()

	currentNode := "AAA"
	instructionIndex := 0
	moveCount := 0

	for currentNode != "ZZZ" {
		switch instructions[instructionIndex] {
		case 'L':
			currentNode = nodes[currentNode].left
		default:
			currentNode = nodes[currentNode].right
		}

		moveCount++

		if instructionIndex == len(instructions)-1 {
			instructionIndex = 0
		} else {
			instructionIndex++
		}
	}

	fmt.Println(moveCount)
}

func getInput() (string, map[string]*node) {
	file, err := os.Open("../input.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()

	nodes := make(map[string]*node)

	for scanner.Scan() {
		line := scanner.Text()

		name := line[:3]
		left := line[7:10]
		right := line[12:15]

		nodes[name] = &node{
			left:  left,
			right: right,
		}
	}

	return instructions, nodes
}

type node struct {
	left  string
	right string
}
