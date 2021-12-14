package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day14.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")
var rules map[pair]byte

type pair struct {
	left, right byte
}

type polymer struct {
	pairs     map[pair]int
	outerPair pair
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	memo := lines[0]

	rules = make(map[pair]byte)
	for i := 2; i < len(lines); i++ {
		s := lines[i]
		parts := strings.Split(s, " -> ")
		rules[pair{parts[0][0], parts[0][1]}] = parts[1][0]
	}

	poly := polymer{pairs: make(map[pair]int), outerPair: pair{memo[0], memo[len(memo)-1]}}

	for i := 0; i < len(memo)-1; i++ {
		p := pair{memo[i], memo[i+1]}
		poly.pairs[p] += 1
	}

	for i := 0; i < 10; i++ {
		poly = applyRules(poly)
	}
	fmt.Println(getFrequencies(poly))

	if *part == "b" {
		for i := 0; i < 30; i++ {
			poly = applyRules(poly)
		}
		fmt.Println(getFrequencies(poly))
	}

}

func applyRules(poly polymer) polymer {
	pairs := make(map[pair]int)
	for p, count := range poly.pairs {
		if i, ok := rules[p]; ok {
			pairs[pair{p.left, i}] += count
			pairs[pair{i, p.right}] += count
		}
	}
	return polymer{pairs: pairs, outerPair: poly.outerPair}
}

func getFrequencies(poly polymer) int {
	elements := make(map[biyte]int)
	for p, c := range poly.pairs {
		elements[p.left] += c
		elements[p.right] += c
	}
	elements[poly.outerPair.left]++
	elements[poly.outerPair.right]++

	largest := 0
	smallest := int(^uint(0)>>1) - 1

	for _, v := range elements {
		if v > largest {
			largest = v
		}
		if v < smallest {
			smallest = v
		}
	}
	// each byte would have bee counted twice, once as the left in a pair, and once as the right in a pair.
	// so divide in two
	return (largest - smallest) / 2
}
