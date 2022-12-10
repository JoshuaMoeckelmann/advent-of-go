package day10

import (
	"fmt"
	"strings"

	"github.com/JoshuaMoeckelmann/advent-of-go/common"
)

func SolveProblem1(lines []string) {
	resultingValue := 0
	x := 1
	cycle := 0
	for _, currentLine := range lines {
		split := strings.Split(currentLine, " ")
		command := split[0]

		switch command {
		case "addx":
			amount := common.ConvertStringToInt(split[1])
			cycle++
			checkCycle(cycle, x, &resultingValue)
			cycle++
			checkCycle(cycle, x, &resultingValue)
			x += amount
		case "noop":
			cycle++
			checkCycle(cycle, x, &resultingValue)
		}
	}

	fmt.Printf("Solution to 1 is: Max value is %d :)\n", resultingValue)
}

func checkCycle(cycle, x int, resultingValue *int) {

	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		*resultingValue += cycle * x
	}
}
