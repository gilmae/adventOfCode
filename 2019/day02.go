package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day02.input", "Relative file path to use as input.")
var debug = flag.Bool("debug", false, "Show debug messages")
var partB = flag.Bool("partB", false, "Do Part B")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, ",")
	tape := make([]int, len(split))
	for i, s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}
		tape[i] = n
	}

	if len(tape) == 0 {
		fmt.Println("No operations on tape")
		return
	}

	if !*partB {
		tape[1] = 12
		tape[2] = 2

		fmt.Println(process(tape))
	} else {
		for noun := 0; noun < 100; noun += 1 {
			for verb := 0; verb < 100; verb += 1 {
				workingTape := make([]int, len(tape))
				copy(workingTape, tape)

				workingTape[1] = noun
				workingTape[2] = verb

				if 19690720 == process(workingTape) {
					fmt.Println(100*noun + verb)
					return
				}

			}

		}
	}

}

func process(tape []int) int {
	cursor := 0
	for {
		if *debug {
			fmt.Printf("%+v\n", tape)
		}
		if cursor >= len(tape) {
			return -1
		}
		if tape[cursor] == 99 {
			return tape[0]
		}
		op := tape[cursor]
		dloca := tape[cursor+1]
		dlocb := tape[cursor+2]
		dstloc := tape[cursor+3]
		switch op {
		case 1: // Add
			tape[dstloc] = tape[dloca] + tape[dlocb]

			cursor += 4
		case 2: //Multiply
			tape[dstloc] = tape[dloca] * tape[dlocb]
			cursor += 4
		default:
			return -1
		}

	}
}
