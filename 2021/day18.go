package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day18.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")
var debug = flag.Bool("debug", false, "Show debug messages")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	parser := NewSnailfishParser()

	numbers := make([]*SnailfishNumber, len(lines))
	for i, line := range lines {
		numbers[i] = parser.Parse([]byte(line))
	}
	sum := numbers[0]

	if len(numbers) > 1 {
		for i := 1; i < len(numbers); i++ {
			sum = Sum(sum, numbers[i])
		}

	}

	fmt.Println(sum.Magnitude())
	if *part == "b" {
		largestMagnitude := 0
		for i, _ := range lines {
			for j, _ := range lines {
				if i == j {
					continue
				}
				a := parser.Parse([]byte(lines[i]))

				b := parser.Parse([]byte(lines[j]))

				magnitude := Sum(a, b).Magnitude()
				if magnitude > largestMagnitude {
					largestMagnitude = magnitude
				}
			}
		}
		fmt.Println(largestMagnitude)
	}

}

func Sum(a, b *SnailfishNumber) *SnailfishNumber {
	pair := SnailfishNumber{
		Left:      nil,
		Right:     nil,
		LeftPair:  a,
		RightPair: b,
	}

	pair.LeftPair.IncrementDepth()
	pair.LeftPair.Parent = &pair
	pair.RightPair.IncrementDepth()
	pair.RightPair.Parent = &pair
	if *debug {
		fmt.Printf("After add:\t\t%s\n", pair)
	}
	pair.Reduce()
	return &pair
}

func (p *SnailfishNumber) Reduce() {
	for {
		if exploded := p.explode(); exploded {
			if *debug {
				fmt.Printf("After explode:\t\t%s\n", p)
			}
			continue
		}
		if split := p.split(); !split {
			break
		} else {
			if *debug {
				fmt.Printf("After split:\t\t%s\n", p)
			}

		}

	}
}

func (p *SnailfishNumber) Magnitude() int {
	value := 0
	if p.LeftPair != nil {
		value += 3 * p.LeftPair.Magnitude()
	} else {
		value += 3 * *p.Left
	}

	if p.RightPair != nil {
		value += 2 * p.RightPair.Magnitude()
	} else {
		value += 2 * *p.Right
	}

	return value
}

func (p *SnailfishNumber) explode() bool {
	// If we are at depth of 4+ and have no nested pairs, explode
	if p.Depth >= 4 && p.Left != nil && p.Right != nil {
		p.pushValueLeft(*p.Left)
		p.pushValueRight(*p.Right)

		// Remove myself
		zero := 0
		if p.Parent.LeftPair == p {
			p.Parent.LeftPair = nil

			p.Parent.Left = &zero
		} else {
			p.Parent.RightPair = nil
			p.Parent.Right = &zero
		}
		return true
	}
	if p.LeftPair != nil {
		if exploded := p.LeftPair.explode(); exploded {
			return exploded
		}
	}
	if p.RightPair != nil {
		if exploded := p.RightPair.explode(); exploded {
			return exploded
		}
	}
	return false
}

func (p *SnailfishNumber) pushValueLeft(value int) {
	parent := p.Parent
	if parent == nil {
		return
	}
	if parent.RightPair == p {
		if parent.Left != nil {
			*parent.Left += value
		} else {
			parent.LeftPair.pushValueDownToFindNeightbourToTheLeft(value)
		}
	} else if parent.LeftPair == p {
		parent.pushValueLeft(value)
	}
}

func (p *SnailfishNumber) pushValueDownToFindNeightbourToTheLeft(value int) {
	if p.Right != nil {
		*p.Right += value
	} else {
		p.RightPair.pushValueDownToFindNeightbourToTheLeft(value)
	}
}

func (p *SnailfishNumber) pushValueRight(value int) {
	parent := p.Parent
	if parent == nil {
		return
	}
	if parent.LeftPair == p {
		if parent.Right != nil {
			*parent.Right += value
		} else {
			parent.RightPair.pushValueDownToFindNeightbourToTheRight(value)
		}
	} else if parent.RightPair == p {
		parent.pushValueRight(value)
	}
}

func (p *SnailfishNumber) pushValueDownToFindNeightbourToTheRight(value int) {
	if p.Left != nil {
		*p.Left += value
	} else {
		p.LeftPair.pushValueDownToFindNeightbourToTheRight(value)
	}
}

func (p *SnailfishNumber) split() bool {
	if p.Left != nil && *p.Left > 9 {
		left := *p.Left / 2
		right := *p.Left - left

		p.LeftPair = &SnailfishNumber{Left: &left, Right: &right, Parent: p, Depth: p.Depth + 1}
		p.Left = nil
		return true
	}

	if p.LeftPair != nil {
		if split := p.LeftPair.split(); split {
			return split
		}
	}
	if p.Right != nil && *p.Right > 9 {
		left := *p.Right / 2
		right := *p.Right - left

		p.RightPair = &SnailfishNumber{Left: &left, Right: &right, Parent: p, Depth: p.Depth + 1}
		p.Right = nil
		return true
	}

	if p.RightPair != nil {
		if split := p.RightPair.split(); split {
			return split
		}
	}
	return false
}

func (p *SnailfishNumber) IncrementDepth() {
	p.Depth++
	if p.LeftPair != nil {
		p.LeftPair.IncrementDepth()
	}
	if p.RightPair != nil {
		p.RightPair.IncrementDepth()
	}
}

type SnailfishNumber struct {
	Left, Right         *int
	LeftPair, RightPair *SnailfishNumber
	Depth               int
	Parent              *SnailfishNumber
}

type SnailFishParser struct {
	data  []byte
	ip    int
	stack []*SnailfishNumber
}

func NewSnailfishParser() *SnailFishParser {
	s := SnailFishParser{data: nil, ip: 0}
	return &s
}

func (p *SnailFishParser) Parse(data []byte) *SnailfishNumber {
	p.data = data
	p.ip = 0
	if p.peekChar() == '[' {

		return p.parseNested()

	} else {
		return nil
	}
}

func (p *SnailfishNumber) String() string {
	var out bytes.Buffer
	out.WriteString("[")

	if p.LeftPair != nil {
		out.WriteString(p.LeftPair.String())
	} else {
		out.WriteString(fmt.Sprintf("%d", *p.Left))
	}
	out.WriteString(",")
	if p.RightPair != nil {
		out.WriteString(p.RightPair.String())
	} else {
		out.WriteString(fmt.Sprintf("%d", *p.Right))
	}
	out.WriteString("]")
	return out.String()
}

func (p *SnailFishParser) parseNested() *SnailfishNumber {
	n := SnailfishNumber{}
	n.Depth = len(p.stack)
	if len(p.stack) != 0 {
		n.Parent = p.stack[len(p.stack)-1]
	}
	p.stack = append(p.stack, &n)
	p.ip++
	if p.peekChar() == '[' {
		n.LeftPair = p.parseNested()
	} else {
		n.Left = p.parseRegular()
	}
	for p.peekChar() == ',' || p.peekChar() == ']' {
		p.ip++
	}
	if p.peekChar() == '[' {
		n.RightPair = p.parseNested()
	} else {
		n.Right = p.parseRegular()
	}
	p.stack = p.stack[:len(p.stack)-1]
	return &n

}

func (p *SnailFishParser) parseRegular() *int {

	str := make([]byte, 0)
	for {
		ch := p.getChar()
		if ch != ',' && ch != ']' {
			str = append(str, ch)
		} else {
			break
		}
	}
	v, _ := strconv.Atoi(string(str))

	return &v
}

func (p *SnailFishParser) getChar() byte {
	b := p.data[p.ip]
	p.ip++
	return b
}
func (p *SnailFishParser) peekChar() byte {
	return p.data[p.ip]
}
