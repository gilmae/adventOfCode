package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day05.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	bucketsRx := regexp.MustCompile("(\\[[A-Z]\\]\\s?|\\s{3}\\s?)")
	instructionRx := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")

	line := 0

	m := bucketsRx.FindAllStringSubmatch(lines[line], -1)
	stacks := make([]*Stack, len(m))
	for i := 0; i < len(m); i++ {
		stacks[i] = NewStack()
	}

	// Get stacks
	for bucketsRx.Match([]byte(lines[line])) {
		m := bucketsRx.FindAllStringSubmatch(lines[line], -1)
		for i, item := range m {
			if item[0][0] == '[' {
				stacks[i].Insert(item[0])
			}
		}
		line += 1
	}
	line += 1

	for i := 0; i < len(stacks); i++ {
		item := stacks[i].Peek()
		if item != "" {
			fmt.Printf("%s", string(item[1]))
		}

	}
	fmt.Println()
	for line < len(lines) {

		m = instructionRx.FindAllStringSubmatch(lines[line], -1)
		if len(m[0]) < 3 {
			fmt.Printf("Invalid instruction %s, %d\n", lines[line], len(m))
			fmt.Println(m[0])
		} else {
			items, _ := strconv.Atoi(m[0][1])
			src, _ := strconv.Atoi(m[0][2])
			dest, _ := strconv.Atoi(m[0][3])
			tempStack := NewStack()
			for i := 0; i < items; i++ {
				item := stacks[src-1].Pop()
				if *part == "a" {
					stacks[dest-1].Push(item)
				} else {
					tempStack.Push(item)
				}
			}
			for tempStack.Peek() != "" {
				stacks[dest-1].Push(tempStack.Pop())
			}
		}
		line += 1
	}

	for i := 0; i < len(stacks); i++ {
		item := stacks[i].Peek()
		if item != "" {
			fmt.Printf("%s", string(item[1]))
		} else {
			fmt.Printf("_")
		}

	}
	fmt.Println()
}

type Stack struct {
	items []string
}

func NewStack() *Stack {
	return &Stack{make([]string, 0)}
}

func (this *Stack) Push(i string) {
	if i != "" {
		this.items = append([]string{i}, this.items[0:]...)
	}
}

func (this *Stack) Pop() string {
	if len(this.items) == 0 {
		return ""
	}
	i := this.items[0]
	this.items = this.items[1:]
	return i
}

func (this *Stack) Peek() string {
	if len(this.items) == 0 {
		return ""
	}
	return this.items[0]
}

func (this *Stack) Insert(i string) {
	this.items = append(this.items[0:], i)
}

func (this *Stack) Print() {
	for _, i := range this.items {
		fmt.Printf("%s\t", i)
		fmt.Println()
	}
}
