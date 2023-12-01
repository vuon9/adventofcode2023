package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/vuon9/adventofcode2023/utils"
)

func main() {
	input := utils.ReadFileOrPanic()
	fmt.Println("Part 1, result is: ", part1(string(input)))
	fmt.Println("Part 2, result is: ", part2(string(input)))
}

func part2(input string) string {
	lines := strings.Split(input, "\n")
	strToNum := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	total := 0
	for _, line := range lines {
		fn, ln := "", ""
		for i, j := 0, len(line)-1; i < len(line); i, j = i+1, j-1 {
			if fn == "" {
				foundFn := false
				for str, num := range strToNum {
					if len(line) >= i+len(str) {
						qw := line[i : i+len(str)]
						if str == qw {
							fn = num
							i += len(str)
							foundFn = true
							break
						}
					}
				}

				if !foundFn {
					ni, err := strconv.Atoi(string(line[i]))
					if err == nil {
						fn = fmt.Sprintf("%d", ni)
					}
				}
			}

			if ln == "" {
				foundLn := false

				for str, num := range strToNum {
					if j-len(str) >= 0 {
						qw := line[j-len(str)+1 : j+1]
						if str == qw {
							ln = num
							foundLn = true
							j -= len(str)
							break
						}
					}
				}

				if !foundLn {
					nj, err := strconv.Atoi(string(line[j]))
					if err == nil {
						ln = fmt.Sprintf("%d", nj)
					}
				}
			}
		}

		fnn, _ := strconv.Atoi(fmt.Sprintf("%s%s", fn, ln))
		total += fnn
	}

	return fmt.Sprintf("%d", total)
}

func part1(input string) string {
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		fn, ln := "", ""

		for i, j := 0, len(line)-1; i < len(line); i, j = i+1, j-1 {
			if fn == "" {
				ni, err := strconv.Atoi(string(line[i]))
				if err == nil {
					fn = fmt.Sprintf("%d", ni)
				}
			}

			if ln == "" {
				nj, err := strconv.Atoi(string(line[j]))
				if err == nil {
					ln = fmt.Sprintf("%d", nj)
				}
			}
		}

		fnn, _ := strconv.Atoi(fmt.Sprintf("%s%s", fn, ln))
		// fmt.Println("line: ", lineNo, ", number is: ", fnn)
		total += fnn
	}

	return fmt.Sprintf("%d", total)
}
