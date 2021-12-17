package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day17.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

type Probe struct {
	x, y, dx, dy int
}

func (p *Probe) HasMissedTarget(minX, maxX, minY, maxY int) bool {
	return p.x > maxX || p.y < minY
}

func (p *Probe) HasHitTarget(minX, maxX, minY, maxY int) bool {
	return p.x >= minX && p.x <= maxX && p.y >= minY && p.y <= maxY
}

func (p *Probe) TakeTurn() {
	p.x += p.dx
	p.y += p.dy
	if p.dx > 0 {
		p.dx--
	}
	p.dy--
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	rx := regexp.MustCompile("target area: x=(-?\\d+)..(-?\\d+), y=(-?\\d+)..(-?\\d+)")
	sm := rx.FindStringSubmatch(lines[0])

	xMin, _ := strconv.Atoi(sm[1])
	xMax, _ := strconv.Atoi(sm[2])
	yMin, _ := strconv.Atoi(sm[3])
	yMax, _ := strconv.Atoi(sm[4])

	fmt.Println(xMin, xMax, yMin, yMax)
	highestHeightAchieved := 0
	hits := 0
	for dy := yMin; dy <= 0-yMin; dy++ {
		for dx := 0; dx <= xMax; dx++ {
			//startingdx, startingdy := dx, dy
			highestHeight := 0
			p := Probe{0, 0, dx, dy}
			for !p.HasMissedTarget(xMin, xMax, yMin, yMax) {
				p.TakeTurn()
				if p.dy == 0 {
					highestHeight = p.y
				}

				if p.HasHitTarget(xMin, xMax, yMin, yMax) {
					if highestHeight > highestHeightAchieved {
						//fmt.Println(startingdx, startingdy, highestHeight)
						highestHeightAchieved = highestHeight
					}
					hits++
					break
				}

			}
		}

	}
	fmt.Println(highestHeightAchieved)
	fmt.Println(hits)
}
