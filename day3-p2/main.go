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

	oxygen := lines
	co2 := lines

	lenOfBinary := len(lines[0])
	for a := 0; a < lenOfBinary; a++ {
		oxygenZeros, oxygenOnes := countAppearence(oxygen, a)
		co2Zeros, co2Ones := countAppearence(co2, a)

		if len(oxygenOnes) > 0 && len(oxygenZeros) > 0 {
			if len(oxygenOnes) > len(oxygenZeros) || len(oxygenZeros) == len(oxygenOnes) {
				oxygen = oxygenOnes
			} else if len(oxygenZeros) > len(oxygenOnes) {
				oxygen = oxygenZeros
			} else {
				panic("Something happend")
			}
		}
		if len(co2Ones) > 0 && len(co2Zeros) > 0 {
			if len(co2Ones) > len(co2Zeros) || len(co2Zeros) == len(co2Ones) {
				co2 = co2Zeros
			} else if len(co2Zeros) > len(co2Ones) {
				co2 = co2Ones
			} else {
				panic("Something happend")
			}
		}
	}
	oxygenInt, err := strconv.ParseInt(oxygen[0], 2, 64)
	check(err)
	co2Int, err := strconv.ParseInt(co2[0], 2, 64)
	check(err)
	fmt.Println(lenOfBinary, oxygenInt, co2Int, oxygenInt*co2Int)
}

func countAppearence(lines []string, a int) (zeros []string, ones []string) {
	for i := 0; i < len(lines); i++ {
		binary := lines[i]
		binChar := binary[a : a+1]

		if binChar == "1" {
			ones = append(ones, binary)
		} else if binChar == "0" {
			zeros = append(zeros, binary)
		}
	}
	return zeros, ones
}
