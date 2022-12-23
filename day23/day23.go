package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	x int
	y int
}

type Elf struct {
	proposed Point
	location Point
}

type Rectangle struct {
	n int
	s int
	w int
	e int
}

var compass = []string{"N", "S", "W", "E", "N"}

func main() {

	readFile, fileScanner := getScannedFile("input.txt")

	y := 0
	grid := make(map[Point]rune)
	elves := make([]Elf, 0)

	// read input
	for fileScanner.Scan() {
		line := fileScanner.Text()

		for x, r := range line {
			grid[Point{x: x, y: y}] = r

			if string(r) == "#" {
				elves = append(elves, Elf{location: Point{x: x, y: y}})
			}
		}
		y += 1

	}
	readFile.Close()

	dir := "N"
	minRectangle := Rectangle{e: -100000, w: 100000, s: -100000, n: 100000}
	elfMoved := true
	part1 := true
	r := 0

	for condition(part1, r, elfMoved) {
		proposed := make(map[Point]int)
		elfMoved = false

		// propose moves
		for i := 0; i < len(elves); i++ {
			// check if elf should move
			if shouldMove(elves[i].location, grid) {
				pp := proposeMove(elves[i].location, dir, grid)
				elves[i].proposed = pp
				proposed[pp] += 1
			}
		}

		// see if you can move
		for i := 0; i < len(elves); i++ {
			pl := elves[i].location
			pp := elves[i].proposed

			if proposed[pp] == 1 && pl != pp {
				elfMoved = true
				elves[i].location = pp
				grid[pl] = '.'
				grid[pp] = '#'
			} else {
				elves[i].proposed = pl
			}
		}

		dir = changeDir(dir)
		r++
	}

	// update min rectangle
	for i := 0; i < len(elves); i++ {
		minRectangle = updateRectangle(minRectangle, elves[i].location)
	}

	fmt.Println(getResult(part1, minRectangle, elves, r))
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

func proposeMove(p Point, dir string, grid map[Point]rune) Point {
	dirs := make([]string, 4)
	dirs[0] = dir
	dirs[1] = changeDir(dirs[0])
	dirs[2] = changeDir(dirs[1])
	dirs[3] = changeDir(dirs[2])

	for _, d := range dirs {
		valid := checkMove(p, d, grid)
		if valid {
			return getMove(p, d)
		}
	}

	return p
}

func getMove(p Point, dir string) Point {
	switch dir {
	case "N":
		p.y -= 1
	case "W":
		p.x -= 1
	case "S":
		p.y += 1
	case "E":
		p.x += 1
	}
	return p
}

func shouldMove(p Point, grid map[Point]rune) bool {
	n := checkMove(p, "N", grid)
	w := checkMove(p, "W", grid)
	e := checkMove(p, "E", grid)
	s := checkMove(p, "S", grid)

	if (n && w && e && s) || (!n && !w && !e && !s) {
		return false
	}
	return true
}

func checkMove(p Point, dir string, grid map[Point]rune) bool {
	p1 := false
	p2 := false
	p3 := false

	switch dir {
	case "N":
		// north west
		p1 = checkIfFree(p.x-1, p.y-1, grid)
		// north
		p2 = checkIfFree(p.x, p.y-1, grid)
		// north east
		p3 = checkIfFree(p.x+1, p.y-1, grid)
	case "E":
		// north east
		p1 = checkIfFree(p.x+1, p.y-1, grid)
		// east
		p2 = checkIfFree(p.x+1, p.y, grid)
		// south east
		p3 = checkIfFree(p.x+1, p.y+1, grid)
	case "S":
		// south east
		p1 = checkIfFree(p.x+1, p.y+1, grid)
		// south
		p2 = checkIfFree(p.x, p.y+1, grid)
		// south west
		p3 = checkIfFree(p.x-1, p.y+1, grid)
	case "W":
		// south west
		p1 = checkIfFree(p.x-1, p.y+1, grid)
		// west
		p2 = checkIfFree(p.x-1, p.y, grid)
		// north west
		p3 = checkIfFree(p.x-1, p.y-1, grid)
	}

	if p1 && p2 && p3 {
		return true
	}

	return false
}

func checkIfFree(x int, y int, grid map[Point]rune) bool {
	if val, ok := grid[Point{x: x, y: y}]; ok {
		if string(val) == "#" {
			return false
		}
	}
	return true
}

func changeDir(dir string) string {
	for i := 0; i < len(compass)-1; i++ {
		if compass[i] == dir {
			return compass[i+1]
		}
	}
	return ""
}

func updateRectangle(r Rectangle, p Point) Rectangle {
	if r.n > p.y {
		r.n = p.y
	}
	if r.s < p.y {
		r.s = p.y
	}
	if r.e < p.x {
		r.e = p.x
	}
	if r.w > p.x {
		r.w = p.x
	}
	return r
}

func condition(part1 bool, r int, elfMoved bool) bool {
	if part1 {
		if r == 10 {
			return false
		}
	} else {
		return elfMoved
	}

	return true
}

func part1solution(r Rectangle, elves []Elf) int {
	a := math.Abs(float64(r.n) - float64(r.s))
	b := math.Abs(float64(r.w) - float64(r.e))

	if r.n < 0 && r.s > 0 {
		a += 1
	}
	if r.w < 0 && r.e > 0 {
		b += 1
	}
	return int(a*b) - len(elves)
}

func getResult(part1 bool, minRectangle Rectangle, elves []Elf, r int) int {
	if part1 {
		return part1solution(minRectangle, elves)
	} else {
		return r
	}
}
