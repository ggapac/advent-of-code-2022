package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	readFile, fileScanner := getScannedFile("input.txt")

	// build grid
	var forest [][]int
	i := 0

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		forest = append(forest, make([]int, len(currentLine)))

		for pos, char := range currentLine {
			forest[i][pos] = int(char - '0')
		}

		i += 1
	}
	readFile.Close()

	// edge trees
	numVisible := 2*len(forest[0]) + 2*(len(forest)-2)

	maxScenicView := 0

	for i := 1; i < len(forest)-1; i++ {
		for j := 1; j < len(forest[i])-1; j++ {

			if isVisible(forest, i, j) {
				numVisible += 1
			}

			scenicView := getScenicScore(forest, i, j)
			if scenicView > maxScenicView {
				maxScenicView = scenicView
			}

		}
	}

	fmt.Println(numVisible)    // Part 1
	fmt.Println(maxScenicView) // Part 2
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

func isVisible(forest [][]int, i int, j int) bool {
	tree := forest[i][j]

	left := scan(forest, tree, i, j-1, 0, -1)
	if left {
		return true
	}

	right := scan(forest, tree, i, j+1, 0, 1)
	if right {
		return true
	}

	up := scan(forest, tree, i-1, j, -1, 0)
	if up {
		return true
	}

	down := scan(forest, tree, i+1, j, 1, 0)
	return down
}

func scan(forest [][]int, tree int, i int, j int, iAdd int, jAdd int) bool {

	for i >= 0 && i < len(forest) && j >= 0 && j < len(forest[0]) {
		if forest[i][j] >= tree {
			return false
		}
		i += iAdd
		j += jAdd
	}
	return true
}

func getScenicScore(forest [][]int, i int, j int) int {
	tree := forest[i][j]

	left := scanViewingDistance(forest, tree, i, j-1, 0, -1)
	right := scanViewingDistance(forest, tree, i, j+1, 0, 1)
	up := scanViewingDistance(forest, tree, i-1, j, -1, 0)
	down := scanViewingDistance(forest, tree, i+1, j, 1, 0)

	return left * right * up * down
}

func scanViewingDistance(forest [][]int, tree int, i int, j int, iAdd int, jAdd int) int {
	numTrees := 0
	for i >= 0 && i < len(forest) && j >= 0 && j < len(forest[0]) {
		numTrees += 1
		if forest[i][j] >= tree {
			break
		}
		i += iAdd
		j += jAdd
	}
	return numTrees
}
