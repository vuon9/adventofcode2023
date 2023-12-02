package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vuon9/adventofcode2023/utils"
)

type expectation struct {
	Red   int
	Green int
	Blue  int
}

type turn struct {
	Red   int
	Green int
	Blue  int
}

func main() {
	input := utils.ReadFileOrPanic()
	// 2551
	fmt.Println("Part 1, result is: ", part1(input, expectation{Red: 12, Green: 13, Blue: 14}))
	// 62811
	fmt.Println("Part 2, result is: ", part2(input))
}

func part1(input string, exp expectation) string {
	total := 0

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		mainStr := strings.Split(line, ": ")
		rawGameID := strings.Split(mainStr[0], " ")[1]
		rawRounds := strings.Split(mainStr[1], "; ")

		isOK := true
		for _, rawRound := range rawRounds {
			rawTurns := strings.Split(rawRound, ", ")
			red, green, blue := 0, 0, 0
			for _, rawTurn := range rawTurns {
				bag := strings.Split(strings.TrimSpace(rawTurn), " ")
				noBag, _ := strconv.Atoi(bag[0])

				if bag[1] == "red" {
					red = noBag
				} else if bag[1] == "green" {
					green = noBag
				} else if bag[1] == "blue" {
					blue = noBag
				}
			}

			if red > exp.Red || green > exp.Green || blue > exp.Blue {
				isOK = false
				break
			}
		}

		if isOK {
			gameID, _ := strconv.Atoi(rawGameID)
			total += gameID
		}
	}

	return fmt.Sprintf("%d", total)
}

func part2(input string) string {
	total := 0

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		mainStr := strings.Split(line, ": ")
		rawRounds := strings.Split(mainStr[1], "; ")

		mRed, mGreen, mBlue := 0, 0, 0
		for _, rawRound := range rawRounds {
			rawTurns := strings.Split(rawRound, ", ")
			red, green, blue := 0, 0, 0
			for _, rawTurn := range rawTurns {
				bag := strings.Split(strings.TrimSpace(rawTurn), " ")
				noBag, _ := strconv.Atoi(bag[0])

				if bag[1] == "red" {
					red = noBag
				} else if bag[1] == "green" {
					green = noBag
				} else if bag[1] == "blue" {
					blue = noBag
				}
			}

			mRed = max(mRed, red)
			mGreen = max(mGreen, green)
			mBlue = max(mBlue, blue)
		}

		total += mRed * mGreen * mBlue
	}

	return fmt.Sprintf("%d", total)
}
