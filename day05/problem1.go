package day05

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/JoshuaMoeckelmann/AOC/common"
)

func SolveProblem1(scanner *bufio.Scanner, lineCount int) {
	stacks := createStacks(scanner)
	for scanner.Scan() {
		amount, from, to := parseInput(scanner.Text())

		for i := 0; i < amount; i++ {
			v := ""
			stacks[from-1], v = stacks[from-1].Pop()
			stacks[to-1] = stacks[to-1].Push(v)
		}
	}

	resultingValue := ""
	for _, sta := range stacks {
		resultingValue += sta.Peek()
	}
	fmt.Printf("Solution to 1 is: %s :)\n", resultingValue)
	common.CheckScannerForError(scanner)
}

func createStacks(scanner *bufio.Scanner) []stack {
	// TODO make generic
	inputLength := 8
	stackLength := 9
	stacks := make([]stack, stackLength)
	i := 0
	for scanner.Scan() {
		i++
		if i > inputLength {
			// read next two lines, so that parsing works further on
			scanner.Text()
			scanner.Scan()
			scanner.Text()
			break
		}
		currentLine := scanner.Text()

		for i := 0; i < stackLength; i++ {
			currentChar := string(currentLine[1+4*i])
			if currentChar != " " {
				stacks[i] = stacks[i].InvertedAndWrongPush(currentChar)
			}
		}
	}
	return stacks
}

func parseInput(s string) (int, int, int) {
	firstSplit := strings.Split(s, "from")
	amountSplit := strings.Split(firstSplit[0], "move")
	fromToSplit := strings.Split(firstSplit[1], "to")

	return common.ConvertStringToInt(amountSplit[1]), common.ConvertStringToInt(fromToSplit[0]), common.ConvertStringToInt(fromToSplit[1])
}
