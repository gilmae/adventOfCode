package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day04.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	contained := 0
	overlaps := 0
	// Set Difference, I think
	for _, line := range lines {
		jobs := strings.Split(line, ",")
		one := parseJobs(jobs[0])
		two := parseJobs(jobs[1])

		three := difference(one, two)
		if len(three) == 0 {
			contained += 1
		} else {

			three = difference(two, one)
			if len(three) == 0 {
				contained += 1
			}
		}
		three = intersect(one, two)
		if len(three) > 0 {

			overlaps += 1
		}
	}

	fmt.Println(contained)
	fmt.Println(overlaps)
}

func parseJobs(job string) []int {
	parts := strings.Split(job, "-")
	if len(parts) < 2 {
		fmt.Println("Invalid job: %s", job)
		return nil
	}
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Errorf("Invalid start in job: %s, %s", start, job)

	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Errorf("Invalid end: %s, %s", end, job)

	}
	bits := make([]int, 0)
	for i := start; i <= end; i++ {
		bits = append(bits, i)
	}

	return bits
}

func difference(one []int, two []int) []int {
	bits := make(map[int]bool)

	for _, b := range one {
		bits[b] = true
	}

	for _, b := range two {
		if _, ok := bits[b]; ok {
			delete(bits, b)
		}
	}
	three := make([]int, 0)
	for k, _ := range bits {
		three = append(three, k)
	}

	return three
}

func intersect(one []int, two []int) []int {
	three := make([]int, 0)
	oneMap := make(map[int]bool)
	for _, b := range one {
		oneMap[b] = true
	}
	twoMap := make(map[int]bool)
	for _, b := range two {
		twoMap[b] = true
	}
	for k, _ := range oneMap {
		if _, ok := twoMap[k]; ok {
			three = append(three, k)
		}
	}

	return three
}
