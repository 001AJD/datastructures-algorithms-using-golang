package main

import "fmt"

func main() {
	fmt.Println("Create new linked List")
	var myLinkedList *LinkedList = NewLinkedList()
	var node1 *Node = NewNode(10)
	var node2 *Node = NewNode(20)
	var node3 *Node = NewNode(30)

	myLinkedList.InsertNodeAtStart(node1)
	myLinkedList.InsertNodeAtStart(node2)
	myLinkedList.InsertNodeAtStart(node3)
	fmt.Println(myLinkedList)
	myLinkedList.TraverseLinkedList()

}
