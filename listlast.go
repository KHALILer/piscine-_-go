package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListLast(l *List) interface{} {
	nodecurrent := l.Head
	for l.Head != nil {
		if nodecurrent.Next == nil {
			return nodecurrent.Data
		}
		nodecurrent = nodecurrent.Next
	}
	return nil
}
