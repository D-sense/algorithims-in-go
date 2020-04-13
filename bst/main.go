// --- Directions
// 1) Implement the Node class to create
// a binary search tree.  The constructor
// should initialize values 'data', 'left',
// and 'right'.
// 2) Implement the 'insert' method for the
// Node class.  Insert should accept an argument
// 'data', then create an insert a new node
// at the appropriate location in the tree.
// 3) Implement the 'contains' method for the Node
// class.  Contains should accept a 'data' argument
// and return the Node in the tree with the same value.

package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	// generating a tree from a file
	nodes := load()
	for _, node := range nodes {
		printNode(&node)
	}

    // manually generating a tree
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

	// printing out nodes in the Tree
	printNode(t.root)
	fmt.Println()

	// printing out nodes in the Tree
	print(os.Stdout, t.root, 0, 'M')

	//searching
	fmt.Println(t.contains(3309))

	//another searching
    fmt.Println(t.root.exist(-33))
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

func (t *Tree) insert(data int) *Tree {
	if t.root == nil {
		t.root = NewNode(data, nil,  nil)
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *Node) insert(data int) {
	if n == nil {
		return
	} else if data <= n.Value {
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

// search for node
func (t *Tree) contains(data int)  bool {
	if t.root == nil {
		return false
	} else {
		return t.root.find(data)
	}
}
func (n *Node) find(data int) bool {
	// checking if no further node in the Tree
	if n == nil {
		return false
	}

	if data < n.Value {
		// treat left node
		if n.Value == data {
			return true
		}else{
			return n.Left.find(data)
		}
	}else{
		// treat right node
		if n.Value == data {
			return true
		}else{
			return n.Right.find(data)
		}
	}
}
//another solution
func (n *Node) exist(data int) bool{
	// checking if no further node in the Tree
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


// generating a Tree from a file
func load() []Node {
	var N int
	fmt.Scanf("%d", &N)
	fmt.Println(N)

	var nodes = make([]Node, N)
	for i := 0; i < N; i++{
		var val, indexLeft, indexRight int
		fmt.Scanf("%d %d %d", &val, &indexLeft, &indexRight)

		nodes[i].Value = val
		if indexLeft >= 0 {
			nodes[i].Left = &nodes[indexLeft]
		}

		if indexRight >= 0 {
			nodes[i].Right = &nodes[indexRight]
		}
	}

	return nodes
}


