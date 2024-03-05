package piscine

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
