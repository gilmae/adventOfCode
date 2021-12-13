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

	board.PrintBoardWithShader(func(c boards.Coords, v interface{}) interface{} {
		if v == nil || !v.(bool) {
			return fmt.Sprint(" ")
		} else {
			return fmt.Sprint("#")
		}
	})
}

func fold(line string, board *boards.Board) *boards.Board {
	foldRx := regexp.MustCompile("fold along (\\w)=(\\d+)")
	newBoard := boards.NewBoard()
	sm := foldRx.FindStringSubmatch(line)
	v, _ := strconv.Atoi(sm[2])
	foldPoint := boards.Coords{int(^uint(0) >> 1), int(^uint(0) >> 1)}

	if sm[1] == "y" {
		foldPoint.Y = v
	} else {
		foldPoint.X = v
	}

	for k, v := range board.Points {
		newK := boards.Coords{k.X, k.Y}
		if k.Y >= foldPoint.Y {
			newK.Y = foldPoint.Y - (k.Y - foldPoint.Y)
		}

		if k.X >= foldPoint.X {
			newK.X = foldPoint.X - (k.X - foldPoint.X)
		}
		newBoard.Points[newK] = v
	}
	return newBoard
}
