package node

import (
	"fmt"
	"io"
)

type Tree struct {
	Root *Node
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

func (t *Tree) Insert(data int) *Tree {
	if t.Root == nil {
		t.Root = NewNode(data, nil,  nil)
	} else {
		t.Root.insert(data)
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


// Print out nodes ---> same logic implementations for the two methods
func Print(w io.Writer, node *Node, ns int, ch rune) {
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
// this method does NOT seem to be working as expected
func PrintNode(n *Node){
	if n == nil {
		return
	}

	fmt.Println("Value: ", n.Value )

	PrintNode(n.Left)
	PrintNode(n.Right)
}


// generating a Tree from a file
func Load() []Node {
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
