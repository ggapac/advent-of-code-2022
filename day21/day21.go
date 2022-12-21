package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	el1   string
	el2   string
	job   string
	value int
	fin   bool
}

const upperBound = 92233720368547  // some digits less than max int
const lowerBound = -92233720368547 // some digits less than min int

func main() {

	readFile, fileScanner := getScannedFile("input.txt")

	monkeys := make(map[string]Monkey)

	// read input
	for fileScanner.Scan() {
		line := fileScanner.Text()
		eq := strings.Fields(line)

		if len(eq) > 2 {
			monkeys[strings.TrimSuffix(eq[0], ":")] = Monkey{el1: eq[1], el2: eq[3], job: eq[2]}
		} else {
			num, _ := strconv.Atoi(eq[1])
			monkeys[strings.TrimSuffix(eq[0], ":")] = Monkey{value: num, fin: true}
		}
	}
	readFile.Close()

	// root result
	part1, _ := getValues(monkeys["root"], monkeys, false)
	fmt.Println(part1)

	// find out in which branch is humn
	left, humnLeft := getValues(monkeys[monkeys["root"].el1], monkeys, false)
	right, _ := getValues(monkeys[monkeys["root"].el2], monkeys, false)

	// binary search to find humn value
	if humnLeft {
		part2, found := binarySearch(right, lowerBound, upperBound, monkeys[monkeys["root"].el1], monkeys)
		fmt.Println(part2, found)
	} else {
		part2, found := binarySearch(left, lowerBound, upperBound, monkeys[monkeys["root"].el2], monkeys)
		fmt.Println(part2, found)
	}
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

func binarySearch(needle int, low int, high int, m Monkey, monkeys map[string]Monkey) (int, bool) {
	found := false

	for low < high {
		median := (low + high) / 2
		tmp := monkeys["humn"]
		tmp.value = median
		monkeys["humn"] = tmp
		result, _ := getValues(m, monkeys, false)

		if result == needle {
			found = true
		}

		if result > needle {
			low = median + 1
		} else {
			high = median
		}
	}

	return low, found
}

func getValues(m Monkey, monkeys map[string]Monkey, humn bool) (int, bool) {
	if m.fin {
		return m.value, humn
	}

	if m.el1 == "humn" || m.el2 == "humn" {
		humn = true
	}

	left, humn := getValues(monkeys[m.el1], monkeys, humn)
	right, humn := getValues(monkeys[m.el2], monkeys, humn)

	if m.job == "+" {
		return left + right, humn
	} else if m.job == "-" {
		return left - right, humn
	} else if m.job == "*" {
		return left * right, humn
	}
	return left / right, humn
}
