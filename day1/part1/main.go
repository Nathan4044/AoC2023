package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	contents := bufio.NewScanner(file)
	var result int64 = 0

	for contents.Scan() {
		number, err := extractNumber(contents.Text())

		if err != nil {
			log.Fatal(err)
		}

		result += number
	}

	fmt.Printf("%d\n", result)
}

func extractNumber(line string) (int64, error) {
	digits := []rune{}

	for _, c := range line {
		if isDigit(c) {
			digits = append(digits, c)
		}
	}

	if len(digits) < 1 {
		return 0, fmt.Errorf("not enough digits found (%s) from line: %s", string(digits), line)
	}

	textResult := string(digits[0]) + string(digits[len(digits)-1])
	result, err := strconv.ParseInt(textResult, 10, 64)

	if err != nil {
		return 0, fmt.Errorf("invalid number: got '%s' from '%s'", textResult, line)
	}

	return result, nil
}

func isDigit(c rune) bool {
	if '0' <= c && c <= '9' {
		return true
	}

	return false
}
