package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines := getLines()
	numbers := getNumbers(lines)

	var total int64

	for _, number := range numbers {
		if number.isPartNumber() {
			total += number.value
		}
	}

	fmt.Printf("%d\n", total)
}

type number struct {
	value      int64
	line       int
	start      int
	end        int
	neighbours string
}

func (n *number) isPartNumber() bool {
	for i := 0; i < len(n.neighbours); i++ {
		if n.neighbours[i] != '.' && !isDigit(n.neighbours[i]) {
			return true
		}
	}

	return false
}

func getNumbers(lines []string) []number {
	nums := []number{}
	var pos int
	var startPos int

	for lineNum, line := range lines {
		pos = 0
		for pos < len(line) {
			if isDigit(line[pos]) {
				startPos = pos

				for pos < len(line) && isDigit(line[pos]) {
					pos++
				}

				val, err := strconv.ParseInt(line[startPos:pos], 10, 64)

				if err != nil {
					log.Fatalf("error parsing number '%s' from line %d, starting at position %d\n",
						line[startPos:pos], lineNum, startPos)
				}

				var neighbours bytes.Buffer

				if startPos > 0 {
					neighbours.WriteByte(line[startPos-1])
				}

				if pos < len(line) {
					neighbours.WriteByte(line[pos])
				}

				neighbourStart := startPos - 1

				if neighbourStart < 0 {
					neighbourStart = 0
				}

				neighbourEnd := pos + 1

				if neighbourEnd > len(line)-1 {
					neighbourEnd = len(line) - 1
				}

				if lineNum > 0 {
					neighbours.WriteString(lines[lineNum-1][neighbourStart:neighbourEnd])
				}

				if lineNum < len(lines)-1 {
					neighbours.WriteString(lines[lineNum+1][neighbourStart:neighbourEnd])
				}

				nums = append(nums, number{
					value:      val,
					start:      startPos,
					end:        pos,
					line:       lineNum,
					neighbours: neighbours.String(),
				})
			} else {
				pos++
			}
		}
	}

	return nums
}

func getLines() []string {
	content, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := []string{}

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
