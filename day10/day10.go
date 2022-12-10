package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, fileScanner := getScannedFile("input.txt")

	register := 1
	cycle := 0
	cycleOfInterest := 20
	increment := 40
	signalStrength := 0

	screen := make([][]string, 6)
	for i := range screen {
		screen[i] = make([]string, increment)
	}
	rowIx := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		var val int

		if line != "noop" {
			addx := strings.Fields(line)
			val, _ = strconv.Atoi(addx[1])

			screen = draw(rowIx, cycle%40, register-1, screen)
			cycle += 1

			if cycle%increment == 0 {
				rowIx += 1
			}
		}

		screen = draw(rowIx, cycle%40, register-1, screen)
		cycle += 1

		if cycleOfInterest <= cycle {
			signalStrength, cycleOfInterest = updateState(signalStrength, cycleOfInterest, increment, register)
		}

		if cycle%increment == 0 {
			rowIx += 1
		}

		register += val

	}
	readFile.Close()

	// Part 1
	fmt.Println(signalStrength)

	// Part 2
	print(screen)
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

func updateState(signalStrength int, cycleOfInterest int, increment int, register int) (int, int) {
	signalStrength += cycleOfInterest * register
	cycleOfInterest += increment
	return signalStrength, cycleOfInterest
}

func draw(rowIx int, cycle int, sprite int, screen [][]string) [][]string {
	if sprite == cycle || (sprite+1) == cycle || (sprite+2) == cycle {
		screen[rowIx][cycle] = "#"
	} else {
		screen[rowIx][cycle] = "."
	}
	return screen
}

func print(screen [][]string) {
	for i := 0; i < len(screen); i++ {
		for j := 0; j < len(screen[i]); j++ {
			fmt.Print(screen[i][j])
		}
		fmt.Println("")
	}
}
