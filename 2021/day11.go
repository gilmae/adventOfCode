package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gilmae/adventOfCode/2021/boards"
)

var inputFile = flag.String("inputFile", "inputs/day11.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)

	board := boards.NewBoard()
	lines := strings.Split(contents, "\n")
	board.Import(lines, func(c boards.Coords, v interface{}) interface{} {
		i, _ := strconv.Atoi(string(v.(rune)))
		return i
	})

	if *part == "a" {
		flashCount := 0
		var fc int
		for i := 0; i < 100; i++ {

			board, fc = doOctopusEnergyCycle(board)
			flashCount += fc
		}
		fmt.Println(flashCount)
		board.PrintBoard()

	} else {
		step := 0
		for !checkAllFlashed(board) {
			board, _ = doOctopusEnergyCycle(board)
			step += 1
		}

		fmt.Printf("Synchronised flash at step %d\n", step)
	}

}

func doOctopusEnergyCycle(board *boards.Board) (*boards.Board, int) {
	flashed := make(map[boards.Coords]bool)
	flashCount := 0
	board = energise(board)
	flashc := processFlash(board, &flashed)
	flashCount += flashc
	for flashc > 0 {
		flashc = processFlash(board, &flashed)
		flashCount += flashc
	}

	rest(board)
	return board, flashCount
}

func energise(board *boards.Board) *boards.Board {
	maxX, maxY := board.Width(), board.Height()

	newBoard := boards.NewBoard()

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			c := boards.Coords{x, y}
			newBoard.Points[c] = board.Points[c].(int) + 1
		}
	}

	return newBoard
}

func processFlash(b *boards.Board, flashed *map[boards.Coords]bool) int {
	flashers := make([]boards.Coords, 0)
	for c, v := range b.Points {

		if v.(int) >= 10 {
			if _, ok := (*flashed)[c]; !ok {
				flashers = append(flashers, c)
				(*flashed)[c] = true
			}
		}
	}

	for _, c := range flashers {
		for _, n := range b.GetNeighbours(c, false) {
			b.Points[n] = b.Points[n].(int) + 1
		}
	}

	return len(flashers)
}

func rest(b *boards.Board) {
	for k, v := range b.Points {
		if v.(int) > 9 {
			b.Points[k] = 0
		}
	}
}

func checkAllFlashed(b *boards.Board) bool {
	sum := 0
	for _, v := range b.Points {
		sum += v.(int)
	}
	return sum == 0
}
