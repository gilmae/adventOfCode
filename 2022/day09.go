package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day09.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	chain := make([]link, 2)
	fmt.Println(followMoves(chain, lines))

	chain = make([]link, 10)
	fmt.Println(followMoves(chain, lines))

}

func followMoves(c []link, moves []string) int {
	visited := make(map[string]bool)
	visited[c[len(c)-1].String()] = true
	for _, line := range moves {
		direction := byte(line[0])
		steps, _ := strconv.Atoi(line[2:])

		for i := 0; i < steps; i++ {
			c[0].move(direction)
			for j, link := range c[1:] {
				link.follow(c[j])
				c[j+1] = link
			}

			visited[c[len(c)-1].String()] = true
		}
	}
	return len(visited)
}

type link struct {
	x, y int
}

func NewLink(x, y int) *link {
	return &link{x: x, y: y}
}

func (l *link) String() string {
	return fmt.Sprintf("%d,%d", l.x, l.y)
}

func (l *link) isTouching(other link) bool {
	return math.Abs(float64(l.x-other.x)) < 2 && math.Abs(float64(l.y-other.y)) < 2
}

func (l *link) move(direction byte) {
	switch direction {
	case 'U':
		l.y += 1
	case 'D':
		l.y -= 1
	case 'L':
		l.x -= 1
	case 'R':
		l.x += 1
	}
}

func (l *link) follow(other link) {
	if l.isTouching(other) {
		return
	}
	dx := other.x - l.x
	if dx < 0 {
		l.x -= 1
	} else if dx > 0 {
		l.x += 1
	}
	dy := other.y - l.y
	if dy < 0 {
		l.y -= 1
	} else if dy > 0 {
		l.y += 1
	}
}
