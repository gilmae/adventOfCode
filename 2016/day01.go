package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day01.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

const (
	north = 0
	east  = 1
	south = 2
	west  = 3
)

type Coords struct {
	x, y int
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	spot := Coords{0, 0}
	facing := 0
	seen := make(map[Coords]int)
	var hasRepeated bool
	var repeat Coords

	for _, step := range strings.Split(lines[0], ", ") {
		switch step[0] {
		case 'L':
			facing = (facing - 1) % 4

		case 'R':
			facing = (facing + 1) % 4

		}

		if facing < north {
			facing = facing + 4
		}
		steps, _ := strconv.Atoi(step[1:])
		for s := 0; s < steps; s++ {
			switch facing {
			case 0:
				spot = Coords{spot.x, spot.y + 1}

			case 1:
				spot = Coords{spot.x + 1, spot.y}

			case 3:
				spot = Coords{spot.x - 1, spot.y}

			case 2:
				spot = Coords{spot.x, spot.y - 1}

			}
			seen[spot] += 1
			if v, ok := seen[spot]; ok && v > 1 && !hasRepeated {
				hasRepeated = true
				repeat = spot

			}
		}

	}

	fmt.Println((math.Abs(float64(spot.x)) + math.Abs(float64(spot.y))))
	fmt.Println("First repeat: ", repeat, (math.Abs(float64(repeat.x)) + math.Abs(float64(repeat.y))))

}
