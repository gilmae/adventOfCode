package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Pair struct {
	x, y string
}

var inputFile = flag.String("inputFile", "inputs/day13.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	pairings := make(map[Pair]int)
	people := make(map[string]bool)
	regex := regexp.MustCompile("([a-zA-Z]+) would (gain|lose) ([0-9]+) happiness units by sitting next to ([a-zA-Z]+).")

	for line := range lines {
		subMatches := regex.FindStringSubmatch(lines[line])

		pair := Pair{subMatches[1], subMatches[4]}
		happiness, _ := strconv.Atoi(subMatches[3])
		if subMatches[2] == "lose" {
			happiness *= -1
		}

		pairings[pair] = happiness
		people[pair.x] = true
		people[pair.y] = true
	}
	people["me"] = true

	individuals := make([]string, len(people))
	i := 0
	for k := range people {
		individuals[i] = k
		i++
	}

	fmt.Println(arrangeTable(nil, nil, pairings, individuals))
}

func arrangeTable(currentPerson, lastPerson *string, happinesses map[Pair]int, people []string) int {
	longest := 0

	for i := range people {
		person := people[i]

		myLast := lastPerson
		if myLast == nil {
			myLast = &person
		}

		remainingPeople := make([]string, len(people)-1)
		copy(remainingPeople[:i], people[:i])
		copy(remainingPeople[i:], people[i+1:])

		total := 0

		if len(remainingPeople) > 0 {
			total += arrangeTable(&person, myLast, happinesses, remainingPeople)
		} else {
			total += happinesses[Pair{person, *myLast}]
			total += happinesses[Pair{*myLast, person}]
		}

		if currentPerson != nil {
			total += happinesses[Pair{person, *currentPerson}]
			total += happinesses[Pair{*currentPerson, person}]
		}

		if total > longest {
			longest = total
		}

	}
	return longest

}
