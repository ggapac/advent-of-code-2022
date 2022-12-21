package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Number struct {
	value int
	index int
}

func main() {

	readFile, fileScanner := getScannedFile("input.txt")

	file := make([]Number, 0)

	part1 := false
	var mult, rep int

	if part1 {
		mult = 1
		rep = 1
	} else {
		mult = 811589153
		rep = 10
	}

	// read the file
	i := 0
	for fileScanner.Scan() {
		num, _ := strconv.Atoi(fileScanner.Text())
		file = append(file, Number{value: num * mult, index: i})
		i += 1
	}
	readFile.Close()

	// update the indices
	for r := 0; r < rep; r++ {
		for i := 0; i < len(file); i++ {
			oldIx := file[i].index
			if file[i].value == 0 {
				continue
			}
			newIx := getNewIx(oldIx, file[i].value, len(file)-1)

			file[i].index = newIx
			file = updateIx(file, i, oldIx, newIx)
		}
	}

	// find zero position and map final values
	zeroPosition := 0
	endFile := make(map[int]int)
	for i := 0; i < len(file); i++ {
		if file[i].value == 0 {
			zeroPosition = file[i].index
		}
		endFile[file[i].index] = file[i].value
	}

	pos1 := getNewIx(zeroPosition, 1000%len(file), len(file))
	pos2 := getNewIx(zeroPosition, 2000%len(file), len(file))
	pos3 := getNewIx(zeroPosition, 3000%len(file), len(file))

	result := endFile[pos1] + endFile[pos2] + endFile[pos3]
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

func getNewIx(curIx int, val int, length int) int {
	newIx := curIx + val

	if newIx <= 0 {
		newIx = (length + newIx%length)
	} else if newIx >= length {
		newIx = (newIx % length)
	}

	return newIx
}

func updateIx(file []Number, i int, oldIx int, newIx int) []Number {
	for j := 0; j < len(file); j++ {
		if i == j {
			continue
		}
		if oldIx < newIx && file[j].index <= newIx && file[j].index > oldIx {
			file[j].index = getUpdatedIx(file[j].index, -1, len(file))
		} else if oldIx > newIx && file[j].index >= newIx && file[j].index < oldIx {
			file[j].index = getUpdatedIx(file[j].index, 1, len(file))
		}
	}
	return file
}

func getUpdatedIx(curIx int, val int, length int) int {
	newIx := curIx + val

	if newIx < 0 {
		newIx = length - 1
	} else if newIx == length {
		newIx = 0
	}

	return newIx
}
