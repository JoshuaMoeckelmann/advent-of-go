package day05

import (
	"fmt"
	"strings"

	"github.com/JoshuaMoeckelmann/advent-of-go/common"
)

func SolveProblem1(lines []string) {
	i := 0
	stacks := createStacks(&i, lines)
	for i < len(lines) {
		amount, from, to := parseInput(lines[i])

		for i := 0; i < amount; i++ {
			v := ""
			stacks[from-1], v = stacks[from-1].Pop()
			stacks[to-1] = stacks[to-1].Push(v)
		}
		i++
	}

	resultingValue := ""
	for _, sta := range stacks {
		resultingValue += sta.Peek()
	}
	fmt.Printf("Solution to 1 is: %s :)\n", resultingValue)
}

func createStacks(i *int, lines []string) []stack {
	// TODO make generic
	inputLength := 8
	stackLength := 9
	stacks := make([]stack, stackLength)
	for *i < inputLength {
		currentLine := lines[*i]

		for i := 0; i < stackLength; i++ {
			currentChar := string(currentLine[1+4*i])
			if currentChar != " " {
				stacks[i] = stacks[i].InvertedAndWrongPush(currentChar)
			}
		}
		*i++
	}

	// read next two lines, so that parsing works further on
	*i += 2
	return stacks
}

func parseInput(s string) (int, int, int) {
	firstSplit := strings.Split(s, "from")
	amountSplit := strings.Split(firstSplit[0], "move")
	fromToSplit := strings.Split(firstSplit[1], "to")

	return common.ConvertStringToInt(amountSplit[1]), common.ConvertStringToInt(fromToSplit[0]), common.ConvertStringToInt(fromToSplit[1])
}
