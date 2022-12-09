package day02

import (
	"fmt"
	"strings"
)

func SolveProblem2(lines []string) {
	score := 0
	i := 0
	for i < len(lines) {
		currentLine := lines[i]
		i++

		split := strings.Split(currentLine, " ")
		oponent := split[0]
		me := split[1]

		meCalculated := calculateMe(oponent, me)

		calculateTotalOfTop3(oponent, meCalculated, &score)
	}

	fmt.Printf("Solution to 2 is: A score of %d :)\n", score)
}

func calculateMe(oponent, me string) string {
	if oponent == O_Rock {
		if me == Rock {
			return Scissors
		} else if me == Paper {
			return Rock
		} else if me == Scissors {
			return Paper
		}
	} else if oponent == O_Paper {
		if me == Rock {
			return Rock
		} else if me == Paper {
			return Paper
		} else if me == Scissors {
			return Scissors
		}
	} else if oponent == O_Scissors {
		if me == Rock {
			return Paper
		} else if me == Paper {
			return Scissors
		} else if me == Scissors {
			return Rock
		}
	}

	return ""
}
