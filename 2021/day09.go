package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/gilmae/adventOfCode/2021/boards"
)

var inputFile = flag.String("inputFile", "inputs/day09.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	board := boards.NewBoard()

	board.Import(lines, func(c boards.Coords, ch interface{}) interface{} {
		return int(ch.(rune) - '0')
	})

	maxX, maxY := board.Width(), board.Height()
	lowpoints := make([]int, 0)
	sizes := make([]int, 0)

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if checkPoint(board, boards.Coords{x, y}) {
				lowpoints = append(lowpoints, board.Points[boards.Coords{x, y}].(int))
				basin := make(map[boards.Coords]bool)
				basin[boards.Coords{x, y}] = true
				getBasin(board, boards.Coords{x, y}, basin)
				sizes = append(sizes, len(basin))
			}
		}
	}

	sum := 0
	for _, lp := range lowpoints {
		sum += lp + 1
	}
	fmt.Println(sum)
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	fmt.Println(sizes[0] * sizes[1] * sizes[2])

}

func getNeighbours(board *boards.Board, c boards.Coords, ignoreDiagonals bool) []boards.Coords {
	neighbours := make([]boards.Coords, 0)
	for j := -1; j < 2; j++ {
		for i := -1; i < 2; i++ {
			if i == 0 && j == 0 {
				continue
			}

			if ignoreDiagonals && i != 0 && j != 0 {
				continue
			}

			p := boards.Coords{c.X + i, c.Y + j}
			if _, ok := board.Points[p]; ok {
				neighbours = append(neighbours, p)
			}
		}
	}

	return neighbours
}

func checkPoint(board *boards.Board, c boards.Coords) bool {
	value := board.Points[c].(int)

	for _, p := range getNeighbours(board, c, false) {
		v := board.Points[p].(int)
		if v <= value {
			return false
		}
	}

	return true
}

func getBasin(board *boards.Board, c boards.Coords, basin map[boards.Coords]bool) {
	value := board.Points[c].(int)
	for _, p := range getNeighbours(board, c, true) {
		if board.Points[p].(int) > value && board.Points[p].(int) < 9 {
			if _, ok := basin[p]; !ok {
				basin[p] = true
				getBasin(board, p, basin)
			}
		}
	}
}
