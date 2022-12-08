package day07

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/JoshuaMoeckelmann/advent-of-go/common"
)

func SolveProblem1(scanner *bufio.Scanner, lineCount int) {
	// skip first line
	scanner.Scan()
	root := Node{
		name:     "/",
		children: make([]*Node, 0),
	}
	parseTree(scanner, &root)

	resultingValue := 0
	parseTreeDFS(&resultingValue, &root)
	fmt.Printf("Max value is %d :)\n", resultingValue)
	common.CheckScannerForError(scanner)
}

func parseTreeDFS(wholeValue *int, node *Node) {
	for _, n := range node.children {
		if n.dir {
			parseTreeDFS(wholeValue, n)
		}

		node.value += n.value
	}

	if node.value < 100000 {
		*wholeValue += node.value
	}
}

func parseTree(scanner *bufio.Scanner, node *Node) *Node {
	for scanner.Scan() {
		currentLine := scanner.Text()
		if currentLine == "$ ls" {
			readDirectory(scanner, node)
			currentLine = scanner.Text()
			if currentLine == "" {
				break
			}
		}

		if currentLine == "$ cd .." {
			break
		} else {
			dirName := strings.Split(currentLine, " cd ")[1]
			for _, v := range node.children {
				if v.name == dirName {
					parseTree(scanner, v)
					break
				}
			}
		}
	}
	return node
}

func readDirectory(scanner *bufio.Scanner, node *Node) {
	for scanner.Scan() {
		currentLine := scanner.Text()
		if string(currentLine[0]) == "$" {
			break
		}
		split := strings.Split(currentLine, " ")
		if split[0] == "dir" {
			// add dir
			dir := Node{
				name:     split[1],
				children: make([]*Node, 0),
				dir:      true,
			}
			node.children = append(node.children, &dir)
		} else {
			// add file
			child := Node{
				name:  split[1],
				value: common.ConvertStringToInt(split[0]),
				dir:   false,
			}
			node.children = append(node.children, &child)
		}
	}
}
