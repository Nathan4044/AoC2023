package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lines := getLines()
	start := getStart(lines)
	_, length := getPipe(lines, start)

	furthestPoint := (length + 1) / 2
	fmt.Println(furthestPoint)
}

func getPipe(lines []string, start *point) (*point, int) {
	current := start.getNeighbours(lines)[0]
	length := 1

	for current.val != 'S' {
		current = current.nextLocation(lines)
		length++
	}

	return current, length
}

func getStart(lines []string) *point {
	var x, y int

	for i, line := range lines {
		x = strings.IndexByte(line, 'S')

		if x > -1 {
			y = i
			break
		}
	}

	return &point{
		x: x,
		y: y,
	}
}

func getLines() []string {
	contents, err := os.ReadFile("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(contents), "\n")
	lines = lines[:len(lines)-1]

	return lines
}

type point struct {
	val  byte
	x, y int
	last *point
}

func (p *point) String() string {
	var buf bytes.Buffer

	buf.WriteString("{ val: ")
	buf.WriteByte(p.val)
	buf.WriteString(
		fmt.Sprintf(", x: %d, y: %d }", p.x, p.y),
	)

	return buf.String()
}

func (p *point) equals(n *point) bool {
	return p.x == n.x && p.y == n.y
}

func (p *point) nextLocation(lines []string) *point {
	var next *point

	switch p.val {
	case '-':
		next = p.getLeft(lines)

		if next.equals(p.last) {
			next = p.getRight(lines)
		}
	case '|':
		next = p.getTop(lines)

		if next.equals(p.last) {
			next = p.getBottom(lines)
		}
	case 'L':
		next = p.getTop(lines)

		if next.equals(p.last) {
			next = p.getRight(lines)
		}
	case 'J':
		next = p.getTop(lines)

		if next.equals(p.last) {
			next = p.getLeft(lines)
		}
	case '7':
		next = p.getLeft(lines)

		if next.equals(p.last) {
			next = p.getBottom(lines)
		}
	case 'F':
		next = p.getRight(lines)

		if next.equals(p.last) {
			next = p.getBottom(lines)
		}
	default:
		log.Fatalf("invalid value %c used in nextLocation", p.val)
	}

	return next
}

func (p *point) getNeighbours(lines []string) []*point {
	neighbours := []*point{}

	left := p.getLeft(lines)

	if left != nil {
		neighbours = append(neighbours, left)
	}

	right := p.getRight(lines)

	if right != nil {
		neighbours = append(neighbours, right)
	}

	top := p.getTop(lines)

	if top != nil {
		neighbours = append(neighbours, top)
	}

	bottom := p.getBottom(lines)

	if bottom != nil {
		neighbours = append(neighbours, bottom)
	}

	return neighbours
}

func (p *point) getLeft(lines []string) *point {
	if p.x > 0 {
		left := lines[p.y][p.x-1]

		if strings.IndexByte("-LFS", left) > -1 {
			return &point{
				val:  left,
				x:    p.x - 1,
				y:    p.y,
				last: p,
			}
		}
	}

	return nil
}

func (p *point) getRight(lines []string) *point {
	if p.x < len(lines[0])-1 {
		right := lines[p.y][p.x+1]

		if strings.IndexByte("-J7S", right) > -1 {
			return &point{
				val:  right,
				x:    p.x + 1,
				y:    p.y,
				last: p,
			}
		}
	}

	return nil
}

func (p *point) getTop(lines []string) *point {
	if p.y > 0 {
		top := lines[p.y-1][p.x]

		if strings.IndexByte("|7FS", top) > -1 {
			return &point{
				val:  top,
				x:    p.x,
				y:    p.y - 1,
				last: p,
			}
		}
	}

	return nil
}

func (p *point) getBottom(lines []string) *point {
	if p.y < len(lines)-1 {
		bottom := lines[p.y+1][p.x]

		if strings.IndexByte("|LJS", bottom) > -1 {
			return &point{
				val:  bottom,
				x:    p.x,
				y:    p.y + 1,
				last: p,
			}
		}
	}

	return nil
}
