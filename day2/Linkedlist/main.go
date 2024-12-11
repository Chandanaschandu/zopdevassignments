package main

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
func main() {
	head := NewNode(10, NewNode(20, NewNode(30, NewNode(40, NewNode(50, nil)))))
	fmt.Printf("Input Linked list is: ")
	TraverseLinkedList(head)
	AddNodeAtEnd(head, 8)
	fmt.Printf("After adding node at end, linked list is: ")
	TraverseLinkedList(head)
	//fmt.Println("The linked list after removal of element is:")
	head = DeleteLastNode(head)
	/*for current := head; current != nil; current = current.next {
		fmt.Println(current.value) //print the modified linked list
	}*/
	fmt.Printf("After deleting last node of the linked list: ")
	TraverseLinkedList(head)
}
