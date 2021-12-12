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

type Path struct {
	path           []string
	revisitedCave  string
	allowRevisting bool
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
	fmt.Println(exits.countPathsToTheEnd(Path{path: []string{"start"}, revisitedCave: "", allowRevisting: *part == "b"}))
}

func (p *Path) checkCaveIsVisitable(cave string) (mayVisit bool, isRevisit bool) {
	if cave[0] >= 'A' && cave[0] <= 'Z' {
		return true, false
	}
	if p.revisitedCave == cave || cave == "start" {
		return false, false
	}
	for _, c := range p.path {
		if c == cave {
			if p.revisitedCave == "" && p.allowRevisting {
				return true, true
			} else {
				return false, false
			}
		}
	}

	return true, false
}

func extend(p []string, cave string) []string {
	pathCopy := make([]string, len(p), len(p)+1)
	copy(pathCopy, p)
	pathCopy = append(pathCopy, cave)
	return pathCopy
}

func (e Exits) countPathsToTheEnd(path Path) int {
	pathsFromHere := 0

	curCave := path.path[len(path.path)-1]
	if curCave == "end" {
		pathsFromHere++

	} else {

		for _, c := range e[curCave] {
			mayVisit, isRevisit := path.checkCaveIsVisitable(c)
			if !mayVisit {
				continue
			}
			nextPath := Path{
				path:           extend(path.path, c),
				revisitedCave:  path.revisitedCave,
				allowRevisting: path.allowRevisting,
			}
			if isRevisit {
				nextPath.revisitedCave = c
			}
			pathsFromHere += e.countPathsToTheEnd(nextPath)
		}
	}

	return pathsFromHere
}
