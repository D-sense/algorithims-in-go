package main

import "fmt"

func main() {

	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	fmt.Println("first", l.First())
	fmt.Println("last", l.Last())
	fmt.Println("Size", l.Size())
	fmt.Println("forward traverse:")
	l.ForwardTraverse()
	fmt.Println("forward traverse:")
	l.BackwardTraverse()
	fmt.Println("form double linkedlist array:", l.formDoubleLinkedList())
	//	fmt.Println("find", l.Find(3))
	//	fmt.Println("next node", l.Find(3).NextNode())
	//	fmt.Println("previous node", l.Find(3).PrevNode())
}

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	value int
	Next  *Node
	Prev  *Node
}

func (l *List) Push(v int) {
	node := &Node{value: v}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.Next = node
		node.Prev = l.tail
	}

	l.tail = node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	if l.head == nil {
		return nil
	}

	node := l.head
	for node != nil {
		if node.Next == nil {
			return node
		}
		node = node.Next
	}

	return node
}

func (n *Node) NextNode() *Node {
	return n.Next
}

func (n *Node) PrevNode() *Node {
	return n.Prev
}

func (l *List) Find(v int) *Node {
	if l.head == nil {
		return nil
	}

	node := l.head
	for node != nil {
		if node.value == v {
			return node
		}
		node = node.Next
	}

	return nil
}

func (l *List) Size() int {
	if l.head == nil {
		return 0
	}
	counter := 0
	node := l.head
	for node != nil {
		counter++
		node = node.Next
	}

	return counter
}

func (l *List) ForwardTraverse() {
	if l.head == nil {
		return
	}

	node := l.First()
	for node != nil {
		fmt.Println(node.value)
		node = node.Next
	}
}

func (l *List) BackwardTraverse() {
	if l.head == nil {
		return
	}

	node := l.Last()
	for node != nil {
		fmt.Println(node.value)
		node = node.Prev
	}
}

func (l *List) formDoubleLinkedList() []int {
	if l.head == nil {
		return nil
	}

	result := make([]int, 0)

	node := l.First()
	for node != nil {
		result = append(result, node.value)
		node = node.Next
	}

	node = l.Last()
	for node != nil {
		result = append(result, node.value)
		node = node.Prev
	}
	return result
}
