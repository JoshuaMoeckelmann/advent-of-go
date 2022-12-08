package day08

import (
	"bufio"
	"fmt"

	"github.com/JoshuaMoeckelmann/AOC/common"
)

func SolveProblem2(scanner *bufio.Scanner, lineCount int) {
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

	maxVisibleTrees := 0
	for y := 1; y < height-1; y++ {
		for x := 1; x < lineLength-1; x++ {
			currentVisibleTrees := visibleTreesTop(grid, x, y) * visibleTreesBottom(grid, x, y, height) * visibleTreesLeft(grid, x, y) * visibleTreesRight(grid, x, y, lineLength)
			if currentVisibleTrees > maxVisibleTrees {
				maxVisibleTrees = currentVisibleTrees
			}
		}
	}

	fmt.Printf("Solution to 2 is: %d trees are visible :)\n", maxVisibleTrees)
	common.CheckScannerForError(scanner)
}

func visibleTreesTop(grid [][]int, x_init, y_init int) int {
	v := grid[y_init][x_init]
	vis := 1
	for y := y_init - 1; y >= 0; y-- {
		if grid[y][x_init] >= v || y == 0 {
			return vis
		} else {
			vis++
		}
	}
	return vis
}

func visibleTreesBottom(grid [][]int, x_init, y_init, height int) int {
	v := grid[y_init][x_init]
	vis := 1
	for y := y_init + 1; y < height; y++ {
		if grid[y][x_init] >= v || y == height-1 {
			return vis
		} else {
			vis++
		}
	}
	return vis
}

func visibleTreesLeft(grid [][]int, x_init, y_init int) int {
	v := grid[y_init][x_init]
	vis := 1
	for x := x_init - 1; x >= 0; x-- {
		if grid[y_init][x] >= v || x == 0 {
			return vis
		} else {
			vis++
		}
	}
	return vis
}

func visibleTreesRight(grid [][]int, x_init, y_init, width int) int {
	v := grid[y_init][x_init]
	vis := 1
	for x := x_init + 1; x < width; x++ {
		if grid[y_init][x] >= v || x == width-1 {
			return vis
		} else {
			vis++
		}
	}
	return vis
}
