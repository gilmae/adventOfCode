package boards

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
