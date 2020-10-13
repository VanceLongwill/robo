package robo

// VisitMap is used to keep track of visited nodes
type VisitMap interface {
	Add(Node)
	Contains(Node) bool
	Len() int
}

// DefaultVisitMap is the default implementation of the VisitMap interface
type DefaultVisitMap map[Node]struct{}

// Add appends a Node to the the visited list
func (v DefaultVisitMap) Add(n Node) {
	v[n] = struct{}{}
}

// Len returns the number of nodes visited
func (v DefaultVisitMap) Len() int {
	return len(v)
}

// Contains checks whether a Node has been visited
func (v DefaultVisitMap) Contains(n Node) bool {
	_, ok := v[n]
	return ok
}

// NewVisitMap creates a new DefaultVisitMap
func NewVisitMap() DefaultVisitMap {
	return DefaultVisitMap{}
}
