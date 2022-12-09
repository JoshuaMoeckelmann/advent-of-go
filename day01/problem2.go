package day01

import (
	"fmt"
	"sort"

	"github.com/JoshuaMoeckelmann/advent-of-go/common"
)

func SolveProblem2(lines []string) {
	values := make([]int, 0)
	currentValue := 0
	i := 0
	for i < len(lines) {
		currentLine := lines[i]
		i++
		if currentLine == "" {
			values = append(values, currentValue)
			currentValue = 0
		} else {
			currentValue += common.ConvertStringToInt(currentLine)
		}
	}

	fmt.Printf("Solution to 2 is: %d calories :)\n", calculateTotalOfTop3(values))
}

func calculateTotalOfTop3(values []int) int {
	sort.Ints(values)
	size := len(values)
	return values[size-1] + values[size-2] + values[size-3]
}
