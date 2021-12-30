package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day04.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	rx := regexp.MustCompile("(\\d+)\\[(\\w+)]")
	sectorIdSum := 0
	for _, line := range lines {
		nameParts := strings.Split(line, "-")
		chars := make(map[rune]int)
		for p := 0; p < len(nameParts)-1; p++ {
			for _, ch := range nameParts[p] {
				chars[ch] += 1
			}
		}
		sm := rx.FindStringSubmatch(nameParts[len(nameParts)-1])
		sectorId, _ := strconv.Atoi(sm[1])
		checksum := sm[2]

		if checksum == getChecksum(chars) {

			sectorIdSum += sectorId
			fmt.Printf("%s - %d\n", decryptRoomName(strings.Join(nameParts[:len(nameParts)-1], "-"), sectorId), sectorId)
		}

		// good luckm search manually, search for 'north'
	}
	fmt.Println(sectorIdSum)
}

func getChecksum(chars map[rune]int) string {
	pairList := PairList{}
	for k, v := range chars {
		pairList = append(pairList, Pair{k, v})
	}

	sort.Sort(sort.Reverse(sort.Interface(pairList)))
	checkSum := ""
	for _, p := range pairList[0:5] {

		checkSum += string(p.key)
	}
	return checkSum
}

func decryptRoomName(roomName string, sectorId int) string {
	plaintext := ""
	for _, ch := range roomName {
		if ch == '-' {
			plaintext += string(' ')
		} else {
			plaintext += string(byte(97 + ((int(ch) - 97 + sectorId) % 26)))
		}
	}
	return plaintext
}

type Pair struct {
	key   rune
	value int
}

type PairList []Pair

func (p PairList) Len() int      { return len(p) }
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool {
	if p[i].value < p[j].value {
		return true
	} else if p[i].value == p[j].value && p[i].key > p[j].key {
		return true
	}
	return false
}
