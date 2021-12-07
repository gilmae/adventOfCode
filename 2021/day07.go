package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day07.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	subs := make(map[int]int)
	minPos := 100000000
	maxPos := -10000000
	for _, s := range strings.Split(lines[0], ",") {
		sub, _ := strconv.Atoi(s)
		subs[sub] += 1
		if sub < minPos {
			minPos = sub
		}

		if sub > maxPos {
			maxPos = sub
		}
	}

	fuelForDistanceChange := make(map[int]int)
	fuelForDistanceChange[0] = 0

	for i := 1; i <= maxPos-minPos; i++ {
		fuelForDistanceChange[i] = i + fuelForDistanceChange[i-1]
	}

	minFuel := 1000000000
	alignmentPos := -1
	for pos := minPos; pos <= maxPos; pos++ {

		fuel := 0
		for k, v := range subs {
			distanceToMove := int(math.Abs(float64(pos - k)))
			if *part == "a" {
				fuel += distanceToMove * v
			} else {
				fuel += fuelForDistanceChange[distanceToMove] * v
			}
		}
		if fuel < minFuel {
			alignmentPos = pos
			minFuel = fuel
		}
	}

	fmt.Printf("Aligining on position %d costs %d\n", alignmentPos, minFuel)

}
