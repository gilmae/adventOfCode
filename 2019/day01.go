package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day01.input", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Show debug messages")
var partB = flag.Bool("partB", false, "Do Part B")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")

	sum := 0

	for _, s := range split {
		weight, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Failed to parse %s", s)
		}
		var fuel int
		if *partB {
			fuel = calculateFuel(weight)
		} else {
			fuel = weight/3 - 2
		}
		if *debug {
			fmt.Println(fuel)
		}
		sum += fuel
	}

	fmt.Printf("Fuel required: %d\n", sum)
}

func calculateFuel(weight int) int {
	fuel := weight/3 - 2
	if fuel > 0 {
		return fuel + calculateFuel(fuel)
	}
	return 0
}
