package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day23.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	registers := map[string]int{"a": 0, "b": 0}

	if *part == "b" {
		registers["a"] = 1
	}

	pc := 0

	cmdRx := regexp.MustCompile("(\\w{3}) ([\\w\\d+-]+),?\\s?([\\d+-]+)?")
	for pc < len(lines) {
		parts := cmdRx.FindStringSubmatch(lines[pc])

		switch parts[1] {
		case "hlf":
			registers[parts[2]] /= 2
		case "tpl":
			registers[parts[2]] *= 3
		case "inc":
			registers[parts[2]] += 1
		case "jmp":
			offset, _ := strconv.Atoi(parts[2])
			pc += offset - 1
		case "jie":
			offset, _ := strconv.Atoi(parts[3])
			if registers[parts[2]]%2 == 0 {
				pc += offset - 1
			}
		case "jio":
			offset, _ := strconv.Atoi(parts[3])
			if registers[parts[2]] == 1 {
				pc += offset - 1
			}
		}
		pc++
	}

	fmt.Println(registers)
}
