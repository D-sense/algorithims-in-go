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
	"os"
)

func main(){
	//nodes := read()
	//for _, node := range nodes {
	//	printNode(&node)
	//}

	tree := &Tree{}
	tree.insert(100).
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

	print(os.Stdout, tree.root, 0, 'M')

	tree.insert(100)

	print(os.Stdout, tree.root, 0, 'M')
}

func read() []Node {
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

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	root *Node
}

func NewNode(data int, left, right *Node) *Node {
	return &Node{
		Value: data,
		Left:  left,
		Right: right,
	}
}


func printNode(n *Node)  {
	fmt.Println("Value: ", n.Value )

	if n.Left != nil {
		fmt.Println(" Left: ", n.Left.Value)
	}else{
		fmt.Println(" Left: nil")
	}

	if n.Right != nil {
		fmt.Println(" Right: ", n.Right.Value)
	}else{
		fmt.Println(" Right: nil")
	}

	fmt.Println()
}

//func (n *Node) add(data int) []Node {
//	n. = append(n.children, NewNode(data, nil, nil))
//	return n.children
//}

func insertRecursively(root *Node, v int) *Node {
	if root == nil {
		root = NewNode(v, nil, nil)
	} else if v < root.Value {
		root.Left = insertRecursively(root.Left, v)
	} else {
		root.Right = insertRecursively(root.Right, v)
	}

	return root
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

//func (n *Node) remove(data string) []Node {
//	nArr := make([]Node, 0)
//
//	for _, v := range n.children {
//		if v.data != data {
//			nArr = append(nArr, v)
//		}
//	}
//
//	return nil
//}
//
func (n *Tree) traverseBF() {
	fmt.Println(n.root)
	node := n.root

	for node. > 0{
    node := arr[counter]

    for _, v := range node.children {
    	 arr = append(arr, v)
	   }
	   fmt.Println(node)
	}
}


//
//func (n *Tree) traverseDF() {
//	fmt.Println(n.root)
//	arr := n.root.children
//	counter := 0
//
//	for len(arr) > 0{
//		node := arr[counter]
//
//		for _, v := range node.children {
//			arr = append(arr, v)
//		}
//		fmt.Println(node)
//	}
//}












//func (n *Node) Filter(data []string, check modFunc) []string {
//	nArr := make([]string, 0)
//
//	for _, v := range data {
//		if !check(v) {
//			nArr  = append(nArr, v)
//		}
//	}
//
//	return nArr
//}

//type ArrBox []string
//
//type modFunc func(s string) bool
//
//func length(n string)  bool {
//	if len(n) >= 3 {
//		return true
//	}
//	return false
//}