package main

import (
	"flag"
	"fmt"
	"strconv"
)

var part = flag.String("part", "a", "Which part")

func main() {
	flag.Parse()
	passwd := 382344
	validCount := 0
	for passwd <= 843167 {
		passwd++
		if isValid(passwd) {
			validCount++
		}
	}

	fmt.Println(validCount)

}

func isValid(passwd int) bool {
	digits := strconv.Itoa(passwd + 1)
	var previous rune
	var double bool
	var exactDouble bool
	run := 1

	for _, r := range digits {
		if previous == r {
			double = true
			run++
		} else {
			if run == 2 {
				exactDouble = true
			}
			run = 1
		}
		if previous > r {
			return false
		}

		previous = r
	}

	if run == 2 {
		exactDouble = true
	}

	if *part == "b" {
		return exactDouble
	} else {
		return double
	}

}
