package linked_list

import "fmt"

type sNode struct {
	data interface{} 
	next *sNode
}

type SinglyLinkedLists struct {
	head *sNode
}

// -------------------------------------------------------------- //
// Create New Single Node
// > Create a Single Linked List with initial data
// -------------------------------------------------------------- //
func createNewSinglyNode(data interface{}) *sNode {
	new_node := &sNode{
		data: data,
		next: nil,
	}

	return new_node
}

// -------------------------------------------------------------- //
// Get Tail
// > Get tail of Single Linked List
// -------------------------------------------------------------- //
func (ll *SinglyLinkedLists) getTail() *sNode {
	node := ll.head

	for node.next != nil && node.next.next != nil{
		node = node.next
	}

	if node.next != nil {
		return node.next
	} else {
		return node
	}
}

// -------------------------------------------------------------- //
// Print
// > Print all element in Single Linked List
// -------------------------------------------------------------- //
func (ll *SinglyLinkedLists) print() {
	fmt.Print("Print: ")
	
	node := ll.head

	for node != nil {
		if node.next != nil {
			fmt.Print(node.data," -> ")
		} else {
			fmt.Println(node.data)
		}
		node = node.next
	}
}

// -------------------------------------------------------------- //
// Append
// > Append node at the end
// -------------------------------------------------------------- //
func (ll *SinglyLinkedLists) append(data interface{}) {
	new_node := createNewSinglyNode(data)

	//case empty linked list
	if ll.head == nil {
		ll.head = new_node	
		return
	}

	// Insert at the end
	ex_tail := ll.getTail()
	new_tail := new_node
	ex_tail.next = new_tail
}

// -------------------------------------------------------------- //
// Pop
// > Remove first node
// -------------------------------------------------------------- //
func (ll *SinglyLinkedLists) pop() {
	if ll.head == nil {
		return
	}

	ll.head = ll.head.next 
}

// -------------------------------------------------------------- //
// Remove
// > Remove all node have value
// -------------------------------------------------------------- //
func (ll *SinglyLinkedLists) remove(data interface{}) {
	// case empty linked list
	if ll.head == nil {
		return 
	}

	// case head node
	if ll.head.data == data {
		ll.pop()
	}
	
	// other node
	node := ll.head
	for node.next != nil {
		if node.next.data == data {
			node.next = node.next.next
			continue
		}

		node = node.next
	}
}

