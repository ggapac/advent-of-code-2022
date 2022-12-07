package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	diskSpace := 70000000
	spaceNeeded := 30000000

	rootNode := Node{FirstChild: nil, NextSibling: nil, Size: 0, Name: "/", Type: "dir"}
	curNode := &rootNode // save a pointer to root node

	readFile, fileScanner := getScannedFile("input.txt")

	for fileScanner.Scan() {
		line := strings.Fields(fileScanner.Text())

		if line[0] == "$" { // command
			switch line[1] {
			case "cd":
				curNode = changeDir(&rootNode, curNode, line[2])
			case "ls":
				continue
			}

		} else { // ls list
			var size int
			var filetype string

			if line[0] == "dir" {
				filetype = "dir"
			} else {
				tmp, err := strconv.Atoi(line[0])
				if err != nil {
					panic(err)
				}
				size = tmp
				filetype = "file"
			}
			curNode = insertChild(curNode, &rootNode, size, line[1], filetype)
		}

	}

	readFile.Close()

	fmt.Println(getDirSum(&rootNode, 100000))

	diff := rootNode.Size - (diskSpace - spaceNeeded)
	if diff > 0 { // do we need to free up?
		suitableDirs := getSuitableDirs(&rootNode, diff)
		fmt.Println(getSmallestSuitableDir(suitableDirs))
	}
}
