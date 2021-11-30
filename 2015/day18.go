package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

// Fucking game of life.

type Coords struct {
	x, y int
}

var inputFile = flag.String("inputFile", "inputs/day18.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	fmt.Print(len(lines))
	lights := make(map[Coords]bool)
	for y := range lines {
		for x := range lines[y] {
			c := Coords{x, y}
			if lines[y][x] == '#' {
				lights[c] = true
			}

		}
	}
	fmt.Println(len(lights))

	for s := 0; s < 100; s++ {
		lights = tick(lights, len(lines))
	}

	fmt.Println(len(lights))
}

func countLitNeightbours(c Coords, state map[Coords]bool, length int) int {
	count := 0
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if c.x+x < 0 || c.y+y < 0 {
				continue
			}
			if c.x+x >= length || c.y+y >= length {
				continue
			}
			light, ok := state[Coords{c.x + x, c.y + y}]
			if !ok {
				continue
			}
			if light {
				count += 1
			}
		}
	}

	return count
}

func tick(state map[Coords]bool, length int) map[Coords]bool {
	nextState := make(map[Coords]bool)
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			c := Coords{x, y}
			neighbours := countLitNeightbours(c, state, length)
			_, ok := state[c]
			if !ok {
				if neighbours == 3 {
					nextState[c] = true
				}
			} else {
				if neighbours == 2 || neighbours == 3 {
					nextState[c] = true
				}
			}
		}
	}

	// Part B corner lights stuck on
	nextState[Coords{0, 0}] = true
	nextState[Coords{0, length - 1}] = true
	nextState[Coords{length - 1, 0}] = true
	nextState[Coords{length - 1, length - 1}] = true
	return nextState
}
