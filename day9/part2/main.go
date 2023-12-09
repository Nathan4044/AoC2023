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
		prevValue := getPreviousValue(ints)
		total += prevValue

		fmt.Printf("%d %+v\n", prevValue, ints)
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

func getPreviousValue(vals []int) int {
	differences := []int{}

	for i := 1; i < len(vals); i++ {
		diff := vals[i] - vals[i-1]
		differences = append(differences, diff)
	}

	if allEqual(differences) {
		return vals[0] - differences[0]
	} else {
		return vals[0] - getPreviousValue(differences)
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
