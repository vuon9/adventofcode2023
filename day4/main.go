package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vuon9/adventofcode2023/utils"
)

func main() {
	input := utils.ReadFileOrPanic()
	// 23028
	fmt.Println("Part 1: ", part1(input))
	//
	// fmt.Println("Part 2: ", part2(input))
}

func part1(input string) int {
	total := 0

	for _, rawCard := range strings.Split(input, "\n") {
		cardNumber := 0
		winNumbers := make(map[int]struct{})
		myCardNumbers := make([]int, 0)
		currentNumber := ""

		winNumbersStart := 0
		myNumbersStart := 0

		for i := 0; i < len(rawCard); i++ {
			if rawCard[i] == ':' {
				noSpaceRawCardName := strings.Replace(rawCard[:i], " ", "", -1)
				cardNumber, _ = strconv.Atoi(strings.Split(noSpaceRawCardName, "Card")[1])
				winNumbersStart = i + 2
			}

			if rawCard[i] == '|' {
				myNumbersStart = i + 2
			}

			if winNumbersStart > 0 && myNumbersStart == 0 {
				if rawCard[i] >= 48 && rawCard[i] <= 57 {
					currentNumber += string(rawCard[i])
				} else if currentNumber != "" {
					num, _ := strconv.Atoi(currentNumber)
					winNumbers[num] = struct{}{}
					currentNumber = ""
				}
			}

			if myNumbersStart > 0 {
				isANumber := false
				if rawCard[i] >= 48 && rawCard[i] <= 57 {
					currentNumber += string(rawCard[i])
					isANumber = true
				}

				if (!isANumber && currentNumber != "") || i == len(rawCard)-1 {
					num, _ := strconv.Atoi(currentNumber)
					myCardNumbers = append(myCardNumbers, num)
					currentNumber = ""
				}
			}
		}

		wonCardNumbers := make([]int, 0)
		for _, num := range myCardNumbers {
			_, found := winNumbers[num]
			if found {
				wonCardNumbers = append(wonCardNumbers, num)
			}
		}

		cardTotal := 0
		for i := 0; i < len(wonCardNumbers); i++ {
			if i == 0 {
				cardTotal += 1
			} else {
				cardTotal *= 2
			}
		}

		total += cardTotal

		fmt.Println(
			cardNumber, ":",
			// "\n--", winNumbers,
			// "\n--", myCardNumbers,
			// "\n--", wonCardNumbers,
			"=>", cardTotal,
		)
	}

	return total
}

func part2(input string) int {
	return 0
}
