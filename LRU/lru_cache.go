package main

import "fmt"

type Domain struct {
	index  int
	domain string
	score  int
}

type Node struct {
	data Domain
	next *Node
	prev *Node
}

type LRU struct {
	capacity    int
	currentSize int
	hashMap     map[int]*Node
	head        *Node
	tail        *Node
}

func NewLRU(c int) *LRU {
	return &LRU{
		capacity:    c,
		currentSize: 0,
		hashMap:     make(map[int]*Node),
		head:        nil,
		tail:        nil,
	}
}

func NewNode(d Domain) *Node {
	return &Node{
		data: d,
		next: nil,
		prev: nil,
	}
}

// TODO: update the hashmap to be capacity bound, currently hashmap never evicts anything
func (l *LRU) Put(item Domain) {
	if l.currentSize == l.capacity {
		// if capacity is reached, remove the tail node
		// reduce the capacity by 1
		fmt.Printf("\n Cache capacity full, invoking eviction!!")
		fmt.Printf("\n Evicting tail node with data %v", l.tail.data)
		l.tail = l.tail.prev
		l.tail.next = nil
		l.currentSize--
		fmt.Printf("\n New tail node with data is %v", l.tail.data)
	}
	var newNode *Node = NewNode(item)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.currentSize++
		return
	}
	l.head.prev = newNode
	newNode.next = l.head
	l.currentSize++
	l.head = newNode
	l.hashMap[newNode.data.index] = newNode
}

func (l *LRU) Get(id int) (bool, Domain) {
	if l.hashMap[id] == nil {
		return false, Domain{}
	}
	var nodeAddress *Node = l.hashMap[id]
	return true, nodeAddress.data
}
