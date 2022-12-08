package day03

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/JoshuaMoeckelmann/advent-of-go/common"
)

func SolveProblem2(scanner *bufio.Scanner, lineCount int) {
	resultingValue := 0
	firstPart := ""
	secondPart := ""
	thirdPart := ""
	for scanner.Scan() {
		firstPart = scanner.Text()
		scanner.Scan()
		secondPart = scanner.Text()
		scanner.Scan()
		thirdPart = scanner.Text()

		for _, v := range firstPart {
			stringV := string(v)
			if strings.Contains(secondPart, stringV) && strings.Contains(thirdPart, stringV) {
				index := strings.Index(alphabet, stringV) + 1
				resultingValue += index
				break
			}
		}
		resetValues(&firstPart, &secondPart, &thirdPart)
	}

	fmt.Printf("Solution to 2 is: %d :)\n", resultingValue)
	common.CheckScannerForError(scanner)
}

func resetValues(string1 *string, string2 *string, string3 *string) {
	*string1 = ""
	*string2 = ""
	*string3 = ""
}
