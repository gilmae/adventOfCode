package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
)

var part = flag.String("part", "a", "Which part to solve")

func main() {
	flag.Parse()
	codePoint := 0
	code := []byte("        ")
	for index := 0; codePoint != 8; index++ {
		test := []byte(fmt.Sprintf("%s%d", "reyedfim", index))
		h := md5.Sum(test)
		hash := hex.EncodeToString(h[:])

		if hash[0] == '0' && hash[1] == '0' && hash[2] == '0' && hash[3] == '0' && hash[4] == '0' {
			if *part == "a" {
				code[codePoint] = hash[5]
			} else {
				pos := hash[5] - '0'
				if pos >= 0 && pos <= 7 && code[pos] == ' ' {

					code[pos] = hash[6]
				} else {
					continue
				}
			}
		} else {
			continue
		}
		fmt.Println(string(code))
		codePoint++
	}

}
