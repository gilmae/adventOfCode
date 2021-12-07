package main

import (
	"fmt"
	"strconv"
)

func main() {
	passwd := 382344
	validCount := 0
	for passwd <= 843167 {
		passwd++
		if isValid(passwd) {
			validCount++
		}
	}

	fmt.Println(validCount)

}

func isValid(passwd int) bool {
	passwdS := strconv.Itoa(passwd + 1)

	for i := 0; i < len(passwdS)-1; i++ {
		if passwdS[i+1] < passwdS[i] {
			return false
		}
	}
	for i := 0; i < len(passwdS)-1; i++ {
		if passwdS[i+1] == passwdS[i] {
			return true
		}
	}
	return false
}
