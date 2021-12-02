package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day19.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	tokens := getTokens(lines)

	molocule := lines[len(lines)-1]

	fmt.Println(len(partA(tokens, molocule)))
	fmt.Println(partB(tokens, molocule))
}

func partA(tokens map[string][]string, molocule string) map[string]bool {
	newMolocules := make(map[string]bool)

	atomRx := regexp.MustCompile("([A-Z]{1}[a-z]*)")
	sm := atomRx.FindAllStringIndex(molocule, -1)
	for i := range sm {
		atom := molocule[sm[i][0]:sm[i][1]]
		for j := range tokens[atom] {
			newMolocules[molocule[:sm[i][0]]+tokens[atom][j]+molocule[sm[i][1]:]] = true

		}

	}

	return newMolocules
}

func partB(tokens map[string][]string, molocule string) int {
	// currentMolocule := molocule
	// replacements, replacementToToken := getReplacementsByLength(tokens)
	// steps := 0
	// // Longest first
	// sort.Slice(replacements, func(i, j int) bool {
	// 	return len(replacements[i]) > len(replacements[j])
	// })
	// for currentMolocule != "e" {
	// 	for _, r := range replacements {
	// 		if strings.Contains(currentMolocule, r) {
	// 			currentMolocule = strings.Replace(currentMolocule, r, replacementToToken[r], 1)
	// 			steps++
	// 			fmt.Printf("Step %d: Replaced %s with %s, leaving %s\n", steps, r, replacementToToken[r], currentMolocule)
	// 			break
	// 		}
	// 	}
	// }
	// return steps

	// This is the fucking cheese solution
	atomRx := regexp.MustCompile("([A-Z]{1}[a-z]*)")
	atomLocs := atomRx.FindAllStringIndex(molocule, -1)
	atoms := 0
	brackets := 0
	commas := 0
	for _, l := range atomLocs {
		atoms++
		switch molocule[l[0]:l[1]] {
		case "Rn":
			brackets++
		case "Ar":
			brackets++
		case "Y":
			commas++
		}
	}

	return atoms - brackets - commas*2 - 1
}

func getTokens(lines []string) map[string][]string {
	rx := regexp.MustCompile("(\\w+) => (\\w+)")
	tokens := make(map[string][]string)
	i := 0
	for lines[i] != "" {
		sm := rx.FindStringSubmatch(lines[i])
		_, ok := tokens[sm[1]]
		if !ok {
			tokens[sm[1]] = make([]string, 0)
		}
		tokens[sm[1]] = append(tokens[sm[1]], sm[2])
		i++
	}

	return tokens
}

func getReplacementsByLength(tokens map[string][]string) ([]string, map[string]string) {
	replacementToToken := make(map[string]string)
	replacements := make([]string, 0)

	for k, v := range tokens {
		replacements = append(replacements, v...)
		for j := range v {
			replacementToToken[v[j]] = k
		}
	}

	return replacements, replacementToToken
}
