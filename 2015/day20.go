package main

import "fmt"

func main() {
	goal := 33100000 / 10
	house := 1
	for sumDivisors(house) < goal {
		house++
	}
	fmt.Println(house)
}

func sumDivisors(house int) int {

	sum := 0
	for i := 1; i*i < house; i++ {
		if house%i == 0 {
			sum += i + house/i
		}
	}
	return sum
}
