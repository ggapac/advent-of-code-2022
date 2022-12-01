package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var calories []int
	elfWeight := 0
	maxElfWeight := 0

	readFile, fileScanner := getScannedFile("input.txt")

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		if currentLine == "" {
			calories = append(calories, elfWeight)

			if elfWeight > maxElfWeight {
				maxElfWeight = elfWeight
			}

			elfWeight = 0
		} else {
			currentWeight, err := strconv.Atoi(currentLine)
			if err != nil {
				fmt.Println(err)
			}
			elfWeight += currentWeight
		}

	}

	readFile.Close()

	// Solution, part 1
	fmt.Println(maxElfWeight)

	// Solution, part 2
	printTop3Elfs(calories)
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

func printTop3Elfs(calories []int) {

	elf1 := 0
	elf2 := 0
	elf3 := 0

	for _, elfCalories := range calories {
		if elfCalories > elf3 {
			elf3 = elfCalories
		}

		if elfCalories > elf2 {
			elf2, elf3 = elfCalories, elf2
		}

		if elfCalories > elf1 {
			elf1, elf2, elf3 = elfCalories, elf1, elf2
		}
	}

	fmt.Println("Top 3 elfs:")
	fmt.Println(elf1)
	fmt.Println(elf2)
	fmt.Println(elf3)

	fmt.Println("Sum:")
	fmt.Println(elf1 + elf2 + elf3)
}
