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

	lines := strings.Split(string(input), "\n\n")
	bingoNumbers := strings.Split(lines[0], ",")

	winningNumber := ""
	winningCount := 0
	winningSum := 0

	for i := 1; i < len(lines); i++ {
		bingoField := strings.Split(string(lines[i]), "\n")
		bingoLine := make([]string, 0)

		for x := 0; x < len(bingoField); x++ {
			line := strings.Split(string(bingoField[x]), " ")
			line = delete_empty(line)
			bingoLine = append(bingoLine, line...)
		}

		bingoMap := make(map[string]int)
		winArray := [25]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for a := 0; a < len(bingoNumbers); a++ {
			bingoNumber := bingoNumbers[a]
			if stringInSlice(bingoNumber, bingoLine) {
				bingoMap[bingoNumber] = 1
				for i, b := range bingoLine {
					if b == bingoNumber {
						winArray[i] = 1
					}
				}
			} else {
				bingoMap[bingoNumber] = 0
			}
			if checkWin(winArray) {
				if a > winningCount {
					lineSum := 0
					winningNumber = bingoNumbers[a]
					winningCount = a
					for index, mark := range winArray {
						if mark == 0 {
							value, err := strconv.Atoi(bingoLine[index])
							check(err)
							lineSum += value
						}
					}
					winningSum = lineSum
				}
				break
			}
		}
	}
	value, err := strconv.Atoi(winningNumber)
	check(err)
	fmt.Println(winningCount, winningSum*value)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func checkWin(bingo [25]int) bool {
	for i := 0; i < len(bingo); i += 5 {
		if bingo[i] == 1 &&
			bingo[i+1] == 1 &&
			bingo[i+2] == 1 &&
			bingo[i+3] == 1 &&
			bingo[i+4] == 1 {
			return true
		}
	}

	for a := 0; a < 5; a++ {
		if bingo[a] == 1 &&
			bingo[a+5] == 1 &&
			bingo[a+10] == 1 &&
			bingo[a+15] == 1 &&
			bingo[a+20] == 1 {
			return true
		}
	}
	return false
}
