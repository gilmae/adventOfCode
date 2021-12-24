package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day02.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

type ALU struct {
	code      []string
	ip        int
	registers map[string]int
}

func NewALU(code []string) *ALU {
	return &ALU{code: code, ip: 0, registers: map[string]int{"w": 0, "x": 0, "y": 0, "z": 0}}
}

func (a *ALU) Process() {
	for a.ip < len(a.code) {
		parts := strings.Split(a.code[a.ip], " ")

		switch parts[0] {
		case "inp":
			a.registers[parts[1]] = getInput()
		case "add":
			a.registers[parts[1]] += a.parseVariable(parts[2])
		case "mul":
			a.registers[parts[1]] *= a.parseVariable(parts[2])
		case "div":
			a.registers[parts[1]] /= a.parseVariable(parts[2])
		case "mod":
			a.registers[parts[1]] %= a.parseVariable(parts[2])
		case "eql":
			var result int
			if a.registers[parts[1]] == a.registers[parts[2]] {
				result = 1
			} else {
				result = 0
			}

			a.registers[parts[1]] = result
		}
		a.ip++
		fmt.Println(a.registers)
	}
}

func (a *ALU) parseVariable(variable string) int {
	n, err := strconv.Atoi(variable)
	if err == nil {
		return n
	}

	if n, ok := a.registers[variable]; ok {
		return n
	}

	fmt.Println("illegal variable: ", variable)
	return -1

}

func getInput() int {
	return 1
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	a := NewALU(lines)
	a.Process()

	fmt.Println(a.registers)

}
