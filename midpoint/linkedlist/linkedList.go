package linkedlist

import (
	"fmt"
	"strings"
)

type Node struct {
	data interface{}
	Next *Node
}

func NewNode(data string, next *Node) *Node {
	return &Node{
		data: data,
		Next: next,
	}
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
	}
}

func (l *LinkedList) InsertFirst(n string){
	l.Head = NewNode(n, l.Head)
}

func (l *LinkedList) size() int {
	counter := 0
	node := l.Head

	for node != nil {
		counter++
		node = node.Next
	}

	return counter
}

func (l *LinkedList) GetFirst() *Node {
	return l.Head
}

func (l *LinkedList) getLast() *Node {
	node := l.Head
	last := node

	for node != nil {
		last = node
		node = node.Next
	}

	return last
}

func (l *LinkedList) clear(){
	l.Head = nil
}

func (l *LinkedList) removeFirst() *Node{
	if l.Head == nil {
		return l.Head
	}

	l.Head = l.Head.Next
	return l.Head
}

func (l *LinkedList) removeLast() *Node {
	if l.Head == nil {
		return l.Head
	}

	if l.Head.Next == nil {
		l.Head = nil
	}

	prev := l.Head
	node := l.Head.Next

	for node.Next != nil {
		prev = node
		node = node.Next
	}

	prev.Next = nil

	return prev.Next
}

func (l *LinkedList) insertLast(n string) {
	last := l.getLast()

	if last != nil{
		last.Next = NewNode(n, nil)
	}else{
		l.Head = NewNode(n, nil)
	}
}

func (l *LinkedList) getAt(index int) *Node {
	counter := 0
	node := l.Head

	for node != nil {
		if counter == index {
			return node
		}

		counter++
		node = node.Next
	}
	return nil
}

func (l *LinkedList) remoteAt(index int) *Node {
	if l.Head == nil {
		return l.Head
	}

	if index == 0{
		l.Head = l.Head.Next
	}

	node := l.getAt(index)
	node.Next = nil

	return l.Head
}

func (l *LinkedList) InsertAt(data string, index int) *Node {
	if l.Head == nil {
		l.Head = NewNode(data, nil)
		return l.Head
	}

	if index == 0 {
		l.Head = NewNode(data, l.Head)
		return l.Head
	}

	previous := l.getAt(index-1)
	if previous == nil {
		previous = l.getLast()
	}

	node := NewNode(data, previous.Next)
	previous.Next = node

	return l.Head
}


type MoodFunc func(n *Node) Node
func (l * LinkedList) ForEach(fn MoodFunc) []Node {
	node := l.Head
	result := make([]Node, 0)

	for node != nil {
		result = append(result, fn(node))
		node = node.Next
	}

	return result
}
func Capitalize(n *Node) Node {
	data := strings.ToUpper(fmt.Sprintf("%s", n.data))
	n.data = data
	return *n
}
