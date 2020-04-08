// --- Directions
// Implement a Queue data structure using two stacks.
// *Do not* create an array inside of the 'Queue' class.
// Queue should implement the methods 'add', 'remove', and 'peek'.
// For a reminder on what each method does, look back
// at the Queue exercise.
// --- Examples
//     const q = new Queue();
//     q.add(1);
//     q.add(2);
//     q.peek();  // returns 1
//     q.remove(); // returns 1
//     q.remove(); // returns 2

package main

import (
	"fmt"
	"github.com/d-sense/algorithims-in-go/qfroms/stack"
)

type Queue struct {
	first  stack.Stack
	second stack.Stack
}


func main(){
	fmt.Println("start ...")
	q := New()
	q.add(1)
	q.add(3)
	q.add(6)

	fmt.Println("peek: ", q.peek())

	fmt.Println("remove: ", q.remove())
	fmt.Println("remove: ", q.remove())
	fmt.Println("remove: ", q.remove())
}


func New() *Queue {
	return &Queue{
		first:  stack.Stack{
			Store: []int{},
		},
		second: stack.Stack{
			Store: []int{},
		},
	}
}

func (q *Queue) add(n int){
	q.first.Store = append([]int{n}, q.first.Store...)
}

func (q *Queue) peek() int {
	for  q.first.Peek() > -1 {
		q.second.Add(q.first.Remove())
	}

    record := q.second.Peek()

	for q.second.Peek() > -1 {
		q.first.Add(q.second.Remove())
	}

    return record
}

func (q *Queue) remove() int {

	for q.first.Peek() > -1 {
		q.second.Add(q.first.Remove())
	}

	record := q.second.Remove()

	for q.second.Peek() > -1 {
		q.first.Add(q.second.Remove())
	}

	return record
}