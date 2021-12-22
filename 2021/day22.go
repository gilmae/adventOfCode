package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	vector "github.com/gilmae/adventOfCode/2021/vectors"
)

var inputFile = flag.String("inputFile", "inputs/day22.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	regexp := regexp.MustCompile("(\\w+) x=(-?\\d+)..(-?\\d+),y=(-?\\d+)..(-?\\d+),z=(-?\\d+)..(-?\\d+)")

	if *part == "a" {
		cube := make(map[vector.Coords3d]bool)

		for _, line := range lines {
			sm := regexp.FindStringSubmatch(line)
			on := sm[1] == "on"

			xmin, _ := strconv.Atoi(sm[2])
			xmax, _ := strconv.Atoi(sm[3])

			ymin, _ := strconv.Atoi(sm[4])
			ymax, _ := strconv.Atoi(sm[5])

			zmin, _ := strconv.Atoi(sm[6])
			zmax, _ := strconv.Atoi(sm[7])

			if xmin < -50 {
				xmin = -50
			}
			if ymin < -50 {
				ymin = -50
			}
			if zmin < -50 {
				zmin = -50
			}
			if xmax > 50 {
				xmax = 50
			}
			if ymax > 50 {
				ymax = 50
			}
			if zmax > 50 {
				zmax = 50
			}

			for x := xmin; x <= xmax; x++ {
				for y := ymin; y <= ymax; y++ {
					for z := zmin; z <= zmax; z++ {
						c := vector.Coords3d{x, y, z}
						cube[c] = on
					}
				}
			}

		}
		count := 0
		for _, v := range cube {
			if v {
				count++
			}
		}

		fmt.Println(count)
	} else {
		reactor := make(map[Prism]bool)

		count := uint64(0)
		for _, line := range lines {
			sm := regexp.FindStringSubmatch(line)
			on := sm[1] == "on"

			xmin, _ := strconv.Atoi(sm[2])
			xmax, _ := strconv.Atoi(sm[3])

			ymin, _ := strconv.Atoi(sm[4])
			ymax, _ := strconv.Atoi(sm[5])

			zmin, _ := strconv.Atoi(sm[6])
			zmax, _ := strconv.Atoi(sm[7])

			prism := Prism{Coords3d{xmin, ymin, zmin}, Coords3d{xmax, ymax, zmax}}

			// Find any existing prisms that this one overlaps with
			// Subtract the overlap (i.e. return rectabngular prisms that )
			for p, state := range reactor {
				overlap := p.overlap(prism)
				if overlap != nil {
					delete(reactor, p)
					slabs := p.Explode(*overlap)

					for _, s := range slabs {
						reactor[s] = state
					}

				}
			}

			reactor[prism] = on
		}

		for k, v := range reactor {
			if v {
				count += k.Volume()
			}
		}

		fmt.Println(count)

	}
}

type Coords3d [3]int

const (
	X = 0
	Y = 1
	Z = 2
)

type Prism struct {
	min, max Coords3d
}

func (p Prism) Volume() uint64 {
	volume := uint64(1)
	for d := X; d <= Z; d++ {
		volume *= uint64(p.max[d] - p.min[d] + 1)
	}
	return volume
}

func (p Prism) Explode(overlap Prism) []Prism {
	result := make([]Prism, 0)
	// Subtract the overlap from p, and explode/breakup p into multiple rectangular prisms
	// Like 3d tetris, but without the L-shape

	// If we imagined overlaps sitting in the exact centre of p, we would get back 6 prisms: 2 capstones, 2 columns, and two plugs
	// We'll treat the z-axis as the capstones, y-axis as the columns, and x-axis as the plugs

	// Obviously the overlap might not sit in the exact centre, so some testing gets done for each possible prism to see if we need it

	// First though, if overlap and p are the same, just get out. There is nothing to return
	if overlap == p {
		return result
	}
	// Top Z capstone
	if overlap.min[Z] != p.min[Z] {
		min := p.min
		max := p.max
		max[Z] = overlap.min[Z] - 1

		result = append(result, Prism{min, max})
	}

	// Bottom Z capstone
	if overlap.max[Z] != p.max[Z] {
		min := p.min
		max := p.max
		min[Z] = overlap.max[Z] + 1

		result = append(result, Prism{min, max})
	}

	// Y columns
	if overlap.min[Y] != p.min[Y] {
		min := p.min
		max := p.max

		min[Z] = overlap.min[Z]
		max[Z] = overlap.max[Z]
		max[Y] = overlap.min[Y] - 1

		result = append(result, Prism{min, max})
	}

	if overlap.max[Y] != p.max[Y] {
		min := p.min
		max := p.max

		min[Z] = overlap.min[Z]
		max[Z] = overlap.max[Z]
		min[Y] = overlap.max[Y] + 1

		result = append(result, Prism{min, max})
	}

	// X plugs, which are basically the overlap, but using the prism's x -axis to "plug the holes"
	if overlap.min[X] != p.min[X] {
		min := overlap.min
		max := overlap.max
		min[X] = p.min[X]
		max[X] = overlap.min[X] - 1

		result = append(result, Prism{min, max})
	}

	if overlap.max[X] != p.max[X] {
		min := overlap.min
		max := overlap.max
		max[X] = p.max[X]
		min[X] = overlap.max[X] + 1

		result = append(result, Prism{min, max})
	}

	return result
}

func (p Prism) overlap(o Prism) *Prism {
	overlap := Prism{}
	for d := X; d <= Z; d++ {

		if p.min[d] > o.max[d] || p.max[d] < o.min[d] {
			// No overlap
			return nil
		} else if p.max[d] <= o.max[d] && p.min[d] >= o.min[d] {
			// p is encased within o
			overlap.max[d] = p.max[d]
			overlap.min[d] = p.min[d]
		} else if o.max[d] <= p.max[d] && o.min[d] >= p.min[d] {
			// o is encased within p
			overlap.max[d] = o.max[d]
			overlap.min[d] = o.min[d]
		} else if o.min[d] <= p.min[d] && o.max[d] <= p.max[d] {
			// Partial overlap
			overlap.max[d] = o.max[d]
			overlap.min[d] = p.min[d]
		} else if p.min[d] <= o.min[d] && p.max[d] <= o.max[d] {
			// Partial overlap
			overlap.max[d] = p.max[d]
			overlap.min[d] = o.min[d]
		}
	}

	return &overlap
}
