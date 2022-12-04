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

	numFullyContained := 0
	numOverlapped := 0

	for fileScanner.Scan() {
		pair := strings.Split(fileScanner.Text(), ",")
		elf1string := strings.Split(pair[0], "-")
		elf2string := strings.Split(pair[1], "-")

		elf1 := stringArrayToInt(elf1string)
		elf2 := stringArrayToInt(elf2string)

		if areFullyContained(elf1, elf2) {
			numFullyContained += 1
		}

		if areOverlapped(elf1, elf2) {
			numOverlapped += 1
		}

	}

	readFile.Close()

	fmt.Println(numFullyContained)
	fmt.Println(numOverlapped)
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

func stringArrayToInt(assignment []string) []int {

	converted := make([]int, len(assignment))

	for i, el := range assignment {
		num, err := strconv.Atoi(el)
		if err != nil {
			panic(err)
		}
		converted[i] = num
	}

	return converted
}

func areFullyContained(elf1 []int, elf2 []int) bool {
	if ((elf1[0] <= elf2[0]) && (elf1[1] >= elf2[1])) || ((elf1[0] >= elf2[0]) && (elf1[1] <= elf2[1])) {
		return true
	}
	return false
}

func areOverlapped(elf1 []int, elf2 []int) bool {
	if ((elf1[0] <= elf2[0]) && (elf1[1] >= elf2[0])) || ((elf2[0] <= elf1[0]) && (elf2[1] >= elf1[0])) {
		return true
	}
	return false
}
