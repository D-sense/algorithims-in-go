package main

type Stack struct {
	items []int
}

func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) pop() int {
	if len(s.items) == 0 {
		return -1
	}

	items, item := s.items[:len(s.items)-1], s.items[len(s.items)-1]
	s.items = items
	return item
}



type Queue struct {
	items []int
}

func (s *Queue) enqueue(i int) {
	s.items = append(s.items, i)
}

func (s *Queue) dequeue() int {
	if len(s.items) == 0 {
		return -1
	}

	items, item := s.items[1:], s.items[0]
	s.items = items
	return item
}
