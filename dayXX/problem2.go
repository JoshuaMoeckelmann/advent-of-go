package dayXX

import (
	"bufio"
	"fmt"

	"github.com/JoshuaMoeckelmann/AOC/common"
)

func SolveProblem2(scanner *bufio.Scanner, lineCount int) {
	resultingValue := 0
	for scanner.Scan() {
		currentLine := scanner.Text()

		fmt.Printf("Current value is %s :)\n", currentLine)
	}

	fmt.Printf("Solution to 2 is: Max value is %d :)\n", resultingValue)
	common.CheckScannerForError(scanner)
}
