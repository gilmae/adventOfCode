package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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

	if *part == "a" {
		partA(lines)
	} else {
		partB(lines)
	}

}

func getCommon(i string, j string) string {
	hash := make(map[rune]bool)
	common := make([]byte, 0)
	for _, r := range i {
		hash[r] = true
	}
	for _, r := range j {
		if _, ok := hash[r]; ok {
			common = append(common, byte(r))
		}
	}

	return string(common[:])
}

func getPriority(common byte) int {
	priority := 0

	if common <= 'Z' {
		priority += (int(common) - 38)
	} else {
		priority += (int(common) - 96)
	}
	return priority
}

func partA(lines []string) {
	priorities := 0
	for _, line := range lines {
		bag1 := len(line) / 2
		common := getCommon(string(line[0:bag1]), string(line[bag1:]))
		priorities += getPriority(common[0])
	}
	fmt.Println((priorities))
}

func partB(lines []string) {
	priority := 0
	for i := 0; i < len(lines); i += 3 {
		common := getCommon(lines[i], lines[i+1])
		common = getCommon(common, lines[i+2])
		// We're assured only one common item, but the simple input returns 'ZZ' for the second set of three, so clayton's dedupe
		priority += getPriority(common[0])
	}

	fmt.Println(priority)
}
