package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	PART_TWO = false
)

func segmentToDigit(r string) int {
	switch r[0] {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	}
	if PART_TWO {
		if len(r) >= 4 && r[0:4] == "zero" {
			return 0
		} else if len(r) >= 3 && r[0:3] == "one" {
			return 1
		} else if len(r) >= 3 && r[0:3] == "two" {
			return 2
		} else if len(r) >= 5 && r[0:5] == "three" {
			return 3
		} else if len(r) >= 4 && r[0:4] == "four" {
			return 4
		} else if len(r) >= 4 && r[0:4] == "five" {
			return 5
		} else if len(r) >= 3 && r[0:3] == "six" {
			return 6
		} else if len(r) >= 5 && r[0:5] == "seven" {
			return 7
		} else if len(r) >= 5 && r[0:5] == "eight" {
			return 8
		} else if len(r) >= 4 && r[0:4] == "nine" {
			return 9
		}
	}
	return -1
}

func main() {
	f, err := os.ReadFile("./data.txt")
	if err != nil {
		panic(err)
	}
	s := string(f)
	lines := strings.Split(s, "\n")
	sum := 0
	for _, line := range lines {
		first := -1
		last := -1
		for i, _ := range line {
			d := segmentToDigit(line[i:])
			if d != -1 {
				if first == -1 {
					first = d
				}
				last = d
			}
		}
		lineNum, _ := strconv.Atoi(fmt.Sprintf("%v%v", first, last))
		sum += lineNum
	}
	fmt.Printf("%v", sum)
}
