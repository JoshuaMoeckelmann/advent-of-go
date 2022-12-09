package day06

import (
	"fmt"
)

func SolveProblem2(lines []string) {
	resultingValue := searchForMarker(14, lines[0])
	fmt.Printf("Solution to 2 is: Marker starts at %d :)\n", resultingValue)
}
