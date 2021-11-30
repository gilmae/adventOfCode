package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

var inputFile = flag.String("inputFile", "inputs/day12.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	fmt.Println(contents)
}
