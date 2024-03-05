package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}


func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}
	if l.Head == nil {
		l.Tail = newNode
		l.Head = newNode
		return
	}
	newNode.Next = l.Head
	l.Head = newNode
}

func ListSize(l *List) int {
	count := 0
	for l.Head != nil {
		count++
		l.Head = l.Head.Next
	}
	return count
}
