package graphs

import (
	"fmt"
	"math"
)

type Shader func(p Coords, curValue interface{}) interface{}
type Coords struct {
	X, Y int
}

func (c *Coords) ManhattanDistance(from Coords) int {
	return int(math.Abs(float64(from.X-c.X)) + math.Abs(float64(from.Y-c.Y)))
}

type Line struct {
	start, end Coords
}

func NewLine(start, end Coords) *Line {
	l := Line{start, end}
	return &l
}

func (l *Line) GetPoints() []Coords {

	length := 1

	delta := getDelta(l.start, l.end)

	if l.start.X != l.end.X {
		length += int(math.Abs(float64(l.end.X) - float64(l.start.X)))
	} else {
		length += int(math.Abs(float64(l.end.Y) - float64(l.start.Y)))
	}

	out := make([]Coords, length)
	for d := 0; d < length; d++ {
		c := Coords{l.start.X + delta.X*d, l.start.Y + delta.Y*d}
		out[d] = c
	}

	return out
}

type Board struct {
	Points map[Coords]interface{}
}

func NewBoard() *Board {
	b := Board{make(map[Coords]interface{})}

	return &b
}

func (b *Board) Draw() {
	max := Coords{int(^uint(0)>>1) * -1, int(^uint(0)>>1) * -1}
	min := Coords{int(^uint(0) >> 1), int(^uint(0) >> 1)}

	for k, _ := range b.Points {
		if k.X > max.X {
			max.X = k.X
		} else if k.X < min.X {
			min.X = k.X
		}

		if k.Y > max.Y {
			max.Y = k.Y
		} else if k.Y < min.Y {
			min.Y = k.Y
		}
	}

	for x := max.X; x >= min.X; x-- {
		for y := min.Y; y <= max.Y; y++ {
			c := Coords{x, y}
			if v, ok := b.Points[c]; !ok {
				fmt.Print(" ")
			} else if v == -1 {
				fmt.Print("X")
			} else {
				fmt.Print("*")
			}

		}
		fmt.Println()
	}
}

func (b *Board) DrawPoint(c Coords, shader Shader) {
	b.Points[c] = shader(c, b.Points[c])
}

func (b *Board) DrawPoints(points []Coords, shader Shader) {
	for _, c := range points {
		b.Points[c] = shader(c, b.Points[c])
	}

}

func getDelta(start, end Coords) Coords {
	var deltaX int

	if start.X > end.X {
		deltaX = -1
	} else if start.X == end.X {
		deltaX = 0
	} else {
		deltaX = 1
	}

	var deltaY int
	if start.Y > end.Y {
		deltaY = -1
	} else if start.Y == end.Y {
		deltaY = 0
	} else {
		deltaY = 1
	}

	return Coords{deltaX, deltaY}
}
