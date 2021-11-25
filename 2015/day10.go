package main

import (
	"fmt"
	"strconv"
)

func lookSay(in, out chan byte) {
	lastDigit := byte('x')
	count := 0

	for curDigit := range in {
		if lastDigit != curDigit && count != 0 {
			countAsString := strconv.Itoa(count)
			for i := range countAsString {
				out <- countAsString[i]
			}
			out <- lastDigit
			count = 0
		}
		count++
		lastDigit = curDigit
	}

	countStr := strconv.Itoa(count)
	for i := range countStr {
		out <- countStr[i]
	}
	out <- lastDigit
	close(out)
}

func main() {
	in := make(chan byte)
	out := make(chan byte)

	input := out

	// Set up 40 (50 for part two) iterations of the lookSay machine,
	// with the output of each einng the input of the next

	for i := 0; i < 50; i++ {
		in = out
		out = make(chan byte)
		go lookSay(in, out)
	}

	output := out
	seed := "3113322113"

	go func() {
		for ch := range seed {
			input <- seed[ch]
		}

		close(input)
	}()

	count := 0
	for _ = range output {
		count++
	}

	fmt.Println(count)

}
