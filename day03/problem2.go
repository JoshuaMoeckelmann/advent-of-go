package day03

import (
	"fmt"
	"strings"
)

func SolveProblem2(lines []string) {
	resultingValue := 0
	i := 2
	for i < len(lines) {
		firstPart := lines[i-2]
		secondPart := lines[i-1]
		thirdPart := lines[i]
		i += 3

		for _, v := range firstPart {
			stringV := string(v)
			if strings.Contains(secondPart, stringV) && strings.Contains(thirdPart, stringV) {
				index := strings.Index(alphabet, stringV) + 1
				resultingValue += index
				break
			}
		}
	}

	fmt.Printf("Solution to 2 is: %d :)\n", resultingValue)
}
