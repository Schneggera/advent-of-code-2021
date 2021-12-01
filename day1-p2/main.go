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

	increased := 0

	for i := 0; i < len(lines)-3; i++ {
		currentNum, err := strconv.Atoi(lines[i])
		check(err)

		nextNum, err := strconv.Atoi(lines[i+1])
		check(err)

		thirdNum, err := strconv.Atoi(lines[i+2])
		check(err)

		fourthNum, err := strconv.Atoi(lines[i+3])
		check(err)

		threeMeasurement1 := currentNum + nextNum + thirdNum
		threeMeasurement2 := nextNum + thirdNum + fourthNum

		if threeMeasurement1 < threeMeasurement2 {
			increased++
		}
	}
	fmt.Println(increased)
}
