package day04

import (
	"fmt"
	"strings"

	"github.com/JoshuaMoeckelmann/advent-of-go/common"
)

func SolveProblem1(lines []string) {
	resultingValue := 0
	i := 0
	for i < len(lines) {
		currentLine := lines[i]
		i++

		lineSplit := strings.Split(currentLine, ",")
		leftLeft, leftRight := divideStringFurther(lineSplit[0])
		rightLeft, rightRight := divideStringFurther(lineSplit[1])

		if (leftLeft >= rightLeft && leftRight <= rightRight) || (leftLeft <= rightLeft && leftRight >= rightRight) {
			resultingValue++
		}
	}

	fmt.Printf("Solution to 1 is: %d :)\n", resultingValue)
}

func divideStringFurther(s string) (int, int) {
	split := strings.Split(s, "-")

	return common.ConvertStringToInt(split[0]), common.ConvertStringToInt(split[1])
}
