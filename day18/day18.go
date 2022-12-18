package main

import (
	"bufio"
	"fmt"
	"os"
)

type Side struct {
	x   int
	y   int
	z   int
	dir string
}

type Cube struct {
	x int
	y int
	z int
}

func main() {

	readFile, fileScanner := getScannedFile("input.txt")

	sides := make(map[Side]int)
	cubes := make([]Cube, 0)
	maxCube := Cube{}

	// build grid
	for fileScanner.Scan() {
		var x, y, z int
		fmt.Sscanf(fileScanner.Text(), "%d,%d,%d", &x, &y, &z)

		// for part 1 every cube is represented with 3 sides
		sides = addSides(sides, x, y, z)

		// unforunately part 1 representation is not very helpful for
		// part 2, so now we also save cubes
		cube := Cube{x: x, y: y, z: z}
		cubes = append(cubes, cube)
		maxCube = checkMaxCoord(maxCube, cube)
	}
	readFile.Close()

	fmt.Println(part1(sides))
	fmt.Println(part2(cubes, maxCube))
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

func addSides(sides map[Side]int, x int, y int, z int) map[Side]int {
	// front
	s := Side{x: x, y: y, z: z, dir: "x"}
	sides[s] += 1
	// back
	s = Side{x: x - 1, y: y, z: z, dir: "x"}
	sides[s] += 1
	// top
	s = Side{x: x, y: y, z: z, dir: "z"}
	sides[s] += 1
	// bottom
	s = Side{x: x, y: y, z: z - 1, dir: "z"}
	sides[s] += 1
	// right
	s = Side{x: x, y: y, z: z, dir: "y"}
	sides[s] += 1
	// left
	s = Side{x: x, y: y - 1, z: z, dir: "y"}
	sides[s] += 1

	return sides
}

func part1(sides map[Side]int) int {
	count := 0
	for _, v := range sides {
		if v == 1 {
			count += 1
		}
	}
	return count
}

func checkMaxCoord(maxCube Cube, cube Cube) Cube {
	if cube.x > maxCube.x {
		maxCube.x = cube.x
	}
	if cube.y > maxCube.y {
		maxCube.y = cube.y
	}
	if cube.z > maxCube.z {
		maxCube.z = cube.z
	}

	return maxCube
}

func part2(cubes []Cube, maxCube Cube) int {
	cubeMap := make(map[Cube]bool)
	for _, cube := range cubes {
		cubeMap[cube] = true
	}
	exterior := make(map[Cube]bool)
	count := getExterior(Cube{x: 0, y: 0, z: 0}, cubeMap, exterior, maxCube)
	return count
}

func getExterior(cube Cube, cubeMap map[Cube]bool, exterior map[Cube]bool, maxCube Cube) int {
	// outside the bounds
	if cube.x < -1 || cube.x > maxCube.x+1 ||
		cube.y < -1 || cube.y > maxCube.y+1 ||
		cube.z < -1 || cube.z > maxCube.z+1 {
		return 0
	}
	// already counted
	if exterior[cube] {
		return 0
	}
	// cube that we have not counted yet
	if cubeMap[cube] {
		return 1
	}

	exterior[cube] = true

	count := 0
	neighbors := getNeighbors(cube)
	for _, neighbor := range neighbors {
		count += getExterior(neighbor, cubeMap, exterior, maxCube)
	}
	return count
}

func getNeighbors(cube Cube) []Cube {
	return []Cube{
		// front
		{x: cube.x + 1, y: cube.y, z: cube.z},
		// back
		{x: cube.x - 1, y: cube.y, z: cube.z},
		// top
		{x: cube.x, y: cube.y, z: cube.z + 1},
		// bottom
		{x: cube.x, y: cube.y, z: cube.z - 1},
		// right
		{x: cube.x, y: cube.y + 1, z: cube.z},
		// left
		{x: cube.x, y: cube.y - 1, z: cube.z},
	}
}
