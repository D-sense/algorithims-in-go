package stack

type Stack struct {
	Store []int
}

func (q *Stack) Add(n int)  {
	q.Store = append([]int{n}, q.Store...)
}

func (q *Stack) Remove() int{
	if len(q.Store) == 0 {
		return 0
	}

	rec := q.Store[0]
	q.Store = q.Store[1:]

	return rec
}

func (q *Stack)  Peek() int {
	if len(q.Store) == 0 {
		return -1
	}
	return q.Store[len(q.Store)-1]
}

func (q *Stack) Print() []int{
	return q.Store
}