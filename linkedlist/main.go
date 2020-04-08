// --- Directions
// Implement classes Node and Linked Lists
// See 'directions' document

package main

import (
	"fmt"
	"strings"
)

func main(){
 // node := NewNode("Hi Shola", nil)
  linkedlist := NewLinkedList()

 // linkedlist.insertFirst(node)
  linkedlist.insertFirst("a")
  linkedlist.insertFirst("b")
  linkedlist.insertFirst("c")

 // linkedlist.removeFirst()
   //linkedlist.removeLast()
   //linkedlist.insertLast("YO!")
   //linkedlist.remoteAt(0)
  //fmt.Println(linkedlist.getAt(0))

  linkedlist.insertAt("NeW  Node", 2)


  fmt.Println(linkedlist.forEach(Capitalize))
  first := linkedlist.getFirst()
  fmt.Println(first)

  last := linkedlist.getLast()
  fmt.Println(last)
}


type Node struct {
	data interface{}
	next *Node
}

func NewNode(data string, next *Node) *Node {
	return &Node{
		data: data,
		next: next,
	}
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
	}
}

func (l * LinkedList) insertFirst(n string){
	l.head = NewNode(n, l.head)
}

func (l * LinkedList) size() int {
	counter := 0
	node := l.head

	for node != nil {
		counter++
		node = node.next
	}

	return counter
}

func (l * LinkedList) getFirst() *Node {
	return l.head
}

func (l * LinkedList) getLast() *Node {
	node := l.head
    last := node

	for node != nil {
		last = node
		node = node.next
	}

	return last
}

func (l * LinkedList) clear(){
	l.head = nil
}

func (l * LinkedList) removeFirst() *Node{
	if l.head == nil {
		return l.head
	}

	 l.head = l.head.next
	 return l.head
}

func (l * LinkedList) removeLast() *Node {
	if l.head == nil {
		return l.head
	}

	if l.head.next == nil {
		l.head = nil
	}

    prev := l.head
    node := l.head.next

    for node.next != nil {
    	prev = node
    	node = node.next
	}

	prev.next = nil

	return prev.next
}

func (l * LinkedList) insertLast(n string) {
	last := l.getLast()

	if last != nil{
		last.next = NewNode(n, nil)
	}else{
		l.head = NewNode(n, nil)
	}
}

func (l * LinkedList) getAt(index int) *Node {
	counter := 0
	node := l.head

	for node != nil {
       if counter == index {
       	return node
	   }

	   counter++
	   node = node.next
	}
	return nil
}

func (l * LinkedList) remoteAt(index int) *Node {
	if l.head == nil {
		return l.head
	}

	if index == 0{
		l.head = l.head.next
	}

   node := l.getAt(index)
   node.next = nil

   return l.head
}

func (l * LinkedList) insertAt(data string, index int) *Node {
	if l.head == nil {
		l.head = NewNode(data, nil)
		return l.head
	}

	if index == 0 {
		l.head = NewNode(data, l.head)
		return l.head
	}

	previous := l.getAt(index-1)
	if previous == nil {
		previous = l.getLast()
	}

	node := NewNode(data, previous.next)
	previous.next = node

   return l.head
}


type MoodFunc func(n *Node) Node
func (l * LinkedList) forEach(fn MoodFunc) []Node {
   node := l.head
   result := make([]Node, 0)

   for node.next != nil {
   	 result = append(result, fn(node))
   	 node = node.next
   }

   return result
}

func Capitalize(n *Node) Node {
	data := strings.ToUpper(fmt.Sprintf("%s", n.data))
	n.data = data
	return *n
}