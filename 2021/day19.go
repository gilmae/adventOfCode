package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	vector "github.com/gilmae/adventOfCode/2021/vectors"
)

var inputFile = flag.String("inputFile", "inputs/day19.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")
var debug = flag.Bool("debug", false, "Show debug messages")

type Pair [2]int
type MatchedPair struct {
	Aligned, Unaligned Pair
}
type Scanner struct {
	seen     []vector.Coords3d
	deltas   map[vector.Coords3d]Pair
	position *vector.Coords3d
	facing   *Facing // Actually the transformation to simulate all scanners facing the same direction
}

type Facing [3][3]int

func (f Facing) Invert() Facing {
	result := Facing{}
	for row := range f {
		for col := range f {
			result[col][row] = f[row][col]
		}
	}

	return result
}

func (m Facing) Transform(n Facing) Facing {
	var ret Facing
	for r := range ret {
		for c := range ret[r] {
			for i := range ret {
				ret[r][c] += m[r][i] * n[i][c]
			}
		}
	}
	return ret
}

func generateAllTransforms() [24]Facing {
	return [24]Facing{
		{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
		{{1, 0, 0}, {0, 0, -1}, {0, 1, 0}},
		{{1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
		{{1, 0, 0}, {0, 0, 1}, {0, -1, 0}},

		{{0, -1, 0}, {1, 0, 0}, {0, 0, 1}},
		{{0, 0, 1}, {1, 0, 0}, {0, 1, 0}},
		{{0, 1, 0}, {1, 0, 0}, {0, 0, -1}},
		{{0, 0, -1}, {1, 0, 0}, {0, -1, 0}},

		{{-1, 0, 0}, {0, -1, 0}, {0, 0, 1}},
		{{-1, 0, 0}, {0, 0, -1}, {0, -1, 0}},
		{{-1, 0, 0}, {0, 1, 0}, {0, 0, -1}},
		{{-1, 0, 0}, {0, 0, 1}, {0, 1, 0}},

		{{0, 1, 0}, {-1, 0, 0}, {0, 0, 1}},
		{{0, 0, 1}, {-1, 0, 0}, {0, -1, 0}},
		{{0, -1, 0}, {-1, 0, 0}, {0, 0, -1}},
		{{0, 0, -1}, {-1, 0, 0}, {0, 1, 0}},

		{{0, 0, -1}, {0, 1, 0}, {1, 0, 0}},
		{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}},
		{{0, 0, 1}, {0, -1, 0}, {1, 0, 0}},
		{{0, -1, 0}, {0, 0, -1}, {1, 0, 0}},

		{{0, 0, -1}, {0, -1, 0}, {-1, 0, 0}},
		{{0, -1, 0}, {0, 0, 1}, {-1, 0, 0}},
		{{0, 0, 1}, {0, 1, 0}, {-1, 0, 0}},
		{{0, 1, 0}, {0, 0, -1}, {-1, 0, 0}},
	}
}

func (s *Scanner) CalculateDeltas() {
	for i, a := range s.seen {
		for j, b := range s.seen {
			if i == j {
				continue
			}
			s.deltas[a.Difference(b)] = Pair{i, j}
		}
	}
}

func (s Scanner) ActualPositions() []vector.Coords3d {
	actualPositions := make([]vector.Coords3d, len(s.seen))
	for i, v := range s.seen {
		actualPositions[i] = v.Transform(*s.facing).Add(*s.position)
	}
	return actualPositions
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}

	beacons := make(map[vector.Coords3d]bool)

	allFacings := generateAllTransforms()
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	scanners := make([]*Scanner, 0)
	var scanner *Scanner
	for _, line := range lines {
		if len(line) == 0 {
			scanner.CalculateDeltas()
			continue
		}
		if line[0:3] == "---" {
			scanner = &Scanner{deltas: make(map[vector.Coords3d]Pair)}
			scanners = append(scanners, scanner)
			continue
		}
		c := vector.Coords3d{}
		parts := strings.Split(line, ",")
		c.X, _ = strconv.Atoi(parts[0])
		c.Y, _ = strconv.Atoi(parts[1])
		c.Z, _ = strconv.Atoi(parts[2])

		scanner.seen = append(scanner.seen, c)
	}
	scanner.CalculateDeltas()

	scanners[0].position = &vector.Coords3d{0, 0, 0}
	scanners[0].facing = &allFacings[0] // The identity facing :-)

normalise:
	// Iterate through unnormalised scanners (ones that do not have a position and facing)
	// 		Iterate through all normalised scanners (ones that have a position and a facing)
	// 			Iterate through the unnormalised scanner's deltas
	// 				Iterate through each of the facings
	// 					Iterate through normalised scanners deltas, map them with the facing. If deltas match, add to a matched list
	// 					If enough matches are found, this facing is the correct one. Normalise the unnormalised scanner
	//					Then restart all the loops again
	for {
		for scanNum, unaligned := range scanners {
			if unaligned.facing != nil && unaligned.position != nil {
				continue
			}
			if *debug {
				fmt.Printf("Aligning scanner %d\n", scanNum)
			}

			for alignNum, aligned := range scanners {

				if aligned.facing == nil || aligned.position == nil {
					continue
				}

				if *debug {
					fmt.Printf("Attempt to align scanner %d with scanner %d\n", scanNum, alignNum)
				}
				deltasMatched := make(map[Facing][]MatchedPair)
				for d, p := range aligned.deltas {
					for _, facing := range allFacings {

						if matchedPair, ok := unaligned.deltas[d.Transform(facing)]; ok {
							deltasMatched[facing] = append(deltasMatched[facing], MatchedPair{Aligned: p, Unaligned: matchedPair})
							break
						}
					}
				}

				for facing, pairs := range deltasMatched {
					uniqueMatches := make(map[int]bool)
					for _, p := range pairs {
						uniqueMatches[p.Aligned[0]] = true
						uniqueMatches[p.Aligned[1]] = true

					}

					// In the real code I need to test if we got 12 unique matches. If we didn't, bail
					if len(uniqueMatches) < 12 {
						continue
					}

					mapped := aligned.facing.Transform(facing)
					unaligned.facing = &mapped
					pair := pairs[0]

					// Theory u[0] is b[0] and u[1] = b[1] and a[0] = b[0] and a[1] is b[1]
					/*
						If true

						a[0] is the vector from aligned's origin to b[0]
							-> aOri + a[0] == b[0]
						    -> 0,0 + 0,2 == 0,2
						a[1] is the vector from aligned's origin to b[1]
							-> aOri + a[1] = b[1]
							-> 0,0 + 4,1 = 4,1

						u[0] is the vector from unaligned's origin to b[0]
							-> uOri + u[0] == b[0]
							-> ? + -5,0 = 0,2
							-> ? = 0,2 - -5,0
							-> 5,2

						u[1] is the vector from unaligned's origin to b[1]
							-> uOri + u[1] == b[1]
							-> ? + -1,-1 = 4,1
							-> ? = 4,1 - -1,-1
							-> 5,2

						If we add a[0] to aOri (i.e. b[0]) and subtract u[0] (the vector from uOri to b[0]) we should be at uOri

						a[0] + aOri - u[o]
						= b[0] - aOri + aOri - (b[0] - uOri)
						= b[0] - aOri + aOri - b[0] + uOri
						= uOri

						if we then add the vector u[1] we should be at b[1]
						if we also add vector a[1] to aOri, we should be at b[1]

					*/
					position := vector.Coords3d{}
					for attempt := 1; attempt <= 2; attempt++ {
						beacon0 := aligned.seen[pair.Aligned[0]].Transform(*aligned.facing).Add(*aligned.position)
						beacon1 := aligned.seen[pair.Aligned[1]].Transform(*aligned.facing).Add(*aligned.position)

						unalignedVectorToBeacon0 := unaligned.seen[pair.Unaligned[0]].Transform(*unaligned.facing)
						unalignedVectorToBeacon1 := unaligned.seen[pair.Unaligned[1]].Transform(*unaligned.facing)

						position = beacon0.Difference(unalignedVectorToBeacon0)

						if position.Add(unalignedVectorToBeacon1) == beacon1 {
							break
						}

						position = beacon1.Difference(unalignedVectorToBeacon0)
						if position.Add(unalignedVectorToBeacon1) == beacon0 {
							break
						}

						if attempt == 2 {
							fmt.Printf("Could not find a working position for origin for scanner %d\n", scanNum)
							fmt.Println(beacon0, beacon1)
							return
						}

						mapped = aligned.facing.Transform(facing.Invert())
					}

					unaligned.position = &position
					if *debug {
						fmt.Printf("Aligned scanner %d with scnner %d\n", scanNum, alignNum)
					}
					continue normalise

				}

			}
		}
		break
	}

	for scanNum, scanner := range scanners {
		if scanner.facing != nil && scanner.position != nil {
			if *debug {
				fmt.Printf("Scanner %d aligned and at position %+v\n", scanNum, scanner.position)
			}
			for _, b := range scanner.ActualPositions() {
				beacons[b] = true
			}
		} else {
			if *debug {
				fmt.Printf("Scanner %d is still unaligned.\n", scanNum)
			}
		}
	}

	fmt.Println(len(beacons))
	largestDistance := 0
	for i, a := range scanners {
		for j, b := range scanners {
			if i == j {
				continue
			}
			distance := a.position.ManhattanDistance(b.position)
			if distance > largestDistance {
				largestDistance = distance
			}
		}
	}

	fmt.Println(largestDistance)

}

// Guess 1: 329
// Guess 2: 342 (this time the sample data had the correct answer)
