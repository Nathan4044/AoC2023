package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines := getLines()
	numbers, symbols := parseLines(lines)

	var total int64

	for _, sym := range symbols {
		neighbours := []int64{}

		for _, num := range numbers {
			if num.isNeighbour(sym) {
				neighbours = append(neighbours, num.value)
			}
		}

		if len(neighbours) == 2 {
			total += neighbours[0] * neighbours[1]
		}
	}

	fmt.Printf("%d\n", total)
}

type number struct {
	value int64
	line  int
	start int
	end   int
}

func (n *number) isNeighbour(sym symbol) bool {
	if n.line-1 <= sym.line &&
		sym.line <= n.line+1 &&
		n.start-1 <= sym.pos &&
		sym.pos <= n.end {
		return true
	}

	return false
}

type symbol struct {
	value string
	line  int
	pos   int
}

func parseLines(lines []string) ([]number, []symbol) {
	nums := []number{}
	syms := []symbol{}

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

				nums = append(nums, number{
					value: val,
					start: startPos,
					end:   pos,
					line:  lineNum,
				})
			} else if line[pos] != '.' {
				syms = append(syms, symbol{
					value: string(line[pos]),
					line:  lineNum,
					pos:   pos,
				})

				pos++
			} else {
				pos++
			}
		}
	}

	return nums, syms
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

func isPartNumber(b byte) bool {
	return b != '.' && !isDigit(b)
}
