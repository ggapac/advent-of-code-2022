package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	FirstChild  *Node
	NextSibling *Node
	Size        int
	Name        string
	Type        string
	ParentNames []string
}

type Dir struct {
	Name string
	Size int
}

// Read file
func getScannedFile(filename string) (*os.File, *bufio.Scanner) {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	return readFile, fileScanner
}

// Commands: change dir
func changeDir(rootNode *Node, n *Node, command string) *Node {
	switch command {
	case "..":
		n = returnNode(rootNode, n.ParentNames)
	case "/":
		n = rootNode
	default:
		n = returnChild(n, command)
	}

	return n
}

// Node: insert
func insertChild(t *Node, rootNode *Node, size int, name string, filetype string) *Node {

	// loop through children, order from smallest to largest
	child := t.FirstChild
	parents := append(t.ParentNames, t.Name)
	newChild := Node{FirstChild: nil, NextSibling: nil, Size: size, Name: name, Type: filetype, ParentNames: parents}
	if child == nil {
		t.FirstChild = &newChild
	} else {

		// check if it already exists
		exists := false
		for child != nil {
			if child.Name == name && child.Type == filetype {
				exists = true
				break
			}
			child = child.NextSibling
		}

		if !exists {
			tmp := t.FirstChild
			t.FirstChild = &newChild
			newChild.NextSibling = tmp
		}

	}

	if size != 0 {
		t = updateSize(rootNode, newChild.ParentNames, newChild.Size)
	}

	return t
}

// Node: update size
func updateSize(rootNode *Node, parentNames []string, size int) *Node {
	curNode := rootNode
	for i := 0; i < len(parentNames); i++ {

		for curNode.Name != parentNames[i] {
			curNode = curNode.NextSibling
		}

		curNode.Size += size

		// we have to return parent
		if i == len(parentNames)-1 {
			break
		}
		curNode = curNode.FirstChild
	}
	return curNode
}

// Node: return
func returnNode(rootNode *Node, parentNames []string) *Node {
	curNode := rootNode
	for i := 0; i < len(parentNames); i++ {

		for curNode.Name != parentNames[i] {
			curNode = curNode.NextSibling
		}

		// we have to return parent
		if i == len(parentNames)-1 {
			break
		}
		curNode = curNode.FirstChild
	}
	return curNode
}

func returnChild(t *Node, name string) *Node {
	child := t.FirstChild
	for child.Name != name {
		child = child.NextSibling
	}
	return child
}

// Part 1 solution
func getDirSum(node *Node, limit int) int {
	if node == nil || node.Type != "dir" {
		return 0
	}

	var size int
	if node.Size > limit {
		size = 0
	} else {
		size = node.Size
	}
	curNode := node.FirstChild
	for curNode != nil {
		size += getDirSum(curNode, limit)
		curNode = curNode.NextSibling
	}
	return size
}

// Part 2 solution
func getSuitableDirs(node *Node, diff int) []Dir {
	if node == nil || node.Type != "dir" {
		return nil
	}

	var dirs []Dir
	if node.Size >= diff {
		dirs = append(dirs, Dir{Name: node.Name, Size: node.Size})
	}
	curNode := node.FirstChild
	for curNode != nil {
		dirs = append(dirs, getSuitableDirs(curNode, diff)...)
		curNode = curNode.NextSibling
	}
	return dirs
}

func getSmallestSuitableDir(dirs []Dir) int {
	minSize := dirs[0].Size

	for i := 0; i < len(dirs); i++ {
		if dirs[i].Size < minSize {
			minSize = dirs[i].Size
		}
	}
	return minSize
}
