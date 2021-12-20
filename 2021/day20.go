package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gilmae/adventOfCode/2021/boards"
)

var inputFile = flag.String("inputFile", "inputs/day20.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")
var debug = flag.Bool("debug", false, "Show debug messages")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	enhancement := make([]bool, 512)
	for i, c := range lines[0] {
		switch c {
		case '.':
			// Pass
		case '#':
			enhancement[i] = true
		default:
			fmt.Println("invalid character.")
			return
		}
	}

	board := boards.NewBoard()
	board.Import(lines[2:], func(c boards.Coords, v interface{}) interface{} {
		return v == '#'
	})

	for i := 1; i <= 50; i++ {
		board = enhance(board, enhancement, i)
		if i == 2 {
			fmt.Println(countOnBits(board))
		}
		if *debug {
			board.PrintBoard()
		}
	}
	//board = enhance(board, enhancement, 1)
	//board = enhance(board, enhancement, 2)
	fmt.Println(countOnBits(board))
}

func countOnBits(b *boards.Board) int {
	minX, minY := b.TopCorner()
	maxX, maxY := b.BottomCorner()

	count := 0
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if v, ok := b.Points[boards.Coords{x, y}]; ok && v != nil && v.(bool) {
				count++
			}
		}
	}

	return count
}

func enhance(b *boards.Board, enchancement []bool, i int) *boards.Board {
	newBoard := boards.NewBoard()

	minX, minY := b.TopCorner()
	maxX, maxY := b.BottomCorner()

	var void bool

	// Sample input is tricky and deliberately avoids toggling the void on and off,
	// but real input absolutely toggles it on and off
	// Which, btw the way, was cruel
	if enchancement[0] && !enchancement[511] {
		void = (i % 2) == 0
	} else {
		void = false
	}

	for y := minY - 2; y <= maxY+2; y++ {
		for x := minX - 2; x <= maxX+2; x++ {

			k := boards.Coords{x, y}

			byyte := 0

			for _, n := range k.GetNeighbours(1, true) {
				byyte = byyte << 1
				if v, ok := b.Points[n]; ok && v.(bool) {

					byyte += 1
				} else if !ok && void {
					byyte += 1
				}

			}

			newBoard.Points[k] = enchancement[byyte]
		}
	}
	return newBoard
}
