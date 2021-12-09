package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day09.input", "Relative file path to use as input.")
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
	board := make(map[Coords]int)
	max := Coords{len(lines[0]) - 1, len(lines) - 1}
	lowpoints := make([]int, 0)
	sizes := make([]int, 0)

	for y, line := range lines {

		for x, p := range line {
			c := Coords{x, y}
			board[c] = int(p - '0')
		}
	}

	for y := 0; y <= max.y; y++ {
		for x := 0; x <= max.x; x++ {
			if checkPoint(board, x, y, max.y, max.x) {
				lowpoints = append(lowpoints, board[Coords{x, y}])
				basin := make(map[Coords]bool)
				basin[Coords{x, y}] = true
				getBasin(board, Coords{x, y}, basin)
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

func getNeighbours(board map[Coords]int, x, y int, ignoreDiagonals bool) []Coords {
	neighbours := make([]Coords, 0)
	for j := -1; j < 2; j++ {
		for i := -1; i < 2; i++ {
			if i == 0 && j == 0 {
				continue
			}

			if ignoreDiagonals && i != 0 && j != 0 {
				continue
			}

			p := Coords{x + i, y + j}
			if _, ok := board[p]; ok {
				neighbours = append(neighbours, p)
			}
		}
	}

	return neighbours
}

func checkPoint(board map[Coords]int, x, y, height, width int) bool {
	value := board[Coords{x, y}]

	for _, p := range getNeighbours(board, x, y, false) {
		if board[p] <= value {
			return false
		}
	}

	return true
}

func getBasin(board map[Coords]int, lp Coords, basin map[Coords]bool) {
	value := board[lp]
	for _, p := range getNeighbours(board, lp.x, lp.y, true) {
		if board[p] > value && board[p] < 9 {
			if _, ok := basin[p]; !ok {
				basin[p] = true
				getBasin(board, p, basin)
			}
		}
	}
}
