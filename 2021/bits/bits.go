package bits

import (
	"sort"
	"strconv"
)

type Packet interface {
	PacketType() string
	Value() int
}

type LiteralPacket struct {
	Version, Literal int
}

const (
	SUM      = "SUM"
	MULTIPLY = "MULTIPLY"
	MIN      = "MIN"
	MAX      = "MAX"
	LT       = "LT"
	GT       = "GT"
	EQ       = "EQ"
)

var PacketTypeToOperator = map[int]string{
	0: SUM,
	1: MULTIPLY,
	2: MIN,
	3: MAX,
	5: GT,
	6: LT,
	7: EQ,
}

func (l *LiteralPacket) PacketType() string { return "LITERAL" }

func (l *LiteralPacket) Value() int { return l.Literal }

type OperatorPacket struct {
	Version    int
	Operator   string
	SubPackets []Packet
}

func (o *OperatorPacket) PacketType() string { return "OPERATOR" }
func (o *OperatorPacket) Value() int {

	switch o.Operator {
	case SUM:
		value := 0
		for _, sb := range o.SubPackets {
			value += sb.Value()
		}
		return value
	case MULTIPLY:
		value := 1
		for _, sb := range o.SubPackets {
			value *= sb.Value()
		}
		return value
	case MIN:
		values := make([]int, len(o.SubPackets))
		for i, sb := range o.SubPackets {
			values[i] = sb.Value()
		}

		sort.Sort(sort.IntSlice(values))
		return values[0]
	case MAX:
		values := make([]int, len(o.SubPackets))
		for i, sb := range o.SubPackets {
			values[i] = sb.Value()
		}

		sort.Sort(sort.Reverse(sort.IntSlice(values)))
		return values[0]
	case GT:
		if o.SubPackets[0].Value() > o.SubPackets[1].Value() {
			return 1
		} else {
			return 0
		}
	case LT:
		if o.SubPackets[0].Value() < o.SubPackets[1].Value() {
			return 1
		} else {
			return 0
		}
	case EQ:
		if o.SubPackets[0].Value() == o.SubPackets[1].Value() {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
}

type Parser struct {
	data []byte
	pc   int
}

func NewParser(data []byte) *Parser {
	p := Parser{data: data, pc: 0}
	return &p
}

func (p *Parser) Parse() Packet {
	packetVersion := binaryToInt(p.data[p.pc : p.pc+3])
	packetType := binaryToInt(p.data[p.pc+3 : p.pc+6])

	p.pc += 6
	switch packetType {
	case 4: //literal
		return p.parseLiteral(packetVersion)
	default:
		return p.parseOperator(packetVersion, PacketTypeToOperator[packetType])
	}
}

func (p *Parser) parseOperator(version int, operator string) *OperatorPacket {
	pkt := OperatorPacket{Version: version, Operator: operator, SubPackets: make([]Packet, 0)}
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
		packets = append(packets, p.Parse())
	}

	return packets
}

func (p *Parser) parseBitsForPackets(bitLength int) []Packet {
	packets := make([]Packet, 0)
	start := p.pc
	for p.pc < start+bitLength {
		packets = append(packets, p.Parse())
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
