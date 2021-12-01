package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day01.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	switch *part {
	case "a":
		fmt.Println(partA(lines))
	case "b":
		fmt.Println(partB(lines))
	default:
		fmt.Printf("Unknown part: %s\n", *part)
	}

}

func partA(lines []string) int {
	previous := 0

	count := -1
	for i := range lines {
		depth, _ := strconv.Atoi(lines[i])

		if depth > previous {
			count++
		}
		previous = depth
	}
	return count
}

func partB(lines []string) int {
	previous := 0

	count := -1
	for i := 0; i < len(lines)-2; i++ {
		one, _ := strconv.Atoi(lines[i])
		two, _ := strconv.Atoi(lines[i+1])
		three, _ := strconv.Atoi(lines[i+2])

		depth := one + two + three
		if depth > previous {
			count++
		}
		previous = depth
	}
	return count
}
