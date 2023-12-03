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
	//
	fmt.Println("Part 1: ", part1(input))
	//
	fmt.Println("Part 2: ", part2(input))
}

func part1(input string) int {
	total := 0

	lines := strings.Split(input, "\n")

	for il, line := range lines {
		if line == "" {
			continue
		}

		lineCharIndex := 0
		currentNumber := ""
		for i := 0; i < len(line); i++ {
			n, found := charCodeToNum[int(line[i])]
			if found {
				currentNumber += fmt.Sprint(n)
			}

			if (!found && currentNumber != "") || (i == len(line)-1 && found) {
				hasTopBorder := il > 0
				hasLeftBorder := lineCharIndex-len(currentNumber) > 0
				hasRightBorder := lineCharIndex <= len(line)-1
				if i == len(line)-1 {
					hasRightBorder = false
				}

				hasBottomBorder := il < len(lines)-1
				if il == len(lines)-1 {
					hasBottomBorder = false
				}

				if n > 0 {
					lineCharIndex++
				}

				var topBorder, rightBorder, leftBorder, bottomBorder string
				var isPartNumber bool

				fmt.Printf("num: %s, tb: %t, rb: %t, bb: %t, lb: %t", currentNumber, hasTopBorder, hasRightBorder, hasBottomBorder, hasLeftBorder)

				// top -> right -> bottom -> left
				if hasTopBorder {
					// direction from left to right
					borderStartFrom := lineCharIndex - len(currentNumber)
					borderLength := len(currentNumber)
					if hasLeftBorder {
						borderStartFrom--
						borderLength++
					}

					topBorder = lines[il-1][borderStartFrom : borderStartFrom+borderLength]
					fmt.Printf("\n-- tb: %s", topBorder)
				}

				if hasRightBorder {
					// direction is from top to bottom
					rightBorder = string(line[lineCharIndex])
					if hasTopBorder {
						rightBorder = string(lines[il-1][lineCharIndex]) + rightBorder
					}

					fmt.Printf("\n-- rb: %s", string(rightBorder))
				}

				if hasBottomBorder {
					// direction is from right to left
					borderStartFrom := lineCharIndex
					borderLength := len(currentNumber)
					if hasRightBorder && borderStartFrom < len(line)-1 {
						borderStartFrom++
						borderLength++
					}

					rvBottomBorder := lines[il+1][borderStartFrom-borderLength : borderStartFrom]
					bottomBorder = ""
					// just for reading purpose
					for _, c := range rvBottomBorder {
						bottomBorder = string(c) + bottomBorder
					}

					fmt.Printf("\n-- bb: %s", bottomBorder)
				}

				if hasLeftBorder {
					// direction from bottom to top
					leftBorder = string(line[lineCharIndex-len(currentNumber)-1])
					if hasBottomBorder {
						leftBorder = string(lines[il+1][lineCharIndex-len(currentNumber)-1]) + leftBorder
					}

					fmt.Printf("\n-- lb: %s", string(leftBorder))
				}

				border := topBorder + rightBorder + bottomBorder + leftBorder
				for _, c := range border {
					if string(c) != "." {
						isPartNumber = true
						break
					}
				}

				if isPartNumber {
					fmt.Printf("\n-- is part number, border: %s", border)
					num, _ := strconv.Atoi(currentNumber)
					total += num
				} else {
					fmt.Printf("\n-- is not part number, border: %s", border)
				}

				fmt.Print("\n")

				// reset current number to empty
				currentNumber = ""
			}

			lineCharIndex++
		}
	}

	// if the first line, we dont have top border
	// if the number is top left, we dont have left border
	// if the number is top right, we dont have right border

	// if the last line, we dont have bottom border
	// if the number is bottom right, we dont have right border
	// if the number is bottom left, we dont have left border

	// get top border
	// get left border
	// get right border
	// get bottom border

	return total
}

func part2(input string) int {
	return 0
}
