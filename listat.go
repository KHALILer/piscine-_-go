package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {

	count := 0
	current := l
	for current != nil {
		if count == pos {
			return current
		}
		count++
		current = current.Next

	}
	return nil

}
