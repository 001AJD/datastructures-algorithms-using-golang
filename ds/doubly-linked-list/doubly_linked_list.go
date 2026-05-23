package main

import "fmt"

type Node struct {
	data int   // actual data
	next *Node // pointer to the next node
	prev *Node // pointer to the previous node
}

// Preserve the head to start the traversal
// it is a pointer to the first node in the linked list
// initially the next pointer will be nil, until the first node is added
// this is container for the list
type DoublyLinkedList struct {
	Head *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Head: nil,
	}
}

// creates a new node with data
// the next and prev pointer are null, it will be populated when called with associated linked list
func NewNode(data int) *Node {
	return &Node{
		data: data,
		prev: nil,
		next: nil,
	}
}

// inserts the node at the start of the linked list
func (l *DoublyLinkedList) InsertNodeAtStart(node *Node) {
	// if the head is nil, that means no node exists
	// if the head is not nil, that means the existing node needs to be rewired
	// since the nodes are inserted at the start, the new node's address needs to be stored in current node.prev
	if l.Head != nil {
		l.Head.prev = node
	}
	node.next = l.Head
	node.prev = nil
	l.Head = node // replace the current head with the new node
}

// function to traverse the linked list
func (l *DoublyLinkedList) Traverse() {
	fmt.Printf("\n**************** Traversing... ****************")
	// start from the head
	// loop through untill node.next != nil
	// currentNode = currentNode.Next in each iteration
	currentNode := l.Head

	for currentNode != nil {
		fmt.Printf("\nNode Memory Address :: %p, Data :: %d", currentNode, currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Printf("\n**************** Traversing completed... ****************")
}

// function to search the doubly linked list by value
func (l *DoublyLinkedList) SearchByValue(item int) bool {
	// visit each node
	// compare the data with the item to be searched
	// if data == item, item found, return true
	// if reached end of linked list, then return false.
	// what is end of linked list node.next = nil

	currentNode := l.Head

	for currentNode != nil {
		if item == currentNode.data {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// function to delete the node by value
func (l *DoublyLinkedList) DeleteNode(item int) bool {
	// traverse the linked list
	// on each iteration check if data == item
	// if item found, rewire the previous and next node
	// currentNode.prev.next = currentNode.next
	// currentNode.next.prev = currentNode.prev
	currentNode := l.Head

	for currentNode != nil {
		if currentNode.data == item {
			fmt.Printf("\nFound the item:: %d, deleting the node %p", currentNode.data, currentNode)
			if currentNode.prev == nil {
				// this is a head node, set head to currentNode.next
				l.Head = currentNode.next
				fmt.Printf("\nThe given item is present at head node")
				return true
			} else if currentNode.next == nil {
				// this is a tail node, set only previous node next as nil
				fmt.Printf("\nThe given item is present in tail node")
				currentNode.prev.next = nil
				return true
			} else {
				currentNode.next.prev = currentNode.prev
				currentNode.prev.next = currentNode.next
				return true
			}

		}
		currentNode = currentNode.next
	}
	fmt.Printf("\nList traversal ended, node with item %d not found", item)
	return false
}
