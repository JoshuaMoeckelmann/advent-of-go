package day04

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/JoshuaMoeckelmann/AOC/common"
)

func SolveProblem1(scanner *bufio.Scanner, lineCount int) {
	resultingValue := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		lineSplit := strings.Split(currentLine, ",")

		leftLeft, leftRight := divideStringFurther(lineSplit[0])
		rightLeft, rightRight := divideStringFurther(lineSplit[1])

		if (leftLeft >= rightLeft && leftRight <= rightRight) || (leftLeft <= rightLeft && leftRight >= rightRight) {
			resultingValue++
		}
	}

	fmt.Printf("Solution to 1 is: %d :)\n", resultingValue)
	common.CheckScannerForError(scanner)
}

func divideStringFurther(s string) (int, int) {
	split := strings.Split(s, "-")

	return common.ConvertStringToInt(split[0]), common.ConvertStringToInt(split[1])
}
