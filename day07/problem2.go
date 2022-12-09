package day07

import (
	"fmt"
)

func SolveProblem2(lines []string) {
	// skip first line
	i := 0
	root := Node{
		name:     "/",
		children: make([]*Node, 0),
	}
	parseTree(&i, lines, &root)

	resultingValue := 0
	parseTreeDFS(&resultingValue, &root)
	spaceCurrentlyFree := 70000000 - root.value
	spaceToFree := 30000000 - spaceCurrentlyFree

	minSize := root.value
	findDirsBiggerThanSpaceToFree(&minSize, &root, spaceToFree)

	fmt.Printf("Min value to delete is %d :)\n", minSize)
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
