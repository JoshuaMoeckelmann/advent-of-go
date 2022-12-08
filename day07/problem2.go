package day07

import (
	"bufio"
	"fmt"

	"github.com/JoshuaMoeckelmann/AOC/common"
)

func SolveProblem2(scanner *bufio.Scanner, lineCount int) {
	// skip first line
	scanner.Scan()
	root := Node{
		name:     "/",
		children: make([]*Node, 0),
	}
	parseTree(scanner, &root)

	resultingValue := 0
	parseTreeDFS(&resultingValue, &root)
	spaceCurrentlyFree := 70000000 - root.value
	spaceToFree := 30000000 - spaceCurrentlyFree

	minSize := root.value
	findDirsBiggerThanSpaceToFree(&minSize, &root, spaceToFree)

	fmt.Printf("Min value to delete is %d :)\n", minSize)
	common.CheckScannerForError(scanner)
}

func findDirsBiggerThanSpaceToFree(minSize *int, node *Node, spaceToFree int) {
	for _, n := range node.children {
		if n.dir {
			if n.value > spaceToFree && n.value < *minSize {
				*minSize = n.value
			}

			findDirsBiggerThanSpaceToFree(minSize, n, spaceToFree)
		}
	}
}
