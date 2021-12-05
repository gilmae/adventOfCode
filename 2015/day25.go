package main

import "fmt"

// To determine the position of Triangle numbers,
// start at the col and row. Until col == 1, decrement col and increment row
// when col == 1, set row to 1 and col to row - 1
// when col == 1 and row == 1, there you are
func determineCellNumber(row, col int) int {
	position := 0
	for {
		if col == 1 && row == 1 {
			return position
		}
		position++
		if col > 1 {
			col--
			row++
		} else {
			col = row - 1
			row = 1
		}
	}
}

func main() {
	position := determineCellNumber(3010, 3019)
	fmt.Println(position)
	value := 20151125
	for i := 0; i < position; i++ {
		value *= 252533
		value %= 33554393

	}

	fmt.Println(value)
}
