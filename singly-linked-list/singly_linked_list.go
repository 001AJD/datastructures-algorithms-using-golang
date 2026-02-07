package main

import "fmt"

type Node struct {
	Data uint8
	Next *Node
}

type LinkedList struct {
	Head *Node
	Size uint8
}

// NewLinkedList function create a new linked list with empty 0 nodes
// and returns a pointer to the linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Size: 0,
	}
}

// NewNode function creates a new node in the linked list and returns a pointer of type Node,
// value is the actual value that the linked list stores,
// next is the pointer to the next node in the linked list,
// if next is nil that means current node is the last node in the linked list
func NewNode(value uint8) *Node {
	return &Node{
		Data: value,
		Next: nil,
	}
}

// InsertNodeAtStart function inserts the node at the beginning of the linked list
func (l *LinkedList) InsertNodeAtStart(node *Node) {
	node.Next = l.Head // new node points to the current Head which will become the next node
	l.Head = node      // current node becomes the Head
	l.Size++           //Increase the size of the linked list
}

// TraverseLinkedList function travers the singly linked list and prints the data stored at each node,
// Traversal stops when the pointer to the next node is empty
func (l *LinkedList) TraverseLinkedList() {
	var current *Node = l.Head
	for current != nil {
		fmt.Printf("Data = %d, Mem Address = %p\n", current.Data, current)
		current = current.Next
	}
	fmt.Println("Reached end of linked list")

}
