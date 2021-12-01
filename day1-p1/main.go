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

	for i := 0; i < len(lines)-1; i++ {
		currentNum, err := strconv.Atoi(lines[i])
		check(err)

		nextNum, err := strconv.Atoi(lines[i+1])
		check(err)

		if currentNum < nextNum {
			increased++
		}
	}
	fmt.Println(increased)
}
