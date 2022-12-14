package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const rock = '#'
const sand = 'o'

func main() {

	readFile, fileScanner := getScannedFile("input.txt")

	grid := make(map[string]rune)
	maxY := 0

	// build grid
	for fileScanner.Scan() {
		points := strings.Split(fileScanner.Text(), " -> ")

		for i := 1; i < len(points); i++ {
			grid, maxY = getLineAndCheckY(points[i-1], points[i], grid, maxY)
		}

	}
	readFile.Close()

	// simulate sand
	x := 500
	y := 0
	numSand := 0
	part := 1 // set to 1 or 2
	floor := maxY + 2*(part-1)

	for y < floor {
		var newX int
		var newY int
		if part == 1 {
			newX, newY = nextSandPosition(x, y, grid, floor+1)
		} else {
			newX, newY = nextSandPosition(x, y, grid, floor)
		}

		// stopping criteria if we cannot spawn any more sand
		if grid["499,1"] == sand && grid["500,1"] == sand && grid["501,1"] == sand {
			numSand += 1
			break
		}

		if newX == x && newY == y {
			// rest sand
			key := fmt.Sprintf("%d,%d", x, y)
			grid[key] = sand
			numSand += 1
			// new sand
			x = 500
			y = 0
			continue
		}

		x = newX
		y = newY
	}

	fmt.Println(numSand)
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

func getLineAndCheckY(point1 string, point2 string, grid map[string]rune, maxY int) (map[string]rune, int) {
	coords1 := strings.Split(point1, ",")
	coords2 := strings.Split(point2, ",")

	x1, _ := strconv.Atoi(coords1[0])
	x2, _ := strconv.Atoi(coords2[0])

	y1, _ := strconv.Atoi(coords1[1])
	y2, _ := strconv.Atoi(coords2[1])

	if x1 < x2 {
		grid = drawLine(x1, x2, y1, grid, true)
	} else if x2 < x1 {
		grid = drawLine(x2, x1, y1, grid, true)
	} else if y1 < y2 {
		grid = drawLine(y1, y2, x1, grid, false)

		if y2 > maxY {
			maxY = y2
		}
	} else if y2 < y1 {
		grid = drawLine(y2, y1, x1, grid, false)

		if y1 > maxY {
			maxY = y1
		}
	}

	return grid, maxY
}

func drawLine(minCoord int, maxCoord int, staticCoord int, grid map[string]rune, horizontal bool) map[string]rune {
	for i := minCoord; i <= maxCoord; i++ {
		var key string
		if horizontal {
			key = fmt.Sprintf("%d,%d", i, staticCoord)
		} else {
			key = fmt.Sprintf("%d,%d", staticCoord, i)
		}
		grid[key] = rock
	}
	return grid
}

func nextSandPosition(x int, y int, grid map[string]rune, floor int) (int, int) {
	if isValid(x, y+1, grid, floor) { // down
		return x, y + 1
	} else if isValid(x-1, y+1, grid, floor) { // left down
		return x - 1, y + 1
	} else if isValid(x+1, y+1, grid, floor) { // right down
		return x + 1, y + 1
	}
	return x, y
}

func isValid(x int, y int, grid map[string]rune, floor int) bool {
	key := fmt.Sprintf("%d,%d", x, y)

	if grid[key] == rock || grid[key] == sand || y == floor {
		return false
	}
	return true
}
