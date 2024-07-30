package linkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type SinglyLinkedList struct {
	Head *Node
}

func NewNode(data int) *Node {
	return &Node{Data: data, Next: nil}
}

func (list *SinglyLinkedList) AddNode(data int) {
	newNode := NewNode(data)
	if list.Head == nil {
		list.Head = newNode
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func (list *SinglyLinkedList) PrintList() {
	current := list.Head
	for current != nil {
		fmt.Print(current.Data, " -> ")
		current = current.Next
	}
	fmt.Print("nil")
}
