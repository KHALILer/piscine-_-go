package piscine

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
