package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Player struct {
	Position, Score int
}

var inputFile = flag.String("inputFile", "inputs/day21.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	p1 := Player{0, 0}
	p1.Position, _ = strconv.Atoi(strings.Split(lines[0], ": ")[1])
	p2 := Player{0, 0}
	p2.Position, _ = strconv.Atoi(strings.Split(lines[1], ": ")[1])

	die := 0
	rolls := 0
	won := false

	for {
		won = p1.TakeTurn(die*3 + 6)
		die += 2
		die = (die % 100) + 1
		rolls += 3
		if won {
			fmt.Printf("Player 1 won on %d after %d rolls", p2.Score*rolls, rolls)
			break
		}

		won = p2.TakeTurn(die*3 + 6)
		die += 2
		die = (die % 100) + 1

		rolls += 3
		if won {
			fmt.Printf("Player 2 won on %d after %d rolls", p1.Score*rolls, rolls)
			break
		}

	}

}

func (p *Player) TakeTurn(roll int) bool {
	p.Position = ((p.Position - 1 + roll) % 10) + 1

	p.Score += p.Position

	return p.Score >= 1000
}

func Roll(die int) int {
	nd := die + 1
	if nd > 100 {
		nd = nd % 100
	}
	return nd
}
