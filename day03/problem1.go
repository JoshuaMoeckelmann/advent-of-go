package day03

import (
	"fmt"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func SolveProblem1(lines []string) {
	resultingValue := 0
	i := 0
	for i < len(lines) {
		currentLine := lines[i]
		i++

		split := len(currentLine) / 2
		leftPart := currentLine[:split]
		rightPart := currentLine[split:]
		for _, v := range leftPart {
			stringV := string(v)
			if strings.Contains(rightPart, stringV) {
				index := strings.Index(alphabet, stringV) + 1
				resultingValue += index
				break
			}
		}
	}

	fmt.Printf("Solution to 1 is: %d :)\n", resultingValue)
}
