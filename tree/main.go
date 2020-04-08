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
 // var list ArrBox
 // list = []string{"111", "222","3","4","544","6","7333","8","911","10"}
 //// list = append(list, "1", "2","3","4","5","6","7","8","9","10")
 // fmt.Println(Filter(list, length))

	var node Node
	node.add("HAMMED")
	node.add("ADESHINA")
	node.add("HASSAN")
	fmt.Println(node.children)

	fmt.Println( node.remove("ADESHINA"))
	fmt.Println(node.children)

}



type Node struct {
	data string
	children []Node
}

func NewNode(data string) Node {
	return Node{
		data: data,
	}
}

func (n *Node) add(data string) []Node {
	n.children = append(n.children, NewNode(data))
	return n.children
}

func (n *Node) remove(data string) []Node{
	nArr := make([]Node, 0)

	for _, v := range n.children{
		if v.data != data {
			nArr  = append(nArr, v)
		}
   }
   n.children = nArr

   return n.children
}



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