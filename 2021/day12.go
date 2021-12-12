package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day12.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

type Exits map[string][]string
type Path []string

func (p *Path) mayVisit(cave string) bool {
	if cave[0] >= 'A' && cave[0] <= 'Z' {
		return true
	}
	for _, c := range *p {
		if c == cave {
			return false
		}

	}
	return true
}

func (p Path) extend(next string) Path {
	pathCopy := make(Path, len(p), len(p)+1)
	copy(pathCopy, p)
	pathCopy = append(pathCopy, next)
	return pathCopy
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	exits := make(Exits)

	for _, l := range lines {
		parts := strings.Split(l, "-")

		exits[parts[0]] = append(exits[parts[0]], parts[1])
		exits[parts[1]] = append(exits[parts[1]], parts[0])
	}

	possiblePaths := []Path{[]string{"start"}}

	fmt.Println(exits.searchForPaths(possiblePaths))
}

func (exits Exits) searchForPaths(queue []Path) int {
	paths := 0
	for len(queue) != 0 {
		var nextQueue []Path
		for _, item := range queue {

			last := item[len(item)-1]
			if last == "end" {
				// This is a unique path that has reached the end.
				paths++
				continue
			}
			for _, n := range exits[last] {
				mayVisit := item.mayVisit(n)
				if !mayVisit {
					continue
				}
				nextItem := item.extend(n)

				nextQueue = append(nextQueue, nextItem)
			}
		}
		queue = nextQueue
	}
	return paths
}
