package dayXX

import (
	"fmt"
)

func SolveProblem1(lines []string) {
	resultingValue := 0
	for _, currentLine := range lines {
		fmt.Printf("Current value is %s :)\n", currentLine)
	}

	fmt.Printf("Solution to 1 is: Max value is %d :)\n", resultingValue)
}
