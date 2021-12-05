package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day24.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	weights := make([]int, len(lines))
	totalWeight := 0

	for i, l := range lines {
		weight, _ := strconv.Atoi(l)
		weights[i] = weight
		totalWeight += weight
	}
	sort.Sort(sort.Reverse(sort.IntSlice(weights)))

	goalWeight := totalWeight / 3
	if *part == "b" {
		goalWeight = totalWeight / 4
	}

	fmt.Println(goalWeight)
	for n := 1; ; n++ {
		if found, quantumEntaglement := fillBag(weights, goalWeight, n); found {
			fmt.Println(quantumEntaglement)
			return
		}
	}

}

func fillBag(weights []int, weight int, numPackages int) (bool, uint64) {
	found := false
	smallestEntaglement := uint64(math.MaxUint64)

	for i, w := range weights {
		remainingWeight := weight - w
		if remainingWeight < 0 {
			continue
		} else if remainingWeight == 0 {
			return true, uint64(w)
		} else if numPackages > 1 {
			if subFound, subEntanglement := fillBag(weights[i+1:], remainingWeight, numPackages-1); subFound {
				found = true
				entanglement := subEntanglement * uint64(w)
				if entanglement < smallestEntaglement {
					smallestEntaglement = entanglement
				}
			}
		}
	}

	return found, smallestEntaglement
}
