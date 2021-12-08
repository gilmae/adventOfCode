package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day08.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	unique_digits := 0

	for _, line := range lines {
		parts := strings.Split(line, " | ")
		for _, output := range strings.Split(parts[1], " ") {
			if len(output) == 2 || len(output) == 3 || len(output) == 4 || len(output) == 7 {
				unique_digits++
			}
		}
	}

	result := 0
	for _, line := range lines {
		signalsByNumOnBits := make(map[int][]string)
		parts := strings.Split(line, " | ")
		for _, output := range strings.Split(parts[0], " ") {
			if _, ok := signalsByNumOnBits[len(output)]; !ok {
				signalsByNumOnBits[len(output)] = []string{output}
			} else {
				signalsByNumOnBits[len(output)] = append(signalsByNumOnBits[len(output)], output)
			}

		}
		deciphered := decipherLine(signalsByNumOnBits)
		plain := decrypt(deciphered, parts[1])
		fmt.Println(plain)
		result += plain

	}

	fmt.Println(result)
}

func decipherLine(signals map[int][]string) map[rune]rune {
	//outputs := make(map[rune]rune)

	deciphered := make(map[rune]rune)
	seven := make(map[rune]bool)
	two := make(map[rune]bool)

	// Map the segments used in the seven and the two digits
	for _, r := range signals[3][0] {
		seven[r] = true
	}
	for _, r := range signals[2][0] {
		two[r] = true
	}

	// Only 7 uses three segment, only 1 uses 2 segments. The segment not in the 1 is the 'a' segment
	// therefore the segment we find in the 7 map but not in the 1 map must be 'a'in the 1 is the 'a' segment
	for k, _ := range seven {
		if _, ok := two[k]; !ok {
			deciphered[k] = 'a'
			break
		}
	}

	// Count the number of times a segment is used across all 5 segment signals
	five_segments := make(map[rune]int)
	for _, poss := range signals[5] {
		for _, r := range poss {
			five_segments[r]++
		}

	}

	// Count the number of times a segment is used across the four segment signal
	four_segments := make(map[rune]int)
	for _, r := range signals[4][0] {
		four_segments[r]++
	}

	for k, v := range five_segments {

		if v == 3 {

			// The 'd' segment is used in all 5 segment digits
			// and is also used in the four
			// The 'g' segment is also used in all 5 segment digits
			// but is not used in either the four or the seven
			if four_segments[k] == 1 {
				deciphered[k] = 'd'
			} else if four_segments[k] == 0 && !seven[k] {
				deciphered[k] = 'g'
			}
		}

		// The 'b' segment is used only once in all 5 segment digits, and is also used in the four
		if v == 1 && four_segments[k] == 1 {
			deciphered[k] = 'b'
		}

		// The 'e' segment is also used only once in the 5 segment digits and is not used in the four
		if _, ok := four_segments[k]; !ok && v == 1 {
			deciphered[k] = 'e'
		}
	}

	six_segments := make(map[rune]int)
	for _, poss := range signals[6] {
		for _, r := range poss {
			six_segments[r]++
		}

	}

	// The 'c' segment is used in only 2 of the  6 segment digits and is also used in the two
	// The 'f' segment is used in all 3 of the 6 segment digits and is also used in the two
	for k, v := range six_segments {
		if v == 2 && two[k] {
			deciphered[k] = 'c'
		} else if v == 3 && two[k] {
			deciphered[k] = 'f'
		}
	}
	return deciphered

}

func decrypt(decipheredSignals map[rune]rune, cypher string) int {
	digitsToSegments := map[string]int{
		"abcefg":  0,
		"cf":      1,
		"acdeg":   2,
		"acdfg":   3,
		"bcdf":    4,
		"abdfg":   5,
		"abdefg":  6,
		"acf":     7,
		"abcdefg": 8,
		"abcdfg":  9,
	}
	result := 0
	for _, c := range strings.Split(cypher, " ") {

		plain := make([]rune, 0)
		for _, r := range c {
			plain = append(plain, decipheredSignals[r])
		}

		sort.Sort(RuneSlice(plain))

		digit := digitsToSegments[string(plain)]
		result = result*10 + digit
	}

	return result
}

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
