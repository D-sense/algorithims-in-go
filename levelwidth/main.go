// --- Directions
// Given the root node of a tree, return
// an array where each element is the width
// of the tree at each level.
// --- Example
// Given:
//     0
//   / |  \
// 1   2   3
// |       |
// 4       5
// Answer: [1, 3, 2]

package main

import (
	"fmt"
	"github.com/d-sense/algorithims-in-go/levelwidth/node"
)

func main(){
	//tree := &node.Tree{Data: 20}
	//tree.Add(0)
	//tree.Add(40)
	//tree.Add(-15)
	//
	//tree.Children[0].Add(12)
	//tree.Children[0].Add(-2)
	//tree.Children[0].Add(1)
	//
	//tree.Children[0].Add(-2)

	tree := &node.Tree{Data: 0}
	tree.Add(1)
	tree.Add(2)
	tree.Add(3)

	tree.Children[0].Add(4)
	tree.Children[2].Add(5)

	fmt.Println("Width Level:")
	fmt.Println(levelWidth(tree))
}

// the implementation is NOT accurate; with the sample above,
// the result should be [1, 3, 2] not [1, 3, 1]
func levelWidth(n *node.Tree) []int{
	counter := make([]int, 0)

	if n == nil {
		return counter
	}

	arr := []*node.Tree{n}

	for len(arr) > 0 {
		counter = append(counter, len(arr))
		currentNode := arr[0]

		fmt.Println("...:", arr[0])
		arr = currentNode.Children
	}
	return counter
}