package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day03.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	validCount := 0
	if *part == "a" {
		for _, line := range lines {
			sideA, _ := strconv.Atoi(strings.TrimSpace(line[0:5]))
			sideB, _ := strconv.Atoi(strings.TrimSpace(line[5:10]))
			sideC, _ := strconv.Atoi(strings.TrimSpace(line[10:15]))

			if sideA+sideB > sideC && sideA+sideC > sideB && sideB+sideC > sideA {
				validCount++
			}
		}
		fmt.Println(validCount)
	} else {
		line := 0
		for line+2 < len(lines) {
			for triangle := 0; triangle < 3; triangle++ {
				sideA, _ := strconv.Atoi(strings.TrimSpace(lines[line][triangle*5 : triangle*5+5]))
				sideB, _ := strconv.Atoi(strings.TrimSpace(lines[line+1][triangle*5 : triangle*5+5]))
				sideC, _ := strconv.Atoi(strings.TrimSpace(lines[line+2][triangle*5 : triangle*5+5]))
				if sideA+sideB > sideC && sideA+sideC > sideB && sideB+sideC > sideA {
					validCount++
				}
			}

			line += 3
		}
		fmt.Println(validCount)
	}

}

//925
