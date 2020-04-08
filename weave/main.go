// --- Directions
// 1) Complete the task in weave/queue.js
// 2) Implement the 'weave' function.  Weave
// receives two queues as arguments and combines the
// contents of each into a new, third queue.
// The third queue should contain the *alterating* content
// of the two queues.  The function should handle
// queues of different lengths without inserting
// 'undefined' into the new one.
// *Do not* access the array inside of any queue, only
// use the 'add', 'remove', and 'peek' functions.
// --- Example
//    const queueOne = new Queue();
//    queueOne.add(1);
//    queueOne.add(2);
//    const queueTwo = new Queue();
//    queueTwo.add('Hi');
//    queueTwo.add('There');
//    const q = weave(queueOne, queueTwo);
//    q.remove() // 1
//    q.remove() // 'Hi'
//    q.remove() // 2
//    q.remove() // 'There'

package main

import (
	"fmt"
	"github.com/d-sense/algorithims-in-go/weave/queue"
)

var fQueue queue.Queue
var sQueue queue.Queue
var resultQueue queue.Queue

func main(){
	fmt.Println("start ...")
	fQueue.Add(1)
	fQueue.Add(2)
    //.Println(fQueue.Print())

	sQueue.Add(8)
	sQueue.Add(9)
	//fmt.Println(sQueue.Print())

	resultQueue = weave(fQueue, sQueue)
	fmt.Println(resultQueue.Print())
}

func weave(fQueue, sQueue queue.Queue) queue.Queue {
	var newQueue queue.Queue

	for  {
		if elem1 := fQueue.Peek(); elem1 != -1 {
			newQueue.Add(elem1)
			fQueue.Remove()
		}

		if elem2 := sQueue.Peek(); elem2 != -1 {
			newQueue.Add(elem2)
			sQueue.Remove()
			continue
		}

		return newQueue
	}
}