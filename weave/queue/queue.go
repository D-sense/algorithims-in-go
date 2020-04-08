// --- Directions
// Implement a 'peek' method in this Queue class.
// Peek should return the last element (the next
// one to be returned) from the queue *without*
// removing it.

package queue

type Queue struct {
	store []int
}

func (q *Queue) Add(n int)  {
	q.store = append([]int{n}, q.store...)
}

func (q *Queue) Remove() int{
	if len(q.store) == 0 {
		return 0
	}

	length := len(q.store)
	q.store = q.store[0:length-1]

	return 1
}

func (q *Queue) Peek() int{
	if len(q.store) == 0 {
		return -1
	}

	length := len(q.store)

	return q.store[length-1]
}

func (q *Queue) Print() []int{
	return q.store
}


