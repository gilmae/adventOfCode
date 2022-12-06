package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day06.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	data := lines[0]
	startOfPacket := scanForPacket(data, 4)
	fmt.Println(startOfPacket)
	fmt.Println(scanForPacket(data[startOfPacket:], 14) + startOfPacket)

}

func scanForPacket(data string, length int) int {
	for i := 0; i < len(data)-length; i++ {
		if isUnique([]rune(data)[i : i+length]) {
			return i + length
			break
		}
	}
	return -1
}

func isUnique(data []rune) bool {
	hash := make(map[rune]int)
	for _, b := range data {
		hash[b] += 1
		if hash[b] > 1 {
			return false
		}
	}
	return true
}
