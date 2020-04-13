// --- Directions
// Given a node, validate the binary search tree,
// ensuring that every node's left hand child is
// less than the parent node's value, and that
// every node's right hand child is greater than
// the parent

package main

import (
	"fmt"
	"github.com/d-sense/algorithims-in-go/validate/node"
)

func main(){
	// generating a tree from a file
	//nodes := node.Load()
	//for _, n := range nodes {
	//	node.PrintNode(&n)
	//}

	fmt.Println()

	// manually generating a tree
	t := &node.Tree{}
	//t.Insert(100).
	//	Insert(-20).
	//	Insert(-50).
	//	Insert(-15).
	//	Insert(-60).
	//	Insert(50).
	//	Insert(60).
	//	Insert(55).
	//	Insert(85).
	//	Insert(15).
	//	Insert(5).
	//	Insert(-10)

	// this should result to a VALID BST
	//t.Insert(5).Insert(15).Insert(0).Insert(20)

	// this should result to an INVALID BST
	t.Insert(10).Insert(5).Insert(15).Insert(0).Insert(20)
	t.Root.Left.Right = node.NewNode(999, nil, nil)


	// printing out nodes in the Tree
	node.PrintNode(t.Root)
	fmt.Println()

	// validate BST
	fmt.Println(validate(t.Root, 0, 0))
}


// validates Binary Search Tree [NOT accurate yet and requires more work]
func validate(node *node.Node, min, max int) bool {
	if max != 0 && node.Value > max {
		return false
	}

	if min != 0 && node.Value < min {
		return false
	}

	if node.Left != nil && !validate(node.Left, min, node.Value) {
		return false
	}

	if node.Right != nil  && !validate(node.Right, node.Value, max) {
		return false
	}

	return true
}





