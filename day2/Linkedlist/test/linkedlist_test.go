package test

import (
	"testing"
)

// Helper function to convert the linked list to a slice
func listToSlice(head *Node) []int {
	var result []int
	temp := head
	for temp != nil {
		result = append(result, temp.value)
		temp = temp.next
	}
	return result
}

func TestAddNodeAtEnd(t *testing.T) {
	var head *Node
	tests := []struct {
		input  int
		output []int
	}{
		{3, []int{3}},    // Add a node with value 3
		{4, []int{3, 4}}, // Add a node with value 4 after a node with value 3
	}

	for _, test := range tests {

		head = AddNodeAtEnd(head, test.input)

		got := listToSlice(head) //converting list to slice

		if !compareSlices(got, test.output) {
			t.Errorf("your result %v and expected output is %v", got, test.output)
		}
	}
}

// Helper function to compare two slices
func compareSlices(res, expected []int) bool {
	if len(res) != len(expected) {
		return false
	}
	for i := range res {
		if res[i] != expected[i] {
			return false
		}
	}
	return true
}
func TestDeleteLastNode(t *testing.T) {
	var head *Node
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{4}, []int{}},
	}
	for _, test := range tests {

		head = DeleteLastNode(head)
		res := listToSlice(head)
		if !compareSlices(res, test.output) {
			t.Errorf("For input %v, after deleting the last node, got %v; want %v", test.input, res, test.output)
		}
	}
}
