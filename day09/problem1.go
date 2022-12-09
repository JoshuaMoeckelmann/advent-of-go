package day09

import (
	"fmt"
	"strings"

	"github.com/JoshuaMoeckelmann/advent-of-go/common"
	mapset "github.com/deckarep/golang-set/v2"
)

type movement struct {
	xDiff int
	yDiff int
}

var movementMap = initMovementMap()

func SolveProblem1(lines []string) {
	s := len(lines) / 2
	xH := s
	yH := s
	xT := s
	yT := s
	visitedPlaces := mapset.NewSet(position{
		x: xT,
		y: yT,
	})
	for _, line := range lines {
		movement, amount := parseLine(line)

		for i := 0; i < amount; i++ {
			// Move head
			xH += movement.xDiff
			yH += movement.yDiff

			xDiff := common.Abs(xH - xT)
			yDiff := common.Abs(yH - yT)

			signedMovementX := common.SignOrNull(xH - xT)
			signedMovementY := common.SignOrNull(yH - yT)
			if xDiff >= 2 || yDiff >= 2 {
				xT += signedMovementX
				yT += signedMovementY
			}
			visitedPlaces.Add(position{
				x: xT,
				y: yT,
			})
		}
	}

	fmt.Printf("Solution to 1 is: %d Points visited :)\n", visitedPlaces.Cardinality())
}

func parseLine(line string) (movement, int) {
	split := strings.Split(line, " ")
	return movementMap[split[0]], common.ConvertStringToInt(split[1])
}

func initMovementMap() map[string]movement {
	return map[string]movement{
		"R": {
			xDiff: 1,
			yDiff: 0,
		},
		"L": {
			xDiff: -1,
			yDiff: 0,
		},
		"U": {
			xDiff: 0,
			yDiff: -1,
		},
		"D": {
			xDiff: 0,
			yDiff: 1,
		},
	}
}
