package day06

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/JoshuaMoeckelmann/advent-of-go/common"
)

func SolveProblem1(scanner *bufio.Scanner, lineCount int) {
	scanner.Scan()
	text := scanner.Text()
	resultingValue := searchForMarker(4, text)

	fmt.Printf("Solution to 1 is: Marker starts at %d :)\n", resultingValue)
	common.CheckScannerForError(scanner)
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
