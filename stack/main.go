// --- Directions
// Create a stack data structure.  The stack
// should be a class with methods 'push', 'pop', and
// 'peek'.  Adding an element to the stack should
// store it until it is removed.
// --- Examples
//   const s = new Stack();
//   s.push(1);
//   s.push(2);
//   s.pop(); // returns 2
//   s.pop(); // returns 1

package main

import "fmt"

func main(){
	fmt.Println("start ...")
	q := Stack{}
	q.add(1)
	q.add(3)
	q.add(5)
	fmt.Println(q.store)
	q.remove()
	fmt.Println(q.store)
}

type Stack struct {
	store []int
}

func (q *Stack) add(n int)  {
	q.store = append([]int{n}, q.store...)
}

func (q *Stack) remove() int{
	if len(q.store) == 0 {
		return 0
	}

	q.store = q.store[1:]

	return 1
}

func (q *Stack) print() []int{
	return q.store
}