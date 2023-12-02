package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/vuon9/adventofcode2023/utils"
)

var strToNum = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var charToNum = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func main() {
	input := utils.ReadFileOrPanic()
	// 54634
	fmt.Println("Part 1, result is: ", part1(string(input)))
	// 53855
	fmt.Println("Part 2, result is: ", part2(string(input)))
}

func part2(input string) string {
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		fn, ln := 0, 0
		for i, j := 0, len(line)-1; i < len(line); i, j = i+1, j-1 {
			if fn == 0 {
				if charToNum[string(line[i])] != 0 {
					fn, _ = charToNum[string(line[i])]
				} else {
					for str, num := range strToNum {
						if strings.HasPrefix(line[i:], str) {
							fn = num
							i += len(str)
							break
						}
					}
				}
			}

			if ln == 0 {
				if charToNum[string(line[j])] != 0 {
					ln = charToNum[string(line[j])]
				} else {
					for str, num := range strToNum {
						if j-len(str) >= 0 && str == line[j-len(str)+1:j+1] {
							ln = num
							j -= len(str)
							break
						}
					}
				}
			}
		}

		fnn, _ := strconv.Atoi(fmt.Sprintf("%d%d", fn, ln))
		total += fnn
	}

	return fmt.Sprintf("%d", total)
}

func part1(input string) string {
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		fn, ln := 0, 0

		for i, j := 0, len(line)-1; i < len(line); i, j = i+1, j-1 {
			if fn == 0 {
				fn = charToNum[string(line[i])]
			}

			if ln == 0 {
				ln = charToNum[string(line[j])]
			}
		}

		fnn, _ := strconv.Atoi(fmt.Sprintf("%d%d", fn, ln))
		total += fnn
	}

	return fmt.Sprintf("%d", total)
}
