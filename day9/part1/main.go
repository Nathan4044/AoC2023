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
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		ints := getValues(scanner.Text())
		nextValue := getNextValue(ints)
		total += nextValue
	}

	fmt.Println(total)
}

func getValues(line string) []int {
	valStrings := strings.Fields(line)

	vals := []int{}

	for _, s := range valStrings {
		num, err := strconv.Atoi(s)

		if err != nil {
			log.Fatal(err)
		}

		vals = append(vals, num)
	}

	return vals
}

func getNextValue(vals []int) int {
	differences := []int{}

	for i := 1; i < len(vals); i++ {
		diff := vals[i] - vals[i-1]
		differences = append(differences, diff)
	}

	if allEqual(differences) {
		return differences[0] + vals[len(vals)-1]
	} else {
		return getNextValue(differences) + vals[len(vals)-1]
	}
}

func allEqual(vals []int) bool {
	for _, n := range vals {
		if n != vals[0] {
			return false
		}
	}

	return true
}
