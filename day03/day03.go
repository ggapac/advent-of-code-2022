package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	readFile, fileScanner := getScannedFile("input.txt")

	prioritiesSum := 0

	// Part 1
	//groupSize := 1
	//numCompartments := 2
	//numComparing := 2

	// Part 2
	groupSize := 3
	numCompartments := 1
	numComparing := 3

	elfIx := 0
	var elfBackpack [52]int

	for fileScanner.Scan() {
		backpack := fileScanner.Text()

		if elfIx%groupSize == 0 { // reset
			elfBackpack = [len(elfBackpack)]int{}
		}

		compartments := splitBackpack(backpack, numCompartments)

		for _, comp := range compartments {

			var elfUnique [52]bool

			for _, char := range comp {
				compIx := -1
				if unicode.IsUpper(char) {
					compIx = int(rune(char)-38) - 1
				} else {
					compIx = int(rune(char)-96) - 1
				}

				if !elfUnique[compIx] {
					elfBackpack[compIx] += 1
				}

				elfUnique[compIx] = true
			}

			for i := 0; i < len(elfBackpack); i++ {
				if elfBackpack[i] == numComparing {
					prioritiesSum += (i + 1)
				}
			}
		}

		elfIx += 1
	}

	readFile.Close()

	// Result
	fmt.Println(prioritiesSum)
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

func splitBackpack(backpack string, numCompartments int) []string {
	compartments := make([]string, numCompartments)

	for i := 0; i < numCompartments; i++ {
		compartments[i] = backpack[i*(len(backpack)/numCompartments) : (i+1)*(len(backpack)/numCompartments)]
	}
	return compartments
}
