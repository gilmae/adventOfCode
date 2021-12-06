package graphs

import "testing"

func TestManhattanDistance(t *testing.T) {
	from := Coords{0, 0}
	to := Coords{1, 2}

	if to.ManhattanDistance(from) != 3 {
		t.Errorf("Expected 3, got %d", to.ManhattanDistance(from))
	}
}

func TestDrawLine(t *testing.T) {
	from := Coords{0, 0}
	to := Coords{0, 8}
	line := Line{from, to}

	points := line.GetPoints()
	pointsMap := make(map[Coords]bool)
	for _, p := range points {
		pointsMap[p] = true
	}

	if len(pointsMap) != 9 {
		t.Errorf("Expected 9 points, got %d", len(pointsMap))
	}
	for i := 0; i <= 8; i++ {
		if _, ok := pointsMap[Coords{0, i}]; !ok {
			t.Errorf("Expected %d,%d to be returned", 0, i)
		}

	}

}
