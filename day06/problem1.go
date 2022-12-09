package day06

import (
	"fmt"
	"strings"
)

func SolveProblem1(lines []string) {
	resultingValue := searchForMarker(4, lines[0])

	fmt.Printf("Solution to 1 is: Marker starts at %d :)\n", resultingValue)
}

func searchForMarker(length int, text string) int {
	internalLength := length - 1
	for i := internalLength; i < len(text); i++ {
		v := text[i-internalLength : i+1]
		valid := true
		for _, c := range v {
			if strings.Count(v, string(c)) > 1 {
				valid = false
				break
			}
		}

		if valid {
			return i + 1
		}
	}
	return -1
}
