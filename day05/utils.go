package main

// Borrowed the linked list implementation from: https://blog.devgenius.io/linked-list-in-go-c663eb684291

type Node struct {
	data rune
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) insertAtHead(data rune) {
	temp1 := &Node{data, nil}
	if l.head == nil {
		l.head = temp1
	} else {
		temp2 := l.head
		l.head = temp1
		temp1.next = temp2
	}
	l.length += 1
}

func (l *LinkedList) deleteAtHead() {
	temp := l.head
	l.head = temp.next
	l.length -= 1
}

func (l *LinkedList) Reverse() {
	var curr, prev, next *Node
	curr = l.head
	prev = nil
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}
