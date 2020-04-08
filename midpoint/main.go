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
	"github.com/d-sense/algorithims-in-go/midpoint/linkedlist"
)

func main(){
	link := linkedlist.NewLinkedList()

	link.InsertFirst("1")
	link.InsertFirst("2")
	link.InsertFirst("3")
	link.InsertFirst("4")
	link.InsertFirst("5")
	link.InsertFirst("6")
	link.InsertFirst("7")
	link.InsertFirst("8")
	link.InsertFirst("9")
	link.InsertFirst("10")

	fmt.Println("MIDPOINT: ", midPoint(link))

	each := link.ForEach(linkedlist.Capitalize)
	fmt.Println(each)
}


// without using size() or/and counter
func midPoint(l *linkedlist.LinkedList) *linkedlist.Node {
	 slow := l.Head
	 fast := l.Head

	 for fast.Next != nil && fast.Next.Next != nil  {
	 		slow = slow.Next
	 		fast = fast.Next.Next
	 }

	 return slow
}


