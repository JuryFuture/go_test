package main

import "fmt"

type Node struct {
	elem interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func newLinkedList() *LinkedList {
	return &LinkedList{new(Node),0}
}

func (li *LinkedList) Add(index int, o interface{}) {
	if index <= li.size && index > -1 {
		node := li.head
		for j:=0; j<index; j++ {
			node = node.next
		}
		newNode := &Node{o, node.next}
		node.next = newNode
		li.size++
	} else {
		panic("out of index")
	}
}

func (li *LinkedList) Delete(index int) Node {
	if index <= li.Size && index > -1 {
		node := li.head
		for i:=0
	} else {
		panic("out of index")
	}
}

func main() {
	li := newLinkedList()
	li.Add(0,0)
	node := li.head
	for i:=0; i<li.size; i++ {
		node = node.next
		fmt.Println(node)
	}
}
