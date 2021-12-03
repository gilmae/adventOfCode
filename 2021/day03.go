package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day03.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	epsilon := 0
	gamma := 0
	for i := 0; i < len(lines[0]); i++ {
		bits := map[byte]int{'0': 0, '1': 0}
		for _, l := range lines {
			bits[l[i]]++
		}
		pos := len(lines[i]) - 1
		if bits['0'] > bits['1'] {
			gamma += 1 << (pos - i)
		} else {
			epsilon += 1 << (pos - i)

		}
	}

	fmt.Println(epsilon * gamma)

	oxygens := make([]string, len(lines))
	copy(oxygens, lines)

	for i := range lines[0] {
		oxygens = findOxygenReading(oxygens, i)
		if len(oxygens) == 1 {
			break
		}
	}

	carbons := make([]string, len(lines))
	copy(carbons, lines)
	for i := range lines[0] {
		carbons = findC02Reading(carbons, i)
		if len(carbons) == 1 {
			break
		}
	}
	fmt.Println(fromBinaryString(oxygens[0]) * fromBinaryString(carbons[0]))
}

func findOxygenReading(readings []string, pos int) []string {
	splits := map[byte][]string{'0': make([]string, 0), '1': make([]string, 0)}

	for _, l := range readings {
		splits[l[pos]] = append(splits[l[pos]], l)
	}
	if len(splits['0']) > len(splits['1']) {
		return splits['0']
	} else {
		return splits['1']
	}
}

func findC02Reading(readings []string, pos int) []string {
	splits := map[byte][]string{'0': make([]string, 0), '1': make([]string, 0)}

	for _, l := range readings {
		splits[l[pos]] = append(splits[l[pos]], l)
	}
	if len(splits['0']) <= len(splits['1']) {
		return splits['0']
	} else {
		return splits['1']
	}
}

func fromBinaryString(str string) int {
	result := 0
	slen := len(str) - 1
	for i, ch := range str {
		pos := slen - i
		if ch == '1' {
			result += 1 << pos
		}
	}

	return result
}
