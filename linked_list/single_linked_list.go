package linked_list

import "fmt"

type Node struct {
	data interface{} 
	next *Node
}

type SingleLinkedLists struct {
	head *Node
}

// Create a Single Linked List with initial data
func createNewNode(data interface{}) *Node {
	new_node := &Node{
		data: data,
		next: nil,
	}

	return new_node
}

// Get tail of Single Linked List
func (ll *SingleLinkedLists) getTail() *Node {
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

// Print all element in Single Linked List
func (ll *SingleLinkedLists) print() {
	fmt.Print("Linked list: ")
	
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

// Append node at the end
func (ll *SingleLinkedLists) append(data interface{}) {
	new_node := createNewNode(data)

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

// remove first node
func (ll *SingleLinkedLists) pop() {
	if ll.head == nil {
		return
	}

	ll.head = ll.head.next 
}


// remove all node have value
func (ll *SingleLinkedLists) remove(data interface{}) {
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

