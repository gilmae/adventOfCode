package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day19.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

type Coords3d struct {
	X, Y, Z int
}

type Scanner struct {
	seen []Coords3d
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	scanners := make([]*Scanner, 0)
	var scanner *Scanner
	for _, line := range lines {
		fmt.Println(line)
		if len(line) == 0 {
			continue
		}
		if line[0:3] == "---" {
			scanner = &Scanner{}
			scanners = append(scanners, scanner)
			continue
		}
		c := Coords3d{}
		parts := strings.Split(line, ",")
		c.X, _ = strconv.Atoi(parts[0])
		c.Y, _ = strconv.Atoi(parts[1])
		c.Z, _ = strconv.Atoi(parts[2])

		scanner.seen = append(scanner.seen, c)
	}

}
