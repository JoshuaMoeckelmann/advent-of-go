package main

import (
	"flag"

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
	purgeCachedFile := flag.Bool("purgeCachedFile", false, "Remove cached input file and force a redownload")
	flag.Parse()

	lines := parseLocalFlag(local, token, day, purgeCachedFile)
	switchDays(day, lines)
}

func switchDays(day *int, lines []string) {
	switch *day {
	case 1:
		day01.SolveProblem1(lines)
		day01.SolveProblem2(lines)
	case 2:
		day02.SolveProblem1(lines)
		day02.SolveProblem2(lines)
	case 3:
		day03.SolveProblem1(lines)
		day03.SolveProblem2(lines)
	case 4:
		day04.SolveProblem1(lines)
		day04.SolveProblem2(lines)
	case 5:
		day05.SolveProblem1(lines)
		day05.SolveProblem2(lines)
	case 6:
		day06.SolveProblem1(lines)
		day06.SolveProblem2(lines)
	case 7:
		day07.SolveProblem1(lines)
		day07.SolveProblem2(lines)
	case 8:
		day08.SolveProblem1(lines)
		day08.SolveProblem2(lines)
	case 9:
		day09.SolveProblem1(lines)
		day09.SolveProblem2(lines)
	}
}

func parseLocalFlag(local *bool, token *string, day *int, purgeCachedFile *bool) []string {
	if *local {
		return common.LoadLocalFile()
	}

	return common.PrepareFile(*token, *day, *purgeCachedFile)
}
