package day04

import (
	"fmt"
	"strings"
)

func SolveProblem2(lines []string) {
	resultingValue := 0
	i := 0
	for i < len(lines) {
		currentLine := lines[i]
		i++

		lineSplit := strings.Split(currentLine, ",")
		leftLeft, leftRight := divideStringFurther(lineSplit[0])
		rightLeft, rightRight := divideStringFurther(lineSplit[1])

		if leftLeft <= rightRight && rightLeft <= leftRight {
			resultingValue++
		}
	}

	fmt.Printf("Solution to 2 is: %d :)\n", resultingValue)
}
