package day10

import (
	"fmt"
	"strings"

	"github.com/JoshuaMoeckelmann/advent-of-go/common"
)

func SolveProblem2(lines []string) {
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
			printCycle(cycle, x, &resultingValue)
			cycle++
			printCycle(cycle, x, &resultingValue)
			x += amount
		case "noop":
			cycle++
			printCycle(cycle, x, &resultingValue)
		}
	}

	fmt.Printf("Solution to 1 is: Max value is %d :)\n", resultingValue)
}

func printCycle(cycle, x int, resultingValue *int) {
	restOfCycle := cycle % 40
	if restOfCycle >= x && restOfCycle < x+3 {
		fmt.Print("🎅")
	} else {
		fmt.Print("🎄")
	}

	if restOfCycle == 0 {
		fmt.Println()
	}
}
