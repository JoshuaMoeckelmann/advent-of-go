package day03

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/JoshuaMoeckelmann/advent-of-go/common"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func SolveProblem1(scanner *bufio.Scanner, lineCount int) {
	resultingValue := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
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
	common.CheckScannerForError(scanner)
}
