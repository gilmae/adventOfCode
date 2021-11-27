package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day16.input", "Relative file path to use as input.")

type preference struct {
	thing  string
	amount int
}

func Intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	rx := regexp.MustCompile("Sue (\\d+): (\\w+): (\\d+), (\\w+): (\\d+), (\\w+): (\\d+)")

	likes := make(map[preference][]string)
	for l := range lines {
		submatches := rx.FindStringSubmatch(lines[l])
		aunt := submatches[1]
		for i := 0; i < 3; i++ {
			thing := submatches[2+i]
			amount, _ := strconv.Atoi(submatches[3+i])
			pref := preference{thing, amount}
			_, ok := likes[pref]
			if !ok {
				likes[pref] = make([]string, 0)
			}
			likes[pref] = append(likes[pref], aunt)
		}
	}

	detectedPreferences := []preference{
		preference{"children", 3},
		preference{"cats", 7},
		preference{"samoyeds", 2},
		preference{"pomeranians", 3},
		preference{"akitas", 0},
		preference{"vizslas", 0},
		preference{"goldfish", 5},
		preference{"trees", 3},
		preference{"cars", 2},
		preference{"perfumes", 1},
	}

	sues := make(map[string]int)
	for i := range detectedPreferences {

		pref := detectedPreferences[i]
		if pref.amount == 0 {
			continue
		}
		found, ok := likes[pref]
		if ok {
			for f := range found {
				sues[found[f]]++
			}
		}

	}

	// Find Sue with most preferences found. In theory there should be only one sue with the max preferences found
	max := 1
	suesWithMax := make([]string, 0)
	for s := range sues {
		if sues[s] > max {
			max = sues[s]
			suesWithMax = append(suesWithMax, s)
		}
	}
	fmt.Println(suesWithMax)
}
