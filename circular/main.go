// --- Directions
// Return the 'middle' node of a linked list.
// If the list has an even number of elements, return
// the node at the end of the first half of the list.
// *Do not* use a counter variable, *do not* retrieve
// the size of the list, and only iterate
// through the list one time.
// --- Example
//   const l = new LinkedList();
//   l.insertLast('a')
//   l.insertLast('b')
//   l.insertLast('c')
//   midpoint(l); // returns { data: 'b' }

package main

import (
	"fmt"
	"github.com/d-sense/algorithims-in-go/circular/linkedlist"
)

func main(){
	link := linkedlist.NewLinkedList()
    last := link.CreateCircularLinkedList("CIRCULAR")

	fmt.Println(last.Next)

    // this will run into infinite loop because of this CreateCircularLinkedList()
	//each := link.ForEach(linkedlist.Capitalize)
	//fmt.Println(each)

	fmt.Println(circular(link))
}

func circular(l *linkedlist.LinkedList) bool {
	 slow := l.Head
	 fast := l.Head

	 for fast.Next != nil && fast.Next.Next != nil  {
		 slow = slow.Next
		 fast = fast.Next.Next

	 	if fast == slow {
	 		return true
		}
	 }

	 return false
}


