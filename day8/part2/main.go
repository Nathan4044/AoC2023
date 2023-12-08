package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	instructions, nodes := getInput()

	currentNodes := []string{}

	for n := range nodes {
		if n[2] == 'A' {
			currentNodes = append(currentNodes, n)
		}
	}

	nodeSteps := []int{}

	for _, n := range currentNodes {
		instructionIndex := 0
		stepCount := 0

		for n[2] != 'Z' {
			switch instructions[instructionIndex] {
			case 'L':
				n = nodes[n].left
			default:
				n = nodes[n].right
			}

			stepCount++

			if instructionIndex == len(instructions)-1 {
				instructionIndex = 0
			} else {
				instructionIndex++
			}
		}

		nodeSteps = append(nodeSteps, stepCount)
	}

	fmt.Println(lowestCommonMultiple(nodeSteps))
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

func lowestCommonMultiple(ints []int) int {
	result := ints[0] * ints[1] / greatestCommonDenominator(ints[0], ints[1])

	for _, n := range ints[2:] {
		result = result * n / greatestCommonDenominator(result, n)
	}

	return result
}

func greatestCommonDenominator(a, b int) int {
	if a < b {
		t := a
		a = b
		b = t
	}

	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}
