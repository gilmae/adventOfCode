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

	for _, step := range strings.Split(lines[0], ", ") {
		//fmt.Printf("Facing %d at %d,%d, next command is %s\n", facing, n, e, step)
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
		switch facing {
		case 0:
			spot.y += steps
		case 1:
			spot.x += steps
		case 3:
			spot.x -= steps
		case 2:
			spot.y -= steps
		}
	}

	fmt.Println((math.Abs(float64(spot.x)) + math.Abs(float64(spot.y))))
}
