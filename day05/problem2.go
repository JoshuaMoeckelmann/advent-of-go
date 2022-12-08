package day05

import (
	"bufio"
	"fmt"

	"github.com/JoshuaMoeckelmann/advent-of-go/common"
)

func SolveProblem2(scanner *bufio.Scanner, lineCount int) {
	stacks := createStacks(scanner)
	for scanner.Scan() {
		amount, from, to := parseInput(scanner.Text())

		s := make([]string, 0)
		for i := 0; i < amount; i++ {
			v := ""
			stacks[from-1], v = stacks[from-1].Pop()
			s = append(s, v)
		}

		for i := len(s) - 1; i >= 0; i-- {
			stacks[to-1] = stacks[to-1].Push(s[i])
		}
	}

	resultingValue := ""
	for _, sta := range stacks {
		resultingValue += sta.Peek()
	}
	fmt.Printf("Solution to 2 is: %s :)\n", resultingValue)
	common.CheckScannerForError(scanner)
}
