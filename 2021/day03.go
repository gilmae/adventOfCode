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
	epsilon := 0
	gamma := 0
	for i := 0; i < len(lines[0]); i++ {
		bits := map[byte]int{'0': 0, '1': 0}
		for _, l := range lines {
			bits[l[i]]++
		}
		pos := len(lines[i]) - 1
		if bits['0'] > bits['1'] {
			gamma += 1 << (pos - i)
		} else {
			epsilon += 1 << (pos - i)

		}
	}

	fmt.Println(epsilon * gamma)
}
