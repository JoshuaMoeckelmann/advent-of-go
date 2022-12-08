package day02

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/JoshuaMoeckelmann/AOC/common"
)

func SolveProblem2(scanner *bufio.Scanner, lineCount int) {
	score := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		split := strings.Split(currentLine, " ")
		oponent := split[0]
		me := split[1]

		meCalculated := calculateMe(oponent, me)

		calculateTotalOfTop3(oponent, meCalculated, &score)
	}

	fmt.Printf("Solution to 2 is: A score of %d :)\n", score)
	common.CheckScannerForError(scanner)
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
