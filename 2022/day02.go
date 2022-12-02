package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day02.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

const (
	OppRock    = 'A'
	OppPaper   = 'B'
	OppSissors = 'C'

	Rock    = 'X'
	Paper   = 'Y'
	Sissors = 'Z'

	Lose = 'X'
	Draw = 'Y'
	Win  = 'Z'
)

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	totalScore := 0
	totalScoreB := 0
	for _, line := range lines {
		if len(line) < 3 {
			fmt.Errorf("invalid line: %s", line)
		}
		opp := line[0]
		me := line[2]
		result := line[2]

		totalScore = totalScore + getScore(me, opp)
		totalScoreB = totalScoreB + getScore(getPlay(opp, result), opp)
	}
	fmt.Println(totalScore)
	fmt.Println(totalScoreB)

}

func getPlay(opp byte, result byte) byte {
	plays := map[byte]map[byte]byte{

		Lose: map[byte]byte{ // Lose
			OppRock: Sissors, OppPaper: Rock, OppSissors: Paper,
		},
		Draw: map[byte]byte{ // Draw
			OppRock: Rock, OppPaper: Paper, OppSissors: Sissors,
		},
		Win: map[byte]byte{ // Win
			OppRock: Paper, OppPaper: Sissors, OppSissors: Rock,
		},
	}

	return plays[result][opp]
}

func getScore(me byte, opponent byte) int {
	typeScores := map[byte]map[byte]int{

		OppRock: map[byte]int{
			Rock: 4, Paper: 8, Sissors: 3,
		},
		OppPaper: map[byte]int{
			Rock: 1, Paper: 5, Sissors: 9,
		},
		OppSissors: map[byte]int{
			Rock: 7, Paper: 2, Sissors: 6,
		},
	}
	return typeScores[opponent][me]
}
