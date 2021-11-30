package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func copy(m map[int]bool) map[int]bool {
	newMap := make(map[int]bool)
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

func fitIntoContainers(remainder int, containers []int, seen map[int]int, path map[int]bool) {
	// THis should probaly be memoized a bit
	// Also, it should probably keep a runnning count of min to save a loop in main()
	if remainder == 0 {
		return
	}
	for i := range containers {
		_, ok := path[i]
		if ok {
			continue
		}
		c := containers[i]
		if path == nil {
			path = make(map[int]bool, 0)
		}

		if remainder == c {

			key := 0
			for k, _ := range path {
				key += 1 << k
			}
			key += 1 << i
			_, ok := seen[key]
			if ok {
				continue
			} else {
				seen[key] = len(path) + 1

			}
		} else if remainder >= c {
			newPath := copy(path)
			newPath[i] = true
			fitIntoContainers(remainder-c, containers, seen, newPath)
		}
	}

}

var inputFile = flag.String("inputFile", "inputs/day17.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	containers := make([]int, len(lines))
	for i := range lines {
		containers[i], _ = strconv.Atoi(lines[i])
	}
	seen := make(map[int]int)
	fitIntoContainers(150, containers, seen, nil)
	fmt.Println(len(seen))

	min := len(containers) + 1
	for _, v := range seen {
		if v < min {
			min = v
		}
	}
	count := 0
	for _, v := range seen {
		if v == min {
			count++
		}
	}

	fmt.Println(count)
}
