package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Coords struct {
	x int
	y int
}

func main() {
	readFile, fileScanner := getScannedFile("input.txt")

	//ropeLen := 2 // Part 1
	ropeLen := 10 // Part 2

	rope := make([]Coords, ropeLen)
	visited := make(map[string]int)
	key0 := fmt.Sprintf("x=%dy=%d", 0, 0)
	visited[key0] += 1

	for fileScanner.Scan() {
		var direction string
		var numSteps int
		fmt.Sscanf(fileScanner.Text(), "%s %d", &direction, &numSteps)

		rope, visited = updatePosition(rope, direction, numSteps, visited)
	}
	readFile.Close()

	// Print result
	fmt.Println(len(visited))
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

func updatePosition(rope []Coords, dir string, n int, visited map[string]int) ([]Coords, map[string]int) {

	for i := 1; i <= n; i++ {
		rope[0] = updateHead(rope[0], dir, n)

		for j := 1; j < len(rope); j++ {
			rope[j] = updateTail(rope[j-1], rope[j])
		}

		key := fmt.Sprintf("x=%dy=%d", rope[len(rope)-1].x, rope[len(rope)-1].y)
		visited[key] += 1
	}

	return rope, visited
}

func updateHead(head Coords, dir string, n int) Coords {
	switch dir {
	case "U":
		head.y -= 1
	case "R":
		head.x += 1
	case "L":
		head.x -= 1
	case "D":
		head.y += 1
	}

	return head
}

func updateTail(head Coords, tail Coords) Coords {
	diff_x := math.Abs(float64(head.x - tail.x))
	diff_y := math.Abs(float64(head.y - tail.y))

	change_x := getChange(head.x, tail.x)
	change_y := getChange(head.y, tail.y)

	if diff_x+diff_y > 2 { // diagonal
		tail.x += change_x
		tail.y += change_y
	} else if diff_x > 1 { // horizontal
		tail.x += change_x
	} else if diff_y > 1 {
		tail.y += change_y // vertical
	}

	return tail
}

func getChange(h int, t int) int {
	if h < t {
		return -1
	} else if h > t {
		return 1
	}
	return 0
}
