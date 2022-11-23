package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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

	columns := make([]map[rune]int, len(lines[0]))
	for col := 0; col < len(columns); col++ {
		columns[col] = make(map[rune]int)

		for _, line := range lines {
			columns[col][rune(line[col])] += 1
		}
	}
	f := getKeyWithMost
	if *part == "b" {
		f = getKeyWithLeast
	}
	for _, chars := range columns {
		fmt.Printf("%s", string(f(chars)))
	}
	fmt.Println()
}

func getKeyWithMost(chars map[rune]int) rune {
	max := -1
	maxKey := rune('a')
	for k, count := range chars {
		if count > max {
			maxKey = k
			max = count
		}
	}
	return maxKey
}

func getKeyWithLeast(chars map[rune]int) rune {
	min := 1000000000
	minKey := rune('a')
	for k, count := range chars {
		if count < min {
			minKey = k
			min = count
		}
	}
	return minKey
}
