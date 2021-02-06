package intcode

import (
	"fmt"
	"strconv"
	"strings"
)

// Tape is a parsed set of opcodes and parameters
type Tape map[int]int

// ParseString parses a comma-delimited set of opcodes into a Tape
func ParseString(input string) Tape {
	split := strings.Split(input, ",")
	tape := make(Tape)
	for i, s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}
		tape[i] = n
	}

	return tape
}

// Copy creates a new copy of the Tape
func (t Tape) Copy() Tape {
	c := make(Tape)
	for k, v := range t {
		c[k] = v
	}
	return c
}

func (tape Tape) Process() int {
	cursor := 0
	for {
		// if *debug {
		// 	fmt.Printf("%+v\n", tape)
		// }
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
