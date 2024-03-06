package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListClear(l *List) {

	current := l.Head
	for current != nil {
		current.Data = nil
		current.Next = nil
		current = current.Next
	}
	l.Head = nil
}
