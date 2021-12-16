package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gilmae/adventOfCode/2021/boards"
)

var inputFile = flag.String("inputFile", "inputs/day15.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

var cave *boards.Board

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	cave = boards.NewBoard()
	cave.Import(lines, func(c boards.Coords, value interface{}) interface{} {
		return int(value.(int32) - '0')
	})

	width := cave.Width()
	height := cave.Height()

	cave.BottomRight = boards.Coords{width, height}
	destination := boards.Coords{width, height}
	start := boards.Coords{0, 0}

	AStarCave(cave, &start, &destination)

	fmt.Println(AStarCave(cave, &start, &destination))

	biggerCave := boards.NewBoard()
	for k, v := range cave.Points {
		for r := 0; r < 5; r++ {
			for c := 0; c < 5; c++ {
				increase := r + c
				value := 1 + (v.(int)+increase-1)%9
				biggerCave.Points[boards.Coords{k.X + r*(width+1), k.Y + c*(height+1)}] = value
			}
		}
	}

	width = biggerCave.Width()
	height = biggerCave.Height()

	destination = boards.Coords{width, height}

	fmt.Println(AStarCave(biggerCave, &start, &destination))
}

func AStarCave(risks *boards.Board, src, dst *boards.Coords) int {
	work := map[boards.Coords]bool{
		*src: true,
	}
	gScore := map[boards.Coords]int{*src: 0}
	fScore := map[boards.Coords]int{*src: src.ManhattanDistance(dst)}

	path := make(map[boards.Coords]boards.Coords)

	for len(work) != 0 {
		var current boards.Coords
		currentScore := int(^uint(0)>>1) - 1
		for v := range work {
			if score, ok := fScore[v]; ok {
				if score < currentScore {
					current = v
					currentScore = score
				}
			}
		}
		// get an item off the work quete

		delete(work, current)

		if current == *dst {
			// We're there, backtrace through path to calculate score
			score := 0
			for current != *src {
				score += risks.Points[current].(int)
				current = path[current]
			}
			return score
		} else {
			for _, n := range risks.GetNeighbours(current, true) {
				tentativeScore := gScore[current] + risks.Points[n].(int)
				if previousScore, ok := gScore[n]; !ok || tentativeScore < previousScore {
					path[n] = current
					gScore[n] = tentativeScore
					fScore[n] = tentativeScore + n.ManhattanDistance(dst)
					if !work[n] {
						work[n] = true
					}
				}
			}
		}

	}

	fmt.Println(gScore)
	return -1
}
