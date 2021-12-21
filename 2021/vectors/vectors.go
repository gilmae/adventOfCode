package vector

type Coords3d struct {
	X, Y, Z int
}

func (c Coords3d) Difference(dest Coords3d) Coords3d {
	return Coords3d{c.X - dest.X, c.Y - dest.Y, c.Z - dest.Z}
}

func (c Coords3d) Add(dest Coords3d) Coords3d {
	return Coords3d{c.X + dest.X, c.Y + dest.Y, c.Z + dest.Z}
}

func (c *Coords3d) Transform(matrix [3][3]int) Coords3d {
	result := Coords3d{}

	result.X = c.X*matrix[0][0] + c.Y*matrix[0][1] + c.Z*matrix[0][2]
	result.Y = c.X*matrix[1][0] + c.Y*matrix[1][1] + c.Z*matrix[1][2]
	result.Z = c.X*matrix[2][0] + c.Y*matrix[2][1] + c.Z*matrix[2][2]

	return result
}
