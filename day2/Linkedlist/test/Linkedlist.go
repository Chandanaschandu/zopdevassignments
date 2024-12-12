package test

import "fmt"

type Node struct {
	value int
	next  *Node
}

func NewNode(value int, next *Node) *Node {
	var n Node
	n.value = value
	n.next = next
	return &n
}
func TraverseLinkedList(head *Node) {
	temp := head
	for temp != nil {
		fmt.Printf("%d ", temp.value)
		temp = temp.next
	}
	fmt.Println()
}
func AddNodeAtEnd(head *Node, data int) *Node {
	if head == nil {
		head = NewNode(data, nil)
		return head
	}
	temp := head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = NewNode(data, nil)
	return head
}

func DeleteLastNode(head *Node) *Node {
	if head == nil {
		return head
	}
	temp := head
	for temp.next.next != nil {
		temp = temp.next
	}
	temp.next = nil
	return head
}
