package main

import (
	"flag"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day02.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

type amphipod byte
type Room struct {
	habitat     amphipod
	top, bottom amphipod
}

const (
	A     byte = 'A'
	B     byte = 'B'
	C     byte = 'C'
	D     byte = 'D'
	Empty byte = ' '
)

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
}
