package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day02.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	depth := 0
	position := 0
	for _, l := range lines {
		parts := strings.Split(l, " ")
		distance, _ := strconv.Atoi(parts[1])

		switch parts[0] {
		case "up":
			depth -= distance
		case "down":
			depth += distance
		case "forward":
			position += distance
		default:
			fmt.Println("Unkonwn direction: %s", parts[0])
		}
	}

	fmt.Println(depth * position)

	if *part == "b" {
		depth = 0
		position = 0
		aim := 0
		for _, l := range lines {
			parts := strings.Split(l, " ")
			distance, _ := strconv.Atoi(parts[1])

			switch parts[0] {
			case "up":
				//depth -= distance
				aim -= distance
			case "down":
				//depth += distance
				aim += distance
			case "forward":
				position += distance
				depth += distance * aim
			default:
				fmt.Println("Unkonwn direction: %s", parts[0])
			}

		}
		fmt.Println(depth * position)
	}
}
