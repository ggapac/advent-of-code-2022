package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	readFile, fileScanner := getScannedFile("input.txt")

	part1 := false

	monkeys := make(map[int]Monkey)
	var ixMonkey Monkey
	ix := 0

	// Read
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if ix%7 == 0 {
			ixMonkey = Monkey{}
		} else if ix%7 < 6 {
			ixMonkey = parseInput(line, ixMonkey, ix)

			if ix%7 == 5 {
				monkeys[ix/7] = ixMonkey
			}
		}
		ix += 1
	}
	readFile.Close()

	// Monkey business
	var numRounds, relief int
	if part1 {
		numRounds = 20
		relief = 3
	} else {
		numRounds = 10000
		relief = getRelief(monkeys)
	}

	for i := 0; i < numRounds; i++ {
		for j := 0; j < len(monkeys); j++ {
			// loop through items
			m := monkeys[j]
			for _, item := range m.items {
				toMonkey, worry := throwTo(m, item, relief, part1)
				monkeys = receive(monkeys, toMonkey, worry)
				m.inspections += 1
			}
			m.items = make([]int, 0) // empty the item list
			monkeys[j] = m
		}

	}

	fmt.Println(getMonkeyBusiness(monkeys))
}

func getScannedFile(filename string) (*os.File, *bufio.Scanner) {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	return readFile, fileScanner
}

func parseInput(line string, m Monkey, ix int) Monkey {
	var num int

	switch ix % 7 {
	case 1:
		re := regexp.MustCompile("[0-9]+")
		items := re.FindAllString(line, -1)
		m.items = make([]int, len(items))
		for i := 0; i < len(items); i++ {
			m.items[i], _ = strconv.Atoi(items[i])
		}
	case 2:
		var operator string
		var element string
		fmt.Sscanf(line, "  Operation: new = old %s %s", &operator, &element)
		m.element = element
		m.operator = operator
	case 3:
		fmt.Sscanf(line, "  Test: divisible by %d", &num)
		m.div = num
	case 4:
		fmt.Sscanf(line, "    If true: throw to monkey %d", &num)
		m.trueId = num
	case 5:
		fmt.Sscanf(line, "    If false: throw to monkey %d", &num)
		m.falseId = num
	}

	return m
}
