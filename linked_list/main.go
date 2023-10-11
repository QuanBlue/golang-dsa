package linked_list

import "fmt"


func Main() {
	ll := new(SingleLinkedLists)
	fmt.Println("ll:",&ll,"- ll.head:",ll.head)

	ll.append(1)
	ll.append(2)
	ll.append(4)
	ll.append(4)
	ll.append(3)
	ll.append(4)
	ll.append(5)
	ll.print()

	ll.pop()
	ll.print()

	ll.remove(4)
	ll.print()

	ll.remove(5)
	ll.print()
}