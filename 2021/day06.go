package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day06.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	states := strings.Split(lines[0], ",")
	state := make(map[int]int)
	for _, s := range states {
		f, _ := strconv.Atoi(s)
		state[f] += 1
	}

	for i := 0; i < 256; i++ {
		//fmt.Println(state)
		state = tick(state)

	}
	//fmt.Println(state)
	count := 0
	for _, v := range state {
		count += v
	}
	fmt.Println(count)
}

func tick(state map[int]int) map[int]int {
	nextState := make(map[int]int)
	for state, num := range state {

		if state == 0 {
			nextState[6] += num
			nextState[8] = num
		} else {
			nextState[state-1] += num
		}
	}

	return nextState
}
