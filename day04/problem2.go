package day04

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/JoshuaMoeckelmann/AOC/common"
)

func SolveProblem2(scanner *bufio.Scanner, lineCount int) {
	resultingValue := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		lineSplit := strings.Split(currentLine, ",")

		leftLeft, leftRight := divideStringFurther(lineSplit[0])
		rightLeft, rightRight := divideStringFurther(lineSplit[1])

		if leftLeft <= rightRight && rightLeft <= leftRight {
			resultingValue++
		}
	}

	fmt.Printf("Solution to 2 is: %d :)\n", resultingValue)
	common.CheckScannerForError(scanner)
}
