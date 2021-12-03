package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input, err := os.ReadFile("./input.txt")
	check(err)

	lines := strings.Split(string(input), "\n")

	gamma := ""
	epsilon := ""

	lenOfBinary := len(lines[0])
	for a := 0; a < lenOfBinary; a++ {
		zeros := 0
		ones := 0
		for i := 0; i < len(lines); i++ {
			binary := lines[i]
			binChar := binary[a : a+1]

			if binChar == "1" {
				ones += 1
			} else if binChar == "0" {
				zeros += 1
			}
		}
		if ones > zeros {
			gamma += "1"
			epsilon += "0"
		} else if zeros > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			panic("even")
		}
	}
	gammaInt, err := strconv.ParseInt(gamma, 2, 64)
	check(err)
	epsilonInt, err := strconv.ParseInt(epsilon, 2, 64)
	check(err)
	fmt.Println(lenOfBinary, gammaInt, epsilonInt, epsilonInt*gammaInt)
}
