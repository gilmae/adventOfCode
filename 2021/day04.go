package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day04.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

type Coords struct {
	x, y int
}

type Board struct {
	grid        map[Coords]bool
	gridNumbers map[int]Coords
}

func NewBoard(grid [][]int) Board {
	b := Board{grid: make(map[Coords]bool), gridNumbers: make(map[int]Coords)}
	for y := range grid {
		for x, n := range grid[y] {
			c := Coords{x, y}
			b.grid[c] = false
			b.gridNumbers[n] = c
		}
	}
	return b
}

func (b *Board) HasWinner() bool {
	// Check rows and columns
	for i := 0; i < 5; i++ {
		// Check row
		if b.check(0, i, 1, 0, 5) {
			return true
		}
		// Check column
		if b.check(i, 0, 0, 1, 5) {
			return true
		}

	}
	return false

}

func (b *Board) Score() int {
	total := 0
	for k, v := range b.gridNumbers {
		if !b.grid[v] {
			total += k
		}
	}

	return total
}

func (b *Board) CallNumber(number int) bool {
	c, ok := b.gridNumbers[number]
	if ok {
		b.grid[c] = true
		return b.HasWinner()
	}
	return false
}

func (b *Board) check(startX, startY, xDelta, yDelta, steps int) bool {
	for i := 0; i < steps; i++ {
		c := Coords{startX + xDelta*i, startY + yDelta*i}
		if !b.grid[c] {
			return false
		}
	}
	return true
}

func main() {
	rx := regexp.MustCompile("(\\d+)\\s+(\\d+)\\s+(\\d+)\\s+(\\d+)\\s+(\\d+)")
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	calls := strings.Split(lines[0], ",")

	numBoards := (len(lines) - 1) / 6
	boards := make([]Board, numBoards)

	for g := 0; g < numBoards; g++ {
		offset := 2 + g*6
		grid := make([][]int, 5)

		for y, line := range lines[offset : offset+5] {
			row := make([]int, 5)
			numbers := rx.FindStringSubmatch(line)
			for x := 1; x <= 5; x++ {
				row[x-1], _ = strconv.Atoi(numbers[x])
			}
			grid[y] = row
		}
		boards[g] = NewBoard(grid)
	}

	if *part == "a" {
		for _, s := range calls {
			num, _ := strconv.Atoi(s)
			for _, board := range boards {

				if board.CallNumber(num) {
					fmt.Println(num * board.Score())
					return
				}
			}
		}
	} else {
		for _, s := range calls {
			num, _ := strconv.Atoi(s)
			newBoards := make([]Board, 0)

			for _, board := range boards {

				if !board.CallNumber(num) {
					newBoards = append(newBoards, board)
				}

			}
			if len(newBoards) == 0 {
				fmt.Println(boards[0].Score() * num)
				return
			}

			boards = newBoards

		}
	}
}
