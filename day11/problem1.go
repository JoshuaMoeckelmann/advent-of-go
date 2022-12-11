package day11

import (
	"fmt"
	"sort"
	"strings"

	"github.com/JoshuaMoeckelmann/advent-of-go/common"
	"github.com/thoas/go-funk"
)

type monkey struct {
	startingItems    queue
	performOperation func(old int) int
	divident         int
	monkeyIfTrue     int
	monkeyIfFalse    int
	itemsInspected   int
}

func SolveProblem1(lines []string) {
	// resultingValue := 0
	monkeys := parseMonkeys(lines)

	for i := 0; i < 20; i++ {
		for i, monkey := range monkeys {
			items := monkey.startingItems
			for _, v := range items {
				new := monkey.performOperation(v)
				newDividedAndFloored := new / 3
				if newDividedAndFloored%monkey.divident == 0 {
					otherMonkey := monkeys[monkey.monkeyIfTrue]
					otherMonkey.startingItems = otherMonkey.startingItems.Push(newDividedAndFloored)
					monkeys[monkey.monkeyIfTrue] = otherMonkey
				} else {
					otherMonkey := monkeys[monkey.monkeyIfFalse]
					otherMonkey.startingItems = otherMonkey.startingItems.Push(newDividedAndFloored)
					monkeys[monkey.monkeyIfFalse] = otherMonkey
				}
				monkey.itemsInspected += 1
			}

			monkey.startingItems = make([]int, 0)
			monkeys[i] = monkey
		}
	}

	itemsInspected := funk.Map(monkeys, func(m monkey) int {
		return m.itemsInspected
	}).([]int)
	sort.Ints(itemsInspected)
	fmt.Printf("Monkey %d items :)\n", itemsInspected[len(itemsInspected)-1]*itemsInspected[len(itemsInspected)-2])

	for i, m := range monkeys {
		fmt.Printf("Monkey %d inspected %d items :)\n", i, m.itemsInspected)
	}
}

func parseMonkeys(lines []string) []monkey {
	monkeys := make([]monkey, 0)
	for i := 0; i+6 <= len(lines); i += 7 {
		monkey := monkey{}
		parseItems(&monkey, lines[i+1])
		parseOperation(lines[i+2], &monkey)
		monkey.divident = common.ConvertStringToInt(strings.Split(lines[i+3], "divisible by")[1])
		monkey.monkeyIfTrue = common.ConvertStringToInt(strings.Split(lines[i+4], "throw to monkey")[1])
		monkey.monkeyIfFalse = common.ConvertStringToInt(strings.Split(lines[i+5], "throw to monkey")[1])
		monkey.itemsInspected = 0
		monkeys = append(monkeys, monkey)
	}
	return monkeys
}

func parseItems(monkey *monkey, s string) {
	startingItems := strings.Split(strings.Split(s, "items: ")[1], ", ")
	monkey.startingItems = make(queue, 0)
	for _, v := range startingItems {
		monkey.startingItems = monkey.startingItems.Push(common.ConvertStringToInt(v))
	}
}

func parseOperation(operationLine string, monkey *monkey) {
	rawOperation := strings.Split(operationLine, ": new = ")[1]
	operators := strings.Split(rawOperation, " ")

	monkey.performOperation = func(old int) int {
		firstValue := getOldOrValue(operators[0], old)
		operation := operators[1]
		secondValue := getOldOrValue(operators[2], old)

		switch operation {
		case "+":
			return firstValue + secondValue
		case "*":
			return firstValue * secondValue
		default:
			panic("unknown Operation")
		}
	}
}

func getOldOrValue(s string, old int) int {
	if s == "old" {
		return old
	}

	return common.ConvertStringToInt(s)
}
