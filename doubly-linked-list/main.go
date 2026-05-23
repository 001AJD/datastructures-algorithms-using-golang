package main

import "fmt"

func main() {
	fmt.Printf("\nDoubly Linked List")
	var dlist *DoublyLinkedList = NewDoublyLinkedList()

	var node1 *Node = NewNode(10)
	var node2 *Node = NewNode(20)
	var node3 *Node = NewNode(30)

	dlist.InsertNodeAtStart(node1)
	dlist.InsertNodeAtStart(node2)
	dlist.InsertNodeAtStart(node3)

	fmt.Printf("\nIs value 30 present in the linked list? :: %t", dlist.SearchByValue(30))
	fmt.Printf("\nIs value 10 present in the linked list? :: %t", dlist.SearchByValue(10))

	fmt.Printf("\nSearch and delete node with value 10")
	dlist.DeleteNode(10)
	dlist.Traverse()
	fmt.Printf("\nSearch and delete node with value 30")
	dlist.DeleteNode(30)
	dlist.Traverse()

	fmt.Printf("\nTravese the linked list")
	dlist.Traverse()

}
