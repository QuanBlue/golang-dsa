package linked_list

import "fmt"

type LinkedLists interface {
	getTail() interface{}
	print()
	append(interface{})
	pop()
	remove(interface{})
}

func Main() {
	fmt.Println("-------------------------")
	fmt.Println("-  SINGLE LINKED LISTS  -")
	fmt.Println("-------------------------")

	s_ll := new(SingleLinkedLists)
	fmt.Println("s_ll:",&s_ll,"- s_ll.head:",s_ll.head)

	s_ll.append(1)
	s_ll.append(4)
	s_ll.append(2)
	s_ll.append(4)
	s_ll.append(3)
	s_ll.append(4)
	s_ll.append(5)
	s_ll.print()

	s_ll.pop()
	s_ll.print()

	s_ll.remove(4)
	s_ll.print()

	s_ll.remove(5)
	s_ll.print()

	fmt.Println("\n-------------------------")
	fmt.Println("-  DOUBLE LINKED LISTS  -")
	fmt.Println("-------------------------")

	d_ll := new(DoubleLinkedLists)
	fmt.Println("d_ll:",&d_ll,"- d_ll.head:",d_ll.head)

	d_ll.append(1)
	d_ll.append(2)
	d_ll.append(3)
	d_ll.append(4)
	d_ll.append(4)
	d_ll.append(5)
	d_ll.append(4)
	d_ll.print()

	d_ll.pop()
	d_ll.print()

	d_ll.remove(4)
	d_ll.print()

	d_ll.remove(5)
	d_ll.print()
}