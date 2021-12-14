package boards

import (
	"fmt"
	"math"
)

type Coords struct {
	X, Y int
}

type Board struct {
	Points      map[Coords]interface{}
	TopLeft     Coords
	BottomRight Coords
}

func NewBoard() *Board {
	b := Board{make(map[Coords]interface{}), Coords{0, 0}, Coords{0, 0}}

	return &b
}

type transformer func(c Coords, value interface{}) interface{}

func (b *Board) Import(lines []string, transform transformer) {
	if transform == nil {
		transform = defaultTransform
	}
	for y, line := range lines {
		for x, ch := range line {
			c := Coords{x, y}
			b.Points[c] = transform(c, ch)
		}
	}
}

func (b *Board) Width() int {
	maxX := -1*int(^uint(0)>>1) - 1

	for k, _ := range b.Points {

		if k.X > maxX {
			maxX = k.X
		}
	}
	return maxX
}

func (b *Board) Height() int {
	maxY := -1*int(^uint(0)>>1) - 1

	for k, _ := range b.Points {
		if k.Y > maxY {
			maxY = k.Y
		}
	}
	return maxY
}

func (b *Board) GetNeighbours(c Coords, ignoreDiagonals bool) []Coords {
	neighbours := make([]Coords, 0)
	for j := -1; j < 2; j++ {
		for i := -1; i < 2; i++ {
			if i == 0 && j == 0 {
				continue
			}

			if ignoreDiagonals && i != 0 && j != 0 {
				continue
			}

			p := Coords{c.X + i, c.Y + j}
			if _, ok := b.Points[p]; ok {
				neighbours = append(neighbours, p)
			}
		}
	}

	return neighbours
}

func (board *Board) FoldY(foldPoint int) *Board {
	newBoard := NewBoard()

	for k, v := range board.Points {
		if k.Y >= foldPoint {
			newBoard.Points[Coords{k.X, foldPoint - (k.Y - foldPoint)}] = v
		} else {
			newBoard.Points[k] = v
		}
	}
	return newBoard
}

func (board *Board) FoldX(foldPoint int) *Board {
	newBoard := NewBoard()

	for k, v := range board.Points {
		if k.X >= foldPoint {
			newBoard.Points[Coords{foldPoint - (k.X - foldPoint), k.Y}] = v
		} else {
			newBoard.Points[k] = v
		}
	}
	return newBoard
}

// FlipX flips the board on the x axis, so max y becomes min y
func (b *Board) FlipX() *Board {
	nb := NewBoard()

	for y := 0; y <= b.BottomRight.Y; y++ {
		newY := b.BottomRight.Y - y
		for x := 0; x <= b.BottomRight.X; x++ {
			if v, ok := b.Points[Coords{x, y}]; ok {
				nb.Points[Coords{x, newY}] = v
			}
		}
	}
	nb.TopLeft = b.TopLeft
	nb.BottomRight = b.BottomRight
	return nb
}

// FlipX flips the board on the y axis, so max x becomes min x
func (b *Board) FlipY() *Board {
	nb := NewBoard()

	for x := 0; x <= b.BottomRight.X; x++ {
		newX := b.BottomRight.X - x
		for y := 0; y <= b.BottomRight.Y; y++ {
			if v, ok := b.Points[Coords{x, y}]; ok {
				nb.Points[Coords{newX, y}] = v
			}
		}
	}
	nb.TopLeft = b.TopLeft
	nb.BottomRight = b.BottomRight
	return nb
}

func (b *Board) RotateClockwise() *Board {
	nb := NewBoard()

	centre := Coords{(b.BottomRight.X - b.TopLeft.X) / 2, (b.BottomRight.Y - b.TopLeft.Y) / 2}
	for x := 0; x <= b.BottomRight.X; x++ {
		for y := 0; y <= b.BottomRight.Y; y++ {
			p := Coords{x, y}
			if v, ok := b.Points[p]; ok {
				nb.Points[rotatePointByDegrees(Coords{x, y}, centre, -90)] = v
			}
		}
	}
	nb.TopLeft = b.TopLeft
	nb.BottomRight = b.BottomRight
	return nb
}

func (b *Board) RotateCounterClockwise() *Board {
	nb := NewBoard()

	centre := Coords{(b.BottomRight.X - b.TopLeft.X) / 2, (b.BottomRight.Y - b.TopLeft.Y) / 2}
	for x := 0; x <= b.BottomRight.X; x++ {
		for y := 0; y <= b.BottomRight.Y; y++ {
			p := Coords{x, y}
			if v, ok := b.Points[p]; ok {
				nb.Points[rotatePointByDegrees(Coords{x, y}, centre, 90)] = v
			}
		}
	}
	nb.TopLeft = b.TopLeft
	nb.BottomRight = b.BottomRight
	return nb
}

func (b *Board) PrintBoard() {
	b.PrintBoardWithShader(defaultShader)
}

func (b *Board) PrintBoardWithShader(transform transformer) {
	if transform == nil {
		transform = func(c Coords, value interface{}) interface{} {
			return fmt.Sprintf("%+v", value)
		}
	}
	maxX, maxY := b.Width(), b.Height()
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			c := Coords{x, y}
			fmt.Printf("%+v", transform(c, b.Points[c]))

		}
		fmt.Println()
	}
}

func defaultTransform(c Coords, value interface{}) interface{} {
	return value
}

func defaultShader(c Coords, v interface{}) interface{} {
	switch v := v.(type) {
	case nil:
		return " "
	case bool:
		if !v {
			return "."
		} else {
			return "#"
		}
	default:
		return fmt.Sprintf("%+v", v)
	}
}

func rotatePointByDegrees(p Coords, centre Coords, degrees int) Coords {
	// Things are a little weird, yo. Because the Top LEft is 0,0 and the bottom right is width, height
	// We actually have to flip the degrees and make clockwise *counter* clockwise
	radians := float64(degrees*-1) * (math.Pi / 180)
	cos := math.Cos(radians)
	sin := math.Sin(radians)

	dx := int(float64(centre.X) + float64(p.X-centre.X)*cos - float64(p.Y-centre.Y)*sin)
	dy := int(float64(centre.Y) + float64(p.X-centre.X)*sin + float64(p.Y-centre.Y)*cos)

	return Coords{dx, dy}

}
