package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gilmae/adventOfCode/2021/boards"
)

type cucumber byte
type seafloor map[boards.Coords]byte

var inputFile = flag.String("inputFile", "inputs/day25.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	seafloor := make(seafloor)

	width := len(lines[0])
	height := len(lines)

	fmt.Println(width, height)

	for row, line := range lines {
		for col, c := range line {
			pos := boards.Coords{col, row}
			switch c {
			case '>':
				seafloor[pos] = '>'
			case 'v':
				seafloor[pos] = 'v'

			}
		}
	}
	//seafloor.draw(width, height)
	//fmt.Println()
	var someoneMoved bool
	for i := 1; ; i++ {

		seafloor, someoneMoved = seafloor.swim(width, height)
		if i == 1 {
			seafloor.draw(width, height)
			fmt.Println(i)

		}
		if !someoneMoved {
			seafloor.draw(width, height)
			fmt.Println(i)
			break
		}
	}

}

func (s seafloor) swim(width, height int) (seafloor, bool) {
	next := make(seafloor)
	var someoneMoved bool

	// Move east herd first
	for k, v := range s {
		if v == '>' {
			dest := boards.Coords{k.X + 1, k.Y}
			wrap(&dest, width, height)
			if _, ok := s[dest]; ok {
				next[k] = v
			} else {
				next[dest] = v
				someoneMoved = true
			}
		} else {
			next[k] = v
		}
	}

	s = next
	next = make(seafloor)
	// Then move south herd
	for k, v := range s {
		if v == 'v' {
			dest := boards.Coords{k.X, k.Y + 1}
			wrap(&dest, width, height)
			if _, ok := s[dest]; ok {
				next[k] = v
			} else {
				next[dest] = v
				someoneMoved = true
			}
		} else {
			next[k] = v
		}
	}
	return next, someoneMoved
}

func (s seafloor) draw(width, height int) {

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if v, ok := s[boards.Coords{col, row}]; ok {
				if v == '>' {
					fmt.Print(">")
				} else if v == 'v' {
					fmt.Print("v")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func wrap(c *boards.Coords, width, height int) {
	if c.X >= width {
		c.X = 0
	}
	if c.Y >= height {
		c.Y = 0
	}
}
