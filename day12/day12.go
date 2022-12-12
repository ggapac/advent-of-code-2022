package main

import (
	"bufio"
	"fmt"
	"os"
)

type Location struct {
	i         int
	j         int
	elevation rune
	dist      int
}

func main() {

	readFile, fileScanner := getScannedFile("input.txt")

	part1 := false
	var grid [][]rune
	var visited [][]bool
	i := 0
	S_i := make([]int, 0)
	S_j := make([]int, 0)
	var E_i, E_j int

	// build grid
	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		grid = append(grid, make([]rune, len(currentLine)))
		visited = append(visited, make([]bool, len(currentLine)))

		for pos, char := range currentLine {
			grid[i][pos] = char

			if string(char) == "S" || (!part1 && string(char) == "a") {
				S_i = append(S_i, i)
				S_j = append(S_j, pos)
				grid[i][pos] = 'a'
			} else if string(char) == "E" {
				E_i = i
				E_j = pos
				grid[i][pos] = 'z'
			}
		}

		i += 1
	}
	readFile.Close()

	result := make([]int, 0)
	for i := 0; i < len(S_i); i++ { // looping for part 2 is not really efficient, but ok

		// reset visited array
		visited = resetVisited(visited)

		// bfs
		dist := bfs(S_i[i], S_j[i], E_i, E_j, grid, visited)
		if dist != 0 {
			result = append(result, dist)
		}
	}

	fmt.Println(getMin(result))
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

func bfs(s_i int, s_j int, e_i int, e_j int, grid [][]rune, visited [][]bool) int {
	queue := make([]Location, 0)
	queue = append(queue, Location{elevation: 'a', i: s_i, j: s_j})
	visited[s_i][s_j] = true

	for len(queue) != 0 {
		// pop
		source := queue[0]
		queue = queue[1:]

		srcEl := source.elevation

		// Destination found;
		if source.i == e_i && source.j == e_j {
			return source.dist
		}

		// up
		if isValid(source.i-1, source.j, srcEl, grid, visited) {
			queue = append(queue, Location{elevation: grid[source.i-1][source.j], i: source.i - 1, j: source.j, dist: source.dist + 1})
			visited[source.i-1][source.j] = true
		}

		// down
		if isValid(source.i+1, source.j, srcEl, grid, visited) {
			queue = append(queue, Location{elevation: grid[source.i+1][source.j], i: source.i + 1, j: source.j, dist: source.dist + 1})
			visited[source.i+1][source.j] = true
		}

		// left
		if isValid(source.i, source.j-1, srcEl, grid, visited) {
			queue = append(queue, Location{elevation: grid[source.i][source.j-1], i: source.i, j: source.j - 1, dist: source.dist + 1})
			visited[source.i][source.j-1] = true
		}

		// right
		if isValid(source.i, source.j+1, srcEl, grid, visited) {
			queue = append(queue, Location{elevation: grid[source.i][source.j+1], i: source.i, j: source.j + 1, dist: source.dist + 1})
			visited[source.i][source.j+1] = true
		}

	}
	return 0
}

func isValid(i int, j int, srcEl rune, grid [][]rune, visited [][]bool) bool {
	if (i >= 0 && j >= 0) &&
		(i < len(grid) && j < len(grid[0])) &&
		(int(grid[i][j])-int(srcEl) <= 1) &&
		!visited[i][j] {
		return true
	}
	return false
}

func resetVisited(visited [][]bool) [][]bool {
	for j := 0; j < len(visited); j++ {
		for k := 0; k < len(visited[j]); k++ {
			visited[j][k] = false
		}
	}
	return visited
}

func getMin(distances []int) int {
	min := distances[0]

	for i := 1; i < len(distances); i++ {
		if distances[i] < min {
			min = distances[i]
		}
	}
	return min
}
