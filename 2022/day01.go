package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
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

	elves := make(map[int]int)
	elves[0] = 0
	elf := 0
	for _, l := range lines {
		if l == "" {
			elf += 1
			elves[elf] = 0
		} else {
			if c, err := strconv.Atoi(l); err == nil {
				elves[elf] += c
			}
		}
	}

	_, max := getMostCalories(elves)

	fmt.Println(max)
	allCalories := values(elves)
	sort.Sort(sort.Reverse(sort.IntSlice(allCalories)))
	fmt.Println(allCalories[0] + allCalories[1] + allCalories[2])
}

func getMostCalories(elves map[int]int) (int, int) {
	max := -1
	maxElf := -1
	for elf, c := range elves {
		if c > max {
			maxElf = elf
			max = c
		}
	}
	return maxElf, max
}

func values(elves map[int]int) []int {
	v := make([]int, 0)
	for _, c := range elves {
		v = append(v, c)
	}

	return v
}
