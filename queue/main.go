// --- Description
// Create a queue data structure.  The queue
// should be a class with methods 'add' and 'remove'.
// Adding to the queue should store an element until
// it is removed
// --- Examples
//     const q = new Queue();
//     q.add(1);
//     q.remove(); // returns 1;
package main

import "fmt"

func main(){
	fmt.Println("start ...")
    q := Queue{}
    q.add(1)
    q.add(3)
    q.add(5)
    fmt.Println(q.store)
    q.remove()
	fmt.Println(q.store)
}

type Queue struct {
	store []int
}

func (q *Queue) add(n int)  {
	 q.store = append([]int{n}, q.store...)
}

func (q *Queue) remove() int{
	if len(q.store) == 0 {
		return 0
	}

	length := len(q.store)
	q.store = q.store[0:length-1]

	return 1
}

func (q *Queue) print() []int{
	return q.store
}


