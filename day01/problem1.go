package day01

import (
	"bufio"
	"fmt"

	"github.com/JoshuaMoeckelmann/AOC/common"
)

func SolveProblem1(scanner *bufio.Scanner, lineCount int) {
	maxValue := 0
	currentValue := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		performActionOnEachLine(&maxValue, &currentValue, currentLine)
	}

	fmt.Printf("Solution to 1 is: %d calories :)\n", maxValue)
	common.CheckScannerForError(scanner)
}

func performActionOnEachLine(maxValue, currentValue *int, currentLine string) {
	if currentLine == "" {
		*currentValue = 0
	} else {
		*currentValue += common.ConvertStringToInt(currentLine)
		if *currentValue > *maxValue {
			*maxValue = *currentValue
		}
	}
}
