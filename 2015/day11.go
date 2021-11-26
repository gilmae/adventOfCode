package main

import (
	"fmt"
)

func isValid(poss []byte) bool {
	// Must include straight
	// Must not include i or o or l
	// Must include two doubles
	hasStraight := false
	for i := 0; i < len(poss)-2; i++ {
		if poss[i]+1 == poss[i+1] && poss[i+1]+1 == poss[i+2] {
			hasStraight = true
		}
	}

	if !hasStraight {
		return false
	}

	for i := range poss {
		if poss[i] == 'i' || poss[i] == 'l' || poss[i] == 'o' {
			return false
		}
	}

	pairs := make(map[byte]bool)
	for i := 0; i < len(poss)-1; i++ {
		if poss[i] == poss[i+1] {
			pairs[poss[i]] = true
		}
	}

	return len(pairs) >= 2
}

func nextPassword(passwd []byte) {
	carry := true
	for position := len(passwd) - 1; position >= 0; position-- {
		if carry {
			if passwd[position] == 'z' {
				passwd[position] = 'a'
			} else {
				carry = false
				passwd[position] = passwd[position] + 1
			}
		}
	}
}

func main() {
	input := []byte("hepxcrrq")
	nextPassword(input)
	for !isValid(input) {
		nextPassword(input)
	}
	fmt.Println(string(input))
	nextPassword(input)
	for !isValid(input) {
		nextPassword(input)
	}
	fmt.Println(string(input))
}
