package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/JoshuaMoeckelmann/advent-of-go/common"
	"github.com/JoshuaMoeckelmann/advent-of-go/day01"
	"github.com/JoshuaMoeckelmann/advent-of-go/day02"
	"github.com/JoshuaMoeckelmann/advent-of-go/day03"
	"github.com/JoshuaMoeckelmann/advent-of-go/day04"
	"github.com/JoshuaMoeckelmann/advent-of-go/day05"
	"github.com/JoshuaMoeckelmann/advent-of-go/day06"
	"github.com/JoshuaMoeckelmann/advent-of-go/day07"
	"github.com/JoshuaMoeckelmann/advent-of-go/day08"
	"github.com/JoshuaMoeckelmann/advent-of-go/day09"
)

func main() {
	token := flag.String("token", "", "Token to request the AOC input")
	day := flag.Int("day", 1, "day to run the AOC of")
	local := flag.Bool("local", false, "Use local input file")
	flag.Parse()

	scanner, fileRowLength, fileToClose, lines := parseLocalFlag(local, token, day)
	defer fileToClose.Close()

	switchDays(day, scanner, fileRowLength, fileToClose, lines)
}

func switchDays(day *int, scanner *bufio.Scanner, fileRowLength int, fileToClose *os.File, lines []string) {
	switch *day {
	case 1:
		day01.SolveProblem1(scanner, fileRowLength)
		day01.SolveProblem2(common.CreateNewScanner(fileToClose), fileRowLength)
	case 2:
		day02.SolveProblem1(scanner, fileRowLength)
		day02.SolveProblem2(common.CreateNewScanner(fileToClose), fileRowLength)
	case 3:
		day03.SolveProblem1(scanner, fileRowLength)
		day03.SolveProblem2(common.CreateNewScanner(fileToClose), fileRowLength)
	case 4:
		day04.SolveProblem1(scanner, fileRowLength)
		day04.SolveProblem2(common.CreateNewScanner(fileToClose), fileRowLength)
	case 5:
		day05.SolveProblem1(scanner, fileRowLength)
		day05.SolveProblem2(common.CreateNewScanner(fileToClose), fileRowLength)
	case 6:
		day06.SolveProblem1(scanner, fileRowLength)
		day06.SolveProblem2(common.CreateNewScanner(fileToClose), fileRowLength)
	case 7:
		day07.SolveProblem1(scanner, fileRowLength)
		day07.SolveProblem2(common.CreateNewScanner(fileToClose), fileRowLength)
	case 8:
		day08.SolveProblem1(scanner, fileRowLength)
		day08.SolveProblem2(common.CreateNewScanner(fileToClose), fileRowLength)
	case 9:
		day09.SolveProblem1(lines)
		day09.SolveProblem2(lines)
	}
}

func parseLocalFlag(local *bool, token *string, day *int) (*bufio.Scanner, int, *os.File, []string) {
	if *local {
		return common.LoadLocalFile()
	}

	return common.PrepareFile(*token, *day)
}
