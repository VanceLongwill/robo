package robo

// Adjacent contains the position and relative direction of an adjacent Node
type Adjacent struct {
	Direction
	Node
}

// Scanner is used to find adjacent nodes
type Scanner interface {
	Scan(Node) []Adjacent
}

// ScannerFunc allows an Scanner to be declared as a standalone function
type ScannerFunc func(Node) []Adjacent

// Scan implements the Scanner interface
func (fn ScannerFunc) Scan(n Node) []Adjacent {
	return fn(n)
}

// DefaultScanner is the default implementation of the scanner interface and returns neighbouring nodes in all directions
var DefaultScanner = ScannerFunc(func(n Node) []Adjacent {
	adjs := make([]Adjacent, len(Directions))
	for d := range Directions {
		node := Navigate(n, d)
		adjs[d] = Adjacent{Node: node, Direction: d}
	}
	return adjs
})
