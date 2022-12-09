package day01

import (
	"fmt"

	"github.com/JoshuaMoeckelmann/advent-of-go/common"
)

func SolveProblem1(lines []string) {
	maxValue := 0
	currentValue := 0
	i := 0
	for i < len(lines) {
		currentLine := lines[i]
		i++
		performActionOnEachLine(&maxValue, &currentValue, currentLine)
	}

	fmt.Printf("Solution to 1 is: %d calories :)\n", maxValue)
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
