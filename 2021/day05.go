package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day05.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

type Coords struct {
	x, y int
}

func traverse(start, end Coords) []Coords {
	var deltaX int
	length := 1
	if start.x != end.x {
		length += int(math.Abs(float64(end.x) - float64(start.x)))
	} else {
		length += int(math.Abs(float64(end.y) - float64(start.y)))
	}
	out := make([]Coords, length)

	if start.x > end.x {
		deltaX = -1
	} else if start.x == end.x {
		deltaX = 0
	} else {
		deltaX = 1
	}

	var deltaY int
	if start.y > end.y {
		deltaY = -1
	} else if start.y == end.y {
		deltaY = 0
	} else {
		deltaY = 1
	}

	for d := 0; d < length; d++ {
		c := Coords{start.x + deltaX*d, start.y + deltaY*d}
		out[d] = c
	}
	return out
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	rx := regexp.MustCompile("(\\d+),(\\d+) -> (\\d+),(\\d+)")

	points := make(map[Coords]int)

	for _, line := range lines {

		sm := rx.FindStringSubmatch(line)
		startX, _ := strconv.Atoi(sm[1])
		startY, _ := strconv.Atoi(sm[2])

		endX, _ := strconv.Atoi(sm[3])
		endY, _ := strconv.Atoi(sm[4])

		if startX != endX && startY != endY {
			if *part == "a" {
				// No diagonals in part a
				continue
			}
		}
		for _, c := range traverse(Coords{startX, startY}, Coords{endX, endY}) {
			points[c] += 1
		}
	}

	numDangerousPoints := 0
	for _, v := range points {
		if v > 1 {
			numDangerousPoints++
		}
	}

	// Debug Simple Input
	// for y := 0; y < 10; y++ {
	// 	for x := 0; x < 10; x++ {
	// 		c := Coords{x, y}
	// 		if v, ok := points[c]; !ok || v == 0 {
	// 			fmt.Print(".")
	// 		} else {
	// 			fmt.Print(v)
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	fmt.Println(numDangerousPoints)

}
