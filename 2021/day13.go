package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"github.com/gilmae/adventOfCode/2021/boards"
)

var inputFile = flag.String("inputFile", "inputs/day13.input", "Relative file path to use as input.")
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

	i := 0
	for lines[i] != "" {
		parts := strings.Split(lines[i], ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		board.Points[boards.Coords{x, y}] = true
		i++
	}

	i++
	board = fold(lines[i], board)
	count := 0
	for _, v := range board.Points {
		if v.(bool) {
			count++
		}
	}
	// Part A
	fmt.Println(count)

	// Part B
	for ; i < len(lines); i++ {
		board = fold(lines[i], board)
	}

	drawBoard(board)
}

func fold(line string, board *boards.Board) *boards.Board {
	foldRx := regexp.MustCompile("fold along (\\w)=(\\d+)")

	sm := foldRx.FindStringSubmatch(line)
	foldPoint, _ := strconv.Atoi(sm[2])

	if sm[1] == "y" {
		return foldy(board, foldPoint)
	} else {
		return foldX(board, foldPoint)
	}
}

func foldy(board *boards.Board, foldPoint int) *boards.Board {
	newBoard := boards.NewBoard()
	height := board.Height()

	for k, v := range board.Points {
		if k.Y < foldPoint {
			newBoard.Points[k] = v
		} else {
			newBoard.Points[boards.Coords{k.X, height - k.Y}] = v
		}
	}
	return newBoard
}

func foldX(board *boards.Board, foldPoint int) *boards.Board {
	newBoard := boards.NewBoard()
	width := board.Width()

	for k, v := range board.Points {
		if k.X < foldPoint {
			newBoard.Points[k] = v
		} else {
			newBoard.Points[boards.Coords{width - k.X, k.Y}] = v
		}
	}
	return newBoard
}

func drawBoard(board *boards.Board) {
	width := board.Width()
	height := board.Height()

	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {

			if _, ok := board.Points[boards.Coords{x, y}]; !ok {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}

		}
		fmt.Println()
	}
}
