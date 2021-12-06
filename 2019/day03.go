package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gilmae/adventOfCode/2019/graphs"
)

var inputFile = flag.String("inputFile", "inputs/day03.input", "Relative file path to use as input.")

//var debug = flag.Bool("debug", false, "Show debug messages")
var partB = flag.String("part", "a", "Which part")

type wirePoint struct {
	wire, distance int
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	board := graphs.NewBoard()
	var length int

	for i, line := range lines {
		length = 0
		shader := func(c graphs.Coords, value interface{}) interface{} {
			if value == nil {
				return map[int]int{i + 1: length}
			}

			v := value.(map[int]int)
			if _, ok := v[i+1]; !ok {
				v[i+1] = length
			}
			return v
		}

		points := make([]graphs.Coords, 0)
		steps := strings.Split(line, ",")
		start := graphs.Coords{0, 0}
		for _, step := range steps {
			distance, _ := strconv.Atoi(step[1:])
			end := graphs.Coords{start.X, start.Y}
			switch step[0] {
			case 'U':
				end.Y += distance
			case 'D':
				end.Y -= distance
			case 'L':
				end.X -= distance
			case 'R':
				end.X += distance
			}

			l := graphs.NewLine(start, end)

			segment := l.GetPoints()
			for _, c := range segment {
				board.DrawPoint(c, shader)
				length++
			}
			length--
			points = append(points, segment...)
			start = end

		}
		board.DrawPoints(points, shader)
	}

	//board.Draw()

	minDistance := 100000000
	minLatency := 1000000000
	for k, v := range board.Points {
		if k.X == 0 && k.Y == 0 {
			continue
		}
		value := v.(map[int]int)
		if value == nil {
			continue
		} else if len(value) > 1 {
			distance := k.ManhattanDistance(graphs.Coords{0, 0})
			if distance < minDistance {
				minDistance = distance
			}
			latency := 0
			for _, l := range value {
				latency += l
			}
			if latency < minLatency {
				minLatency = latency
			}
		}
	}
	fmt.Println(minDistance, minLatency)

}
