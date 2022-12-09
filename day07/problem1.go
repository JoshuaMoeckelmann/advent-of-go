package day07

import (
	"fmt"
	"strings"

	"github.com/JoshuaMoeckelmann/advent-of-go/common"
)

func SolveProblem1(lines []string) {
	i := 0
	root := Node{
		name:     "/",
		children: make([]*Node, 0),
	}
	parseTree(&i, lines, &root)

	resultingValue := 0
	parseTreeDFS(&resultingValue, &root)
	fmt.Printf("Max value is %d :)\n", resultingValue)
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

func parseTree(i *int, lines []string, node *Node) *Node {
	for *i+1 < len(lines) {
		*i++
		currentLine := lines[*i]
		if currentLine == "$ ls" {
			readDirectory(i, lines, node)
			if *i+1 >= len(lines) {
				break
			}
			currentLine = lines[*i]
		}

		if currentLine == "$ cd .." {
			break
		} else {
			dirName := strings.Split(currentLine, " cd ")[1]
			for _, v := range node.children {
				if v.name == dirName {
					parseTree(i, lines, v)
					break
				}
			}
		}
	}
	return node
}

func readDirectory(i *int, lines []string, node *Node) {
	for *i+1 < len(lines) {
		*i++
		currentLine := lines[*i]
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
