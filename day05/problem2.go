package day05

import (
	"fmt"
)

func SolveProblem2(lines []string) {
	i := 0
	stacks := createStacks(&i, lines)
	for i < len(lines) {
		amount, from, to := parseInput(lines[i])

		s := make([]string, 0)
		for i := 0; i < amount; i++ {
			v := ""
			stacks[from-1], v = stacks[from-1].Pop()
			s = append(s, v)
		}

		for i := len(s) - 1; i >= 0; i-- {
			stacks[to-1] = stacks[to-1].Push(s[i])
		}
		i++
	}

	resultingValue := ""
	for _, sta := range stacks {
		resultingValue += sta.Peek()
	}
	fmt.Printf("Solution to 2 is: %s :)\n", resultingValue)
}
