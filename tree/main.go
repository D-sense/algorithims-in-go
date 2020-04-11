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

import "fmt"

func main(){

	tree := &Tree{Data: 20}
	tree.add(0)
	tree.add(40)
	tree.add(-15)

	tree.Children[0].add(12)
	tree.Children[0].add(-2)
	tree.Children[0].add(1)

	tree.Children[2].add(-2)

	fmt.Println("Breadth-First Traversal")
	tree.traverseBF(NodePrint)
	fmt.Println()

	fmt.Println("Depth-First Traversal")
	tree.traverseDF(NodePrint)
	fmt.Println()

	//tree.remove(4)
	//tree.traverseBF(NodePrint)

}


type Tree struct {
	Data int
	Children []*Tree
}

func NewTree(data int) *Tree {
	return &Tree{
		Data: data,
	}
}


func (n *Tree) add(data int) {
	if n.Children == nil {
		n.Children =  append(n.Children, NewTree(data))
		return
	}
	n.Children = append(n.Children, NewTree(data))
}

//not perfect; currently removed the parent only (does not look through their children)
func (n *Tree) remove(data int) {
	n.Children = n.Filter(data)
}


func NodePrint(node *Tree) {
	fmt.Printf("%v \n", node.Data)
}
func (n *Tree) Filter(data int) []*Tree {
	nArr := make([]*Tree, 0)

	for _, v := range n.Children {
		if data != v.Data  {
			nArr = append(nArr, v)
		}
	}

	return nArr
}


func (n *Tree) traverseBF(f func(*Tree)) {
	if n == nil{
		return
	}

	arr := []*Tree{n}

	for len(arr) > 0 {
		node := arr[0]
		arr = arr[1:]

		arr = append(arr, node.Children...)
		f(node)
	}
}


func (n *Tree) traverseDF(f func(*Tree)) {
	if n == nil{
		return
	}

	arr := []*Tree{n}

	for len(arr) > 0 {
		node := arr[0]
		arr = arr[1:]

		arr = append(node.Children, arr...)
		f(node)
	}
}