package node


type Tree struct {
	Data int
	Children []*Tree
}

func NewTree(data int) *Tree {
	return &Tree{
		Data: data,
	}
}

func (n *Tree) Add(data int) {
	if n.Children == nil {
		n.Children =  append(n.Children, NewTree(data))
		return
	}
	n.Children = append(n.Children, NewTree(data))
}