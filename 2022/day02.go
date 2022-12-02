package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day02.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

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
		parts := strings.Split(line, " ")
		// Part A
		opp := parts[0][0]
		me := parts[1][0]
		result := parts[1][0]

		totalScore = totalScore + getScore(me, opp)
		totalScoreB = totalScoreB + getScore(getPlay(opp, result), opp)
	}
	fmt.Println(totalScore)
	fmt.Println(totalScoreB)

}

func getPlay(opp byte, result byte) byte {
	plays := map[byte]map[byte]byte{

		'X': map[byte]byte{ // Lose
			'A': 'Z', 'B': 'X', 'C': 'Y',
		},
		'Y': map[byte]byte{ // Draw
			'A': 'X', 'B': 'Y', 'C': 'Z',
		},
		'Z': map[byte]byte{ // Win
			'A': 'Y', 'B': 'Z', 'C': 'X',
		},
	}

	return plays[result][opp]
}

func getScore(me byte, opponent byte) int {
	typeScores := map[byte]map[byte]int{

		'A': map[byte]int{
			'X': 4, 'Y': 8, 'Z': 3,
		},
		'B': map[byte]int{
			'X': 1, 'Y': 5, 'Z': 9,
		},
		byte('C'): map[byte]int{
			'X': 7, 'Y': 2, 'Z': 6,
		},
	}
	return typeScores[opponent][me]
}
