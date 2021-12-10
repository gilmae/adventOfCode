package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day10.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	scores := map[rune]int{')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	expectedOpener := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	closers := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	scoresForClosers := map[rune]int{')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	score := 0
	closeScores := make([]int, 0)

	for _, line := range lines {

		stack := make([]rune, len(line))
		pointer := 0
		illegalFound := false

		for _, b := range line {

			switch b {
			case '(', '[', '{', '<':
				stack[pointer] = b
				pointer++
			case '>', '}', ']', ')':

				expected := expectedOpener[b]
				if stack[pointer-1] != expected {
					score += scores[b]
					illegalFound = true
				} else {
					pointer--
				}
			}
			if illegalFound {
				break
			}
		}

		if !illegalFound {
			closeScore := 0
			for p := pointer - 1; p >= 0; p-- {
				closer := closers[stack[p]]
				closeScore = closeScore*5 + scoresForClosers[closer]
				line += string(closer)
			}
			closeScores = append(closeScores, closeScore)

		}
	}

	sort.Sort(sort.IntSlice(closeScores))

	fmt.Println(score, closeScores[int(len(closeScores)/2)])
}
