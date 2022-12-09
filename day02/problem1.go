package day02

import (
	"fmt"
	"strings"
)

const (
	O_Rock     = "A"
	O_Paper    = "B"
	O_Scissors = "C"

	Rock     = "X"
	Paper    = "Y"
	Scissors = "Z"
)

func SolveProblem1(lines []string) {
	score := 0
	i := 0
	for i < len(lines) {
		currentLine := lines[i]
		i++

		split := strings.Split(currentLine, " ")
		oponent := split[0]
		me := split[1]

		calculateTotalOfTop3(oponent, me, &score)
	}

	fmt.Printf("Solution to 1 is: A score of %d :)\n", score)
}

func calculateTotalOfTop3(oponent, me string, score *int) {
	if oponent == O_Rock {
		if me == Rock {
			*score += 1 + 3
		} else if me == Paper {
			*score += 2 + 6
		} else if me == Scissors {
			*score += 3 + 0
		}
	} else if oponent == O_Paper {
		if me == Rock {
			*score += 1 + 0
		} else if me == Paper {
			*score += 2 + 3
		} else if me == Scissors {
			*score += 3 + 6
		}
	} else if oponent == O_Scissors {
		if me == Rock {
			*score += 1 + 6
		} else if me == Paper {
			*score += 2 + 0
		} else if me == Scissors {
			*score += 3 + 3
		}
	}
}
