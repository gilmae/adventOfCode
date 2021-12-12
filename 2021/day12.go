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
type PathWithOneRevisit struct {
	Path    Path
	revisit string
}

func (p *PathWithOneRevisit) mayVisit(cave string) (allowed bool, theRevisit bool) {
	if cave[0] >= 'A' && cave[0] <= 'Z' {
		return true, false
	}

	if cave == p.revisit || cave == "start" {
		return false, false
	}
	for _, c := range *&p.Path {
		if c == cave {
			if p.revisit == "" {
				// We've been here before, but never revisited a cave before so this is our one revisit
				return true, true
			} else {
				return false, false
			}

		}

	}
	return true, false
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

	possiblePaths := []PathWithOneRevisit{
		{Path: Path{"start"}, revisit: ""},
	}

	fmt.Println(exits.searchForPaths(possiblePaths))
}

func (exits Exits) searchForPaths(queue []PathWithOneRevisit) int {
	paths := 0
	for len(queue) != 0 {
		var nextQueue []PathWithOneRevisit
		for _, item := range queue {
			last := item.Path[len(item.Path)-1]
			if last == "end" {
				paths++
				continue
			}
			for _, n := range exits[last] {
				mayVisit, setRevisit := item.mayVisit(n)
				if !mayVisit {
					continue
				}
				nextItem := PathWithOneRevisit{
					Path:    item.Path.extend(n),
					revisit: item.revisit,
				}

				if setRevisit {
					nextItem.revisit = n
				}
				nextQueue = append(nextQueue, nextItem)
			}
		}
		queue = nextQueue
	}
	return paths
}
