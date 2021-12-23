package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day02.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

type Coords struct {
	x, y int
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	keypadA := map[Coords]byte{
		Coords{0, 0}: '1',
		Coords{1, 0}: '2',
		Coords{2, 0}: '3',
		Coords{0, 1}: '4',
		Coords{1, 1}: '5',
		Coords{2, 1}: '6',
		Coords{0, 2}: '7',
		Coords{1, 2}: '8',
		Coords{2, 2}: '9'}

	keypadB := map[Coords]byte{
		Coords{2, 0}: '1',
		Coords{1, 1}: '2',
		Coords{2, 1}: '3',
		Coords{3, 1}: '4',
		Coords{0, 2}: '5',
		Coords{1, 2}: '6',
		Coords{2, 2}: '7',
		Coords{3, 2}: '8',
		Coords{4, 2}: '9',
		Coords{1, 3}: 'A',
		Coords{2, 3}: 'B',
		Coords{3, 3}: 'C',
		Coords{2, 4}: 'D',
	}
	code := make([]byte, len(lines))
	var position Coords
	var keys map[Coords]byte

	if *part == "a" {
		position = Coords{1, 1}
		keys = keypadA
	} else {
		position = Coords{0, 2}
		keys = keypadB

	}
	var nextPosition Coords

	for digit, line := range lines {

		for i, _ := range line {
			switch line[i] {
			case 'U':
				nextPosition = Coords{position.x, position.y - 1}
			case 'D':
				nextPosition = Coords{position.x, position.y + 1}
			case 'L':
				nextPosition = Coords{position.x - 1, position.y}
			case 'R':
				nextPosition = Coords{position.x + 1, position.y}
			}

			if _, ok := keys[nextPosition]; ok {
				position = nextPosition
			}

			//fmt.Println(string(line[i]), nextPosition, " => ", position)
		}
		//fmt.Println()
		code[digit] = keys[position]
	}

	fmt.Println(string(code))

}
