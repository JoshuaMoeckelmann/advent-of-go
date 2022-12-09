package day09

import (
	"fmt"

	"github.com/JoshuaMoeckelmann/advent-of-go/common"
	mapset "github.com/deckarep/golang-set/v2"
)

type position struct {
	x int
	y int
}

func SolveProblem2(lines []string) {
	s := len(lines) / 2
	snake := make([]position, 10)
	for i := range snake {
		snake[i] = position{
			x: s,
			y: s,
		}
	}

	visitedPlaces := mapset.NewSet(snake[len(snake)-1])
	for _, line := range lines {
		movement, amount := parseLine(line)

		for i := 0; i < amount; i++ {
			// Move head
			head := snake[0]
			head.x += movement.xDiff
			head.y += movement.yDiff
			snake[0] = head

			for j := 1; j < len(snake); j++ {
				prev := snake[j-1]
				curr := snake[j]

				xDiff := common.Abs(prev.x - curr.x)
				yDiff := common.Abs(prev.y - curr.y)

				signedMovementX := common.SignOrNull(prev.x - curr.x)
				signedMovementY := common.SignOrNull(prev.y - curr.y)

				if xDiff >= 2 || yDiff >= 2 {
					curr.x += signedMovementX
					curr.y += signedMovementY
				}
				snake[j] = curr
			}

			tail := snake[len(snake)-1]
			visitedPlaces.Add(tail)
		}
	}
	fmt.Printf("Solution to 2 is: snake visited %d places :)\n", visitedPlaces.Cardinality())
}
