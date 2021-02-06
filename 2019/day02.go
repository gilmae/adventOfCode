package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/gilmae/adventOfCode/2019/intcode"
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
	tape := intcode.ParseString(contents)

	if len(tape) == 0 {
		fmt.Println("No operations on tape")
		return
	}

	if !*partB {
		tape[1] = 12
		tape[2] = 2

		fmt.Println(tape.Process())
	} else {
		for noun := 0; noun < 100; noun += 1 {
			for verb := 0; verb < 100; verb += 1 {
				workingTape := tape.Copy()

				workingTape[1] = noun
				workingTape[2] = verb

				if 19690720 == workingTape.Process() {
					fmt.Println(100*noun + verb)
					return
				}

			}

		}
	}

}
