package boards

import "fmt"

type Coords struct {
	X, Y int
}

type Board struct {
	Points map[Coords]interface{}
}

func NewBoard() *Board {
	b := Board{make(map[Coords]interface{})}

	return &b
}

type transformer func(c Coords, value interface{}) interface{}

func (b *Board) Import(lines []string, transform transformer) {
	if transform == nil {
		transform = func(c Coords, value interface{}) interface{} {
			return value
		}
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

func (b *Board) PrintBoard() {
	maxX, maxY := b.Width(), b.Height()
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			c := Coords{x, y}
			fmt.Printf("%+v", b.Points[c])

		}
		fmt.Println()
	}
}
