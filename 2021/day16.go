package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Packet interface {
	PacketType() string
	Value() int
}

type LiteralPacket struct {
	Version, Literal int
}

func (l *LiteralPacket) PacketType() string { return "LITERAL" }

func (l *LiteralPacket) Value() int { return l.Literal }

type OperatorPacket struct {
	Version    int
	Type       int
	SubPackets []Packet
}

func (o *OperatorPacket) PacketType() string { return "OPERATOR" }
func (o *OperatorPacket) Value() int {

	switch o.Type {
	case 0:
		value := 0
		for _, sb := range o.SubPackets {
			value += sb.Value()
		}
		return value
	case 1:
		value := 1
		for _, sb := range o.SubPackets {
			value *= sb.Value()
		}
		return value
	case 2:
		values := make([]int, len(o.SubPackets))
		for i, sb := range o.SubPackets {
			values[i] = sb.Value()
		}

		sort.Sort(sort.IntSlice(values))
		return values[0]
	case 3:
		values := make([]int, len(o.SubPackets))
		for i, sb := range o.SubPackets {
			values[i] = sb.Value()
		}

		sort.Sort(sort.Reverse(sort.IntSlice(values)))
		return values[0]
	case 5:
		if o.SubPackets[0].Value() > o.SubPackets[1].Value() {
			return 1
		} else {
			return 0
		}
	case 6:
		if o.SubPackets[0].Value() < o.SubPackets[1].Value() {
			return 1
		} else {
			return 0
		}
	case 7:
		if o.SubPackets[0].Value() == o.SubPackets[1].Value() {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
	return 0
}

type Parser struct {
	data []byte
	pc   int
}

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

	bits := make([]byte, len(lines[0])*4)
	for i, h := range lines[0] {
		copy(bits[i*4:i*4+4], hexTobyteArray[byte(h)])
	}

	parser := Parser{data: []byte(bits), pc: 0}
	pkt := parser.parse()
	fmt.Println(sumVersionNumbers(pkt))
	fmt.Println(pkt.Value())
}

func sumVersionNumbers(p Packet) int {
	switch p := p.(type) {
	case *LiteralPacket:
		return p.Version
	case *OperatorPacket:
		sum := p.Version
		for _, sp := range p.SubPackets {
			sum += sumVersionNumbers(sp)
		}
		return sum
	default:
		return 0
	}
}

func (p *Parser) parse() Packet {
	packetVersion := binaryToInt(p.data[p.pc : p.pc+3])
	packetType := binaryToInt(p.data[p.pc+3 : p.pc+6])

	p.pc += 6
	switch packetType {
	case 4: //literal
		return p.parseLiteral(packetVersion)
	default:
		return p.parseOperator(packetVersion, packetType)
	}
}

func (p *Parser) parseOperator(version, packetType int) *OperatorPacket {
	pkt := OperatorPacket{Version: version, Type: packetType, SubPackets: make([]Packet, 0)}
	lengthType := p.data[p.pc]
	switch lengthType {
	case '0':
		bitLengthOfPackets := binaryToInt(p.data[p.pc+1 : p.pc+16])
		p.pc += 16
		pkt.SubPackets = p.parseBitsForPackets(bitLengthOfPackets)

	case '1':

		numberOfPackets := binaryToInt(p.data[p.pc+1 : p.pc+12])
		p.pc += 12
		pkt.SubPackets = p.parseNumberOfPackets(numberOfPackets)

	}
	return &pkt
}

func (p *Parser) parseNumberOfPackets(number int) []Packet {
	packets := make([]Packet, 0)
	for i := 0; i < number; i++ {
		packets = append(packets, p.parse())
	}

	return packets
}

func (p *Parser) parseBitsForPackets(bitLength int) []Packet {
	packets := make([]Packet, 0)
	start := p.pc
	for p.pc < start+bitLength {
		packets = append(packets, p.parse())
	}

	return packets
}

func (p *Parser) parseLiteral(version int) *LiteralPacket {
	pkt := LiteralPacket{Version: version}
	var keepReading byte
	keepReading = '1'
	literalRaw := make([]byte, 0)
	for keepReading == '1' {
		keepReading = p.data[p.pc]
		data := p.data[p.pc+1 : p.pc+5]
		literalRaw = append(literalRaw, data...)
		p.pc += 5
	}
	pkt.Literal = binaryToInt(literalRaw)
	return &pkt

}

func binaryToInt(bytes []byte) int {
	i, _ := strconv.ParseInt(string(bytes), 2, 0)

	return int(i)
}
