// Took help from lizthegrey on this, using her readLoop technique. Didn't work until I copied her
// regexps as well, so possibly my original algorthim was hampered by that as well. _May_ re-adjust and check
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day16.input", "Relative file path to use as input.")

func main() {

	detectedPreferences := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	rx := regexp.MustCompile("Sue ([0-9]+): ((?:[a-z]+: [0-9]+(?:, )?)+)")
	rx2 := regexp.MustCompile("([a-z]+): ([0-9]+)(?:, )?")

readLoop:

	for l := range lines {
		submatches := rx.FindStringSubmatch(lines[l])
		aunt := submatches[1]
		properties := make(map[string]int)
		props := rx2.FindAllStringSubmatch(submatches[2], -1)
		for i := range props {
			athing := props[i][1]
			amount, _ := strconv.Atoi(props[i][2])
			properties[athing] = amount
		}

		for k := range detectedPreferences {
			v, ok := properties[k]

			if k == "cats" || k == "trees" {
				if ok && v <= detectedPreferences[k] {
					continue readLoop
				}
			} else if k == "pomeranians" || k == "goldfish" {
				if ok && v >= detectedPreferences[k] {
					continue readLoop
				}
			} else {
				if ok && v != detectedPreferences[k] {
					continue readLoop
				}
			}
		}

		fmt.Println(aunt)
	}

}
