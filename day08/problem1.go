package day08

import (
	"bufio"
	"fmt"

	"github.com/JoshuaMoeckelmann/advent-of-go/common"
)

func SolveProblem1(scanner *bufio.Scanner, lineCount int) {
	grid := make([][]int, lineCount)
	height := 0
	lineLength := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		lineLength = len(currentLine)
		grid[height] = make([]int, lineLength)
		for x, v := range currentLine {
			grid[height][x] = common.ConvertStringToInt(string(v))
		}
		height++
	}

	// // negative 2 so that the no outer tree is counted twice
	visibleTrees := 2*height + 2*(lineLength-2)
	for y := 1; y < height-1; y++ {
		for x := 1; x < lineLength-1; x++ {
			if visibleTop(grid, x, y) || visibleBottom(grid, x, y, height) || visibleLeft(grid, x, y) || visibleRight(grid, x, y, lineLength) {
				visibleTrees++
			}
		}
	}

	fmt.Printf("Solution to 1 is: %d trees are visible :)\n", visibleTrees)
	common.CheckScannerForError(scanner)
}

func visibleTop(grid [][]int, x_init, y_init int) bool {
	v := grid[y_init][x_init]
	for y := y_init - 1; y >= 0; y-- {
		if grid[y][x_init] >= v {
			return false
		}
	}
	return true
}

func visibleBottom(grid [][]int, x_init, y_init, height int) bool {
	v := grid[y_init][x_init]
	for y := y_init + 1; y < height; y++ {
		if grid[y][x_init] >= v {
			return false
		}
	}
	return true
}

func visibleLeft(grid [][]int, x_init, y_init int) bool {
	v := grid[y_init][x_init]
	for x := x_init - 1; x >= 0; x-- {
		if grid[y_init][x] >= v {
			return false
		}
	}
	return true
}

func visibleRight(grid [][]int, x_init, y_init, width int) bool {
	v := grid[y_init][x_init]
	for x := x_init + 1; x < width; x++ {
		if grid[y_init][x] >= v {
			return false
		}
	}
	return true
}
