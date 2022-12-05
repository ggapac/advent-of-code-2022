package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	readFile, fileScanner := getScannedFile("input.txt")

	firstRow := true
	instructions := false
	part1 := false

	var crateStacks []LinkedList

	for fileScanner.Scan() {

		currentLine := fileScanner.Text()

		// From the number of characters in the first row we can
		// calculate how many stacks we have and initialize the
		// linked lists.
		if firstRow {
			crateStacks = initializeStacks(currentLine, crateStacks)
			firstRow = false
		}

		// Check if we finished reading the structure and
		// are moving on to the instructions part.
		if len(currentLine) <= 1 {
			instructions = true

			// reverse
			for i := 0; i < len(crateStacks); i++ {
				crateStacks[i].Reverse()
			}
			continue
		}

		if !instructions { // Read the stacks.
			for pos, char := range currentLine {
				if unicode.IsLetter(char) {
					stackIx := pos / 4
					crateStacks[stackIx].insertAtHead(char)
				}
			}
		} else { // Read the instructions.
			var num, from, to int
			fmt.Sscanf(currentLine, "move %d from %d to %d\n", &num, &from, &to)

			if part1 {
				crateMover9000(crateStacks, num, from, to)
			} else {
				crateMover9001(crateStacks, num, from, to)
			}
		}
	}

	readFile.Close()

	// Print the top crates.
	result := ""
	for i := 0; i < len(crateStacks); i++ {
		el := crateStacks[i].head.data
		result = result + string(el)
	}

	fmt.Println(result)

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

func initializeStacks(currentLine string, crateStacks []LinkedList) []LinkedList {
	numStacks := 1 + len(currentLine)/4
	crateStacks = make([]LinkedList, numStacks)
	for i := 0; i < numStacks; i++ {
		crateStacks[i] = LinkedList{nil, 0}
	}
	return crateStacks
}

func crateMover9000(crateStacks []LinkedList, num int, from int, to int) {
	for i := 0; i < num; i++ {
		el := crateStacks[from-1].head.data
		crateStacks[from-1].deleteAtHead()
		crateStacks[to-1].insertAtHead(el)
	}
}

func crateMover9001(crateStacks []LinkedList, num int, from int, to int) {
	els := make([]rune, num)

	for i := 0; i < num; i++ {
		els[i] = crateStacks[from-1].head.data
		crateStacks[from-1].deleteAtHead()
	}

	for i := len(els) - 1; i >= 0; i-- {
		crateStacks[to-1].insertAtHead(els[i])
	}
}
