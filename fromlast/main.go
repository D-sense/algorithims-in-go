// --- Directions
// Given a linked list, return the element n spaces
// from the last node in the list.  Do not call the 'size'
// method of the linked list.  Assume that n will always
// be less than the length of the list.
// --- Examples
//    const list = new List();
//    list.insertLast('a');
//    list.insertLast('b');
//    list.insertLast('c');
//    list.insertLast('d');
//    fromLast(list, 2).data // 'b'

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

	fmt.Println(fromLast(link, 3))

	//each := link.ForEach(linkedlist.Capitalize)
	//fmt.Println(each)
}


// without using size() or/and counter
func fromLast(l *linkedlist.LinkedList, n int) *linkedlist.Node {
	 slow := l.GetFirst()  // or l.head
	 fast := l.GetFirst()  //or, l.head

	 for n > 0  {
	 	fast = fast.Next
	 	n--;
	 }

	for fast.Next != nil  {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}


