package day11

import (
	"fmt"
	"sort"

	"github.com/thoas/go-funk"
)

func SolveProblem2(lines []string) {
	monkeys := parseMonkeys(lines)
	p := 1
	for _, m := range monkeys {
		p *= m.divident
	}
	for i := 0; i < 10000; i++ {
		for i, monkey := range monkeys {
			items := monkey.startingItems
			for _, v := range items {
				new := monkey.performOperation(v)
				newDividedAndFloored := new % p
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
