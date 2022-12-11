package main

// all the monkey business is here

import "strconv"

type Monkey struct {
	items       []int
	operator    string
	element     string
	div         int
	trueId      int
	falseId     int
	inspections int
}

func throwTo(m Monkey, item int, relief int, part1 bool) (int, int) {
	// get worry level
	worry := item
	element := m.element
	if element == "old" {
		if m.operator == "+" {
			worry = worry * 2
		} else if m.operator == "*" {
			worry = worry * worry
		}
	} else {
		num, _ := strconv.Atoi(m.element)
		if m.operator == "+" {
			worry = worry + num
		} else if m.operator == "*" {
			worry = worry * num
		}
	}

	// worry manager
	if part1 {
		worry /= relief
	} else {
		worry = worry % relief
	}

	// check who to throw
	if worry%m.div == 0 {
		return m.trueId, worry
	}
	return m.falseId, worry
}

func receive(monkeys map[int]Monkey, toMonkey int, worry int) map[int]Monkey {
	tmpMonkey := monkeys[toMonkey]
	tmpMonkey.items = append(tmpMonkey.items, worry)
	monkeys[toMonkey] = tmpMonkey
	return monkeys
}

func getMonkeyBusiness(monkeys map[int]Monkey) int {
	var m1, m2 int
	for i := 0; i <= len(monkeys); i++ {
		if monkeys[i].inspections >= m1 {
			m2 = m1
			m1 = monkeys[i].inspections
		} else if monkeys[i].inspections > m2 {
			m2 = monkeys[i].inspections
		}
	}

	return m1 * m2
}

func getRelief(monkeys map[int]Monkey) int {
	relief := 1
	for i := 0; i < len(monkeys); i++ {
		relief *= monkeys[i].div
	}
	return relief
}
