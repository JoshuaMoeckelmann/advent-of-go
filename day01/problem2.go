package day01

import (
	"bufio"
	"fmt"
	"sort"

	"github.com/JoshuaMoeckelmann/AOC/common"
)

func SolveProblem2(scanner *bufio.Scanner, lineCount int) {
	values := make([]int, 0)
	currentValue := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		if currentLine == "" {
			values = append(values, currentValue)
			currentValue = 0
		} else {
			currentValue += common.ConvertStringToInt(currentLine)
		}
	}

	fmt.Printf("Solution to 2 is: %d calories :)\n", calculateTotalOfTop3(values))
	common.CheckScannerForError(scanner)
}

func calculateTotalOfTop3(values []int) int {
	sort.Ints(values)
	size := len(values)
	return values[size-1] + values[size-2] + values[size-3]
}
