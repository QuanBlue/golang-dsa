package linked_list

import "fmt"

type dNode struct {
	data interface{}
	prev *dNode
	next *dNode
}

type DoublyLinkedLists struct {
	head *dNode
}

// -------------------------------------------------------------- //
// Create New Double Node
// > Create a Single Linked List with initial data
// -------------------------------------------------------------- //
func createNewDoublyNode(data interface{}) *dNode {
	new_node := &dNode{
		data: data,
		prev: nil,
		next: nil,
	}

	return new_node
}

// -------------------------------------------------------------- //
// Get Tail
// > Get tail of Single Linked List
// -------------------------------------------------------------- //
func (ll *DoublyLinkedLists) getTail() *dNode {
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
func (ll *DoublyLinkedLists) print() {
	// Print forward
	fmt.Print("Print (forward) : ")
	
	node := ll.head

	for node != nil {
		if node.next != nil {
			fmt.Print(node.data," <-> ")
		} else {
			fmt.Println(node.data)
		}
		node = node.next
	}

	// Print backward
	fmt.Print("Print (backward): ")
	
	node = ll.getTail()

	for node != nil {
		if node.prev != nil {
			fmt.Print(node.data," <-> ")
		} else {
			fmt.Print(node.data,"\n\n")
		}
		node = node.prev
	}
}

// -------------------------------------------------------------- //
// Append
// > Append node at the end
// -------------------------------------------------------------- //
func (ll *DoublyLinkedLists) append(data interface{}) {
	new_node := createNewDoublyNode(data)

	//case empty linked list
	if ll.head == nil {
		ll.head = new_node	
		return
	}

	// Insert at the end
	ex_tail := ll.getTail()
	new_tail := new_node
	new_tail.prev = ex_tail
	ex_tail.next = new_tail
}

// -------------------------------------------------------------- //
// Pop
// > Remove first node
// -------------------------------------------------------------- //
func (ll *DoublyLinkedLists) pop() {
	if ll.head == nil {
		return
	}

	ll.head = ll.head.next 
	ll.head.prev = nil
}

// -------------------------------------------------------------- //
// Remove
// > Remove all node have value
// -------------------------------------------------------------- //
func (ll *DoublyLinkedLists) remove(data interface{}) {
	// case empty linked list
	if ll.head == nil {
		return 
	}

	// case head node
	if ll.head.data == data {
		ll.pop()
	}
	
	// case tail node
	tail := ll.getTail()
	if tail.data == data {
		prev_node := tail.prev
		prev_node.next = nil
	}

	// other node
	node := ll.head
	for node.next != nil {
		if node.data == data {
			prev_node := node.prev
			next_node := node.next

			prev_node.next = next_node
			next_node.prev = prev_node
		}

		node = node.next
	}
}