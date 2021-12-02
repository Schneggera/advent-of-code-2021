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
	aim := 0

	for i := 0; i < len(lines); i++ {
		line := strings.Split(string(lines[i]), " ")
		value, err := strconv.Atoi(line[1])
		check(err)

		switch line[0] {
		case "forward":
			horizontal += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}
	fmt.Println(horizontal, depth, horizontal*depth)
}
