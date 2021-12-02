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

	horizontal := 0
	depth := 0

	for i := 0; i < len(lines); i++ {
		line := strings.Split(string(lines[i]), " ")
		value, err := strconv.Atoi(line[1])
		check(err)

		switch line[0] {
		case "forward":
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}
	fmt.Println(horizontal, depth, horizontal*depth)
}
