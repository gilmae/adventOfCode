package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gilmae/adventOfCode/2021/boards/bits"
)

var inputFile = flag.String("inputFile", "inputs/day16.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

var hexTobyteArray = map[byte][]byte{
	'0': []byte{'0', '0', '0', '0'},
	'1': []byte{'0', '0', '0', '1'},
	'2': []byte{'0', '0', '1', '0'},
	'3': []byte{'0', '0', '1', '1'},
	'4': []byte{'0', '1', '0', '0'},
	'5': []byte{'0', '1', '0', '1'},
	'6': []byte{'0', '1', '1', '0'},
	'7': []byte{'0', '1', '1', '1'},
	'8': []byte{'1', '0', '0', '0'},
	'9': []byte{'1', '0', '0', '1'},
	'A': []byte{'1', '0', '1', '0'},
	'B': []byte{'1', '0', '1', '1'},
	'C': []byte{'1', '1', '0', '0'},
	'D': []byte{'1', '1', '0', '1'},
	'E': []byte{'1', '1', '1', '0'},
	'F': []byte{'1', '1', '1', '1'},
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	bitData := make([]byte, len(lines[0])*4)
	for i, h := range lines[0] {
		copy(bitData[i*4:i*4+4], hexTobyteArray[byte(h)])
	}

	parser := bits.NewParser(bitData)
	pkt := parser.Parse()
	fmt.Println(sumVersionNumbers(pkt))
	fmt.Println(pkt.Value())
}

func sumVersionNumbers(p bits.Packet) int {
	switch p := p.(type) {
	case *bits.LiteralPacket:
		return p.Version
	case *bits.OperatorPacket:
		sum := p.Version
		for _, sp := range p.SubPackets {
			sum += sumVersionNumbers(sp)
		}
		return sum
	default:
		return 0
	}
}
