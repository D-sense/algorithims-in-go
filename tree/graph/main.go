// --- Directions
// 1) Create a node class.  The constructor
// should accept an argument that gets assigned
// to the data property and initialize an
// empty array for storing children. The node
// class should have methods 'add' and 'remove'.
// 2) Create a tree class. The tree constructor
// should initialize a 'root' property to null.
// 3) Implement 'traverseBF' and 'traverseDF'
// on the tree class.  Each method should accept a
// function that gets called with each element in the tree


package main

import (
	"fmt"
	"io"
)

func main(){
	t := &Tree{}
	t.insert(100).
		insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10)

	//print(os.Stdout, tree.root, 0, 'M')
	printNode(t.root)
	fmt.Println()

	//t.search(50, 0)
	//
	//res := t.root.exist(-10)
	//fmt.Println(res)

	t.traverseBF(iterator)
}



type Tree struct {
	root *Node
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(data int, left, right *Node) *Node {
	return &Node{
		Value: data,
		Left:  left,
		Right: right,
	}
}


// inserting a node
func (t *Tree) insert(data int) *Tree {
	if t.root == nil {
		t.root = NewNode(data, nil,  nil)
	} else {
		t.root.insert(data)
	}
	return t
}
func (n *Node) insert(data int) {
	if data <= n.Value {
		if n.Left == nil {
			n.Left = NewNode(data, nil,  nil)
		} else {
			n.Left.insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = NewNode(data, nil,  nil)
		} else {
			n.Right.insert(data)
		}
	}
}


// Print out nodes ---> same logic implementations for the two methods
func print(w io.Writer, node *Node, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}

	fmt.Fprintf(w, "%c:%v\n", ch, node.Value)

	print(w, node.Left, ns+2, 'L')
	print(w, node.Right, ns+2, 'R')
}
func printNode(n *Node){
	if n == nil {
		return
	}

	fmt.Println("Value: ", n.Value )

	printNode(n.Left)
	printNode(n.Right)
}


// search for node
func (t *Tree) search(data, counter int) *Tree {
	if t.root == nil {
		fmt.Println("DOES NOT EXIST")
	} else {
		t.root.find(data, counter)
	}
	return t
}
func (n *Node) find(data, counter int){
	if data <= n.Value {
		// treat left node
		if n.Value == data {
			fmt.Println("FOUND: ", n.Value , "after ", counter, " counts")
		}else{
			counter++
			n.Left.find(data, counter)
		}
	}else{
		// treat right node
		if n.Value == data {
			fmt.Println("FOUND: ", n.Value , "after", counter, "counts")
		}else{
			counter++
			n.Right.find(data, counter)
		}
	}

}
//another solution
func (n *Node) exist(data int) bool{
	if n == nil {
		return false
	}

	if n.Value == data {
		return true
	}

	if data < n.Value{
		return n.Left.exist(data)
	}else{
		return n.Right.exist(data)
	}
}


// traversal
type iteratorFunc func ([]Node)
func iterator(nodes []Node){
	for _, v := range nodes {
		fmt.Println(v.Value)
	}
}
func getFirst(n []*Node) []*Node{
	result := n[:1]
	return result
}
//func (n *Tree) traverseBF(fn iteratorFunc) {
//	arr := []Node{*n.root}
//    counter := 0
//
//    fmt.Println("A: ", arr)
//
//	for len(arr) > counter {
//		arr := arr[counter:counter+1] // write a function to get first element here
//		node := arr
//		counter++
//
//		fmt.Println("B: ", node)
//
//		for _, v := range node {
//			arr = append(arr, v)
//		}
//
//		fmt.Println("C: ", arr)
//
//		fn(node)
//	}
//}

func (root *Tree) Bfs(f func(*Tree)) {
	if root == nil {
		return
	}
	queue := []*Tree{root}
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		f(t)
		queue = append(queue, t.Children...)
	}
}