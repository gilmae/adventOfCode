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
			} else {
				xDelta := 1
				yDelta := 1
				if startX > endX {
					xDelta = -1
				}
				if startY > endY {
					yDelta = -1
				}
				for d := 0; d <= int(math.Abs(float64(endX)-float64(startX))); d++ {
					c := Coords{startX + xDelta*d, startY + yDelta*d}
					points[c]++
				}
				continue
			}
		}
		if startX > endX {
			startX, endX = endX, startX
		}
		if startY > endY {
			startY, endY = endY, startY
		}
		if startX != endX {

			for x := startX; x <= endX; x++ {
				c := Coords{x, startY}
				points[c]++
			}
		} else {

			for y := startY; y <= endY; y++ {
				c := Coords{startX, y}
				points[c]++
			}
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
