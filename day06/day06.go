package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	readFile, fileScanner := getScannedFile("input.txt")

	var alphabet = map[byte]int{}
	markerLen := 14

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		for pos := range currentLine {

			if pos >= markerLen {
				// check if you found the marker
				if checkIfUnique(alphabet, markerLen) {
					fmt.Println(currentLine[pos-markerLen : pos]) // print the marker
					fmt.Println(pos)                              // print the position of the marker
					break
				}

				// -1 for the character that is out of scope
				alphabet[currentLine[pos-markerLen]] -= 1
			}

			// +1 for the newest character
			alphabet[currentLine[pos]] += 1
		}
	}

	readFile.Close()
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

func checkIfUnique(alphabet map[byte]int, markerLen int) bool {

	for _, val := range alphabet {
		if val > 1 {
			return false
		}
	}

	return true
}
