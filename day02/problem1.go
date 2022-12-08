package day02

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/JoshuaMoeckelmann/AOC/common"
)

const (
	O_Rock     = "A"
	O_Paper    = "B"
	O_Scissors = "C"

	Rock     = "X"
	Paper    = "Y"
	Scissors = "Z"
)

func SolveProblem1(scanner *bufio.Scanner, lineCount int) {
	score := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		split := strings.Split(currentLine, " ")
		oponent := split[0]
		me := split[1]

		calculateTotalOfTop3(oponent, me, &score)
	}

	fmt.Printf("Solution to 1 is: A score of %d :)\n", score)
	common.CheckScannerForError(scanner)
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
