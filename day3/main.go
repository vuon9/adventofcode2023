package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vuon9/adventofcode2023/utils"
)

var charCodeToNum = map[int]int{
	48: 0,
	49: 1,
	50: 2,
	51: 3,
	52: 4,
	53: 5,
	54: 6,
	55: 7,
	56: 8,
	57: 9,
}

func main() {
	input := utils.ReadFileOrPanic()
	// 512794
	fmt.Println("Part 1: ", part1(input))
	// 67779080 - 308 gears
	fmt.Println("Part 2: ", part2(input))
}

func part1(input string) int {
	total := 0

	lines := strings.Split(input, "\n")

	for il, line := range lines {
		if line == "" {
			continue
		}

		currentNumber := ""
		for i := 0; i < len(line); i++ {
			n, found := charCodeToNum[int(line[i])]
			if found {
				currentNumber += fmt.Sprint(n)
			}

			if i < len(line)-1 {
				if _, f := charCodeToNum[int(line[i+1])]; f {
					continue
				}
			}

			if currentNumber == "" {
				continue
			}

			lastCharOfNum := i
			if !found {
				// in case the number is not found, it's the next char, so we have to go back 1 char
				lastCharOfNum = i - 1
			}

			firstCharOfNum := lastCharOfNum - len(currentNumber) + 1

			// has top border
			tb := ""
			if il > 0 {
				// start default as the index of the first char of the number
				start := firstCharOfNum
				end := start + len(currentNumber)
				if firstCharOfNum > 0 {
					start--
				}

				if lastCharOfNum+1 < len(line)-1 {
					end++
				}

				tb = lines[il-1][start:end]
			}

			bb := ""
			if il < len(lines)-1 {
				start := firstCharOfNum
				end := start + len(currentNumber)
				if firstCharOfNum > 0 {
					start--
				}

				if lastCharOfNum+1 < len(line)-1 {
					end++
				}

				bb = lines[il+1][start:end]
			}

			lb := ""
			if firstCharOfNum-1 > 0 {
				lb = string(line[firstCharOfNum-1])
			}

			rb := ""
			if lastCharOfNum+1 < len(line)-1 {
				rb = string(line[lastCharOfNum+1])
			}

			for _, l := range tb + lb + rb + bb {
				if string(l) != "." {
					num, _ := strconv.Atoi(currentNumber)
					total += num
				}
			}

			// fmt.Println("* number: ", currentNumber, "\n fc", string(line[firstCharOfNum]), ", lc", string(line[lastCharOfNum]), "\n -- tb: ", tb, "\n -- lb: ", lb, "\n -- rb: ", rb, "\n -- bb: ", bb)

			currentNumber = ""

		}
	}

	return total
}

type pt struct {
	line int
	col  int
}

func part2(input string) int {
	total := 0

	indices := make(map[pt][]int)
	lines := strings.Split(input, "\n")

	for il, line := range lines {
		if line == "" {
			continue
		}

		currentNumber := ""
		for i := 0; i < len(line); i++ {
			n, found := charCodeToNum[int(line[i])]
			if found {
				currentNumber += fmt.Sprint(n)
			}

			if i < len(line)-1 {
				if _, f := charCodeToNum[int(line[i+1])]; f {
					continue
				}
			}

			if currentNumber == "" {
				continue
			}

			lastCharOfNum := i
			if !found {
				// in case the number is not found, it's the next char, so we have to go back 1 char
				lastCharOfNum = i - 1
			}

			firstCharOfNum := lastCharOfNum - len(currentNumber) + 1

			// has top border
			tb := ""
			if il > 0 {
				// start default as the index of the first char of the number
				start := firstCharOfNum
				end := start + len(currentNumber)
				if firstCharOfNum > 0 {
					start--
				}

				if lastCharOfNum+1 < len(line)-1 {
					end++
				}

				tb = lines[il-1][start:end]
			}

			bb := ""
			if il < len(lines)-1 {
				start := firstCharOfNum
				end := start + len(currentNumber)
				if firstCharOfNum > 0 {
					start--
				}

				if lastCharOfNum+1 < len(line)-1 {
					end++
				}

				bb = lines[il+1][start:end]
			}

			lb := ""
			if firstCharOfNum-1 > 0 {
				lb = string(line[firstCharOfNum-1])
			}

			rb := ""
			if lastCharOfNum+1 < len(line)-1 {
				rb = string(line[lastCharOfNum+1])
			}

			num, _ := strconv.Atoi(currentNumber)

			if string(lb) != "" && string(lb) != "." {
				p := pt{il, firstCharOfNum - 1}
				indices[p] = append(indices[p], num)
			}

			if string(rb) != "" && string(rb) != "." {
				p := pt{il, lastCharOfNum + 1}
				indices[p] = append(indices[p], num)
			}

			for index, l := range tb {
				if string(l) != "." {
					col := index
					if firstCharOfNum > 0 {
						col = firstCharOfNum - 1 + index
					}

					p := pt{il - 1, col}
					indices[p] = append(indices[p], num)
					break
				}
			}

			for index, l := range bb {
				if string(l) != "." {
					col := index
					if firstCharOfNum > 0 {
						col = firstCharOfNum - 1 + index
					}

					p := pt{il + 1, col}
					indices[p] = append(indices[p], num)
					break
				}
			}

			// fmt.Println("* number: ", currentNumber, "\n fc", string(line[firstCharOfNum]), ", lc", string(line[lastCharOfNum]), "\n -- tb: ", tb, "\n -- lb: ", lb, "\n -- rb: ", rb, "\n -- bb: ", bb)

			currentNumber = ""

		}
	}

	for _, gears := range indices {
		if len(gears) > 1 {
			gear := 0
			for gi, n := range gears {
				if gi == 0 {
					gear = n
				} else {
					gear *= n
				}
			}

			total += gear
			// fmt.Println("p", p, "gear", gears)
		} else {
			// fmt.Println("p", p, "gear", gears)
		}
	}

	return total
}
