package main

import "fmt"

func main() {
	fmt.Println("Create new linked List")
	var myLinkedList *LinkedList = NewLinkedList()
	var node1 *Node = NewNode(10, nil)
	myLinkedList.InsertNodeAtStart(node1)

	fmt.Println(myLinkedList)
	myLinkedList.TraverseLinkedList()

}
