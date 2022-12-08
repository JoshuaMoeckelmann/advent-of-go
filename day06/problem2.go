package day06

import (
	"bufio"
	"fmt"

	"github.com/JoshuaMoeckelmann/AOC/common"
)

func SolveProblem2(scanner *bufio.Scanner, lineCount int) {
	scanner.Scan()
	text := scanner.Text()
	resultingValue := searchForMarker(14, text)
	fmt.Printf("Solution to 2 is: Marker starts at %d :)\n", resultingValue)
	common.CheckScannerForError(scanner)
}
