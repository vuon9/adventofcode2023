package utils

import (
	"os"
)

func ReadFileOrPanic() string {
	if len(os.Args) < 2 {
		panic("Please provide a file name.")
	}

	fileName := os.Args[1]
	input, err := os.ReadFile(fileName)
	if err != nil {
		panic("Error reading file:" + err.Error())
	}

	return string(input)
}
