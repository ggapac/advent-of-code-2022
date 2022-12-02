package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scoreSum := 0

	readFile, fileScanner := getScannedFile("input.txt")

	for fileScanner.Scan() {
		game := strings.Fields(fileScanner.Text())

		//gameScore := getGameScore1(game[0], game[1]) // Strategy 1
		gameScore := getGameScore2(game[0], game[1]) // Strategy 2

		scoreSum += gameScore
	}

	readFile.Close()

	// Result
	fmt.Println(scoreSum)
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

func getGameScore1(opponent string, response string) int {
	responseVal := 0

	switch response {
	case "X":
		responseVal = 1
	case "Y":
		responseVal = 2
	case "Z":
		responseVal = 3
	}

	if opponent == "A" && responseVal == 2 || opponent == "B" && responseVal == 3 || opponent == "C" && responseVal == 1 {
		return responseVal + 6
	} else if opponent == "A" && responseVal == 1 || opponent == "B" && responseVal == 2 || opponent == "C" && responseVal == 3 {
		return responseVal + 3
	}

	return responseVal
}

func getGameScore2(opponent string, response string) int {
	opponentVal := 0

	switch opponent {
	case "A":
		opponentVal = 1
	case "B":
		opponentVal = 2
	case "C":
		opponentVal = 3
	}

	gameScore := 0

	switch response {
	case "X":
		gameScore = (opponentVal+1)%3 + 1
	case "Y":
		gameScore = opponentVal + 3 // + 3 for draw
	case "Z":
		gameScore = opponentVal%3 + 7 // + 1 for score calculation, + 6 for winning
	}

	return gameScore
}
