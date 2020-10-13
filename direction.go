package robo

// Direction represents the four cardinal directions
type Direction int

const (
	// Up is an integer representation of the direction
	Up Direction = iota
	// Right is an integer representation of the direction
	Right
	// Down is an integer representation of the direction
	Down
	// Left is an integer representation of the direction
	Left
)

// String returns a human readable representation of the direction
func (d Direction) String() string {
	return [...]string{"Up", "Right", "Down", "Left"}[d]
}

// Directions is defines each direction relative to the origin Node (0, 0)
var Directions = map[Direction]Node{
	Up:    Node{X: 0, Y: 1},
	Right: Node{X: 1, Y: 0},
	Down:  Node{X: 0, Y: -1},
	Left:  Node{X: -1, Y: 0},
}

// oppositeDirection returns the direction opposite (180degrees)
func oppositeDirection(d Direction) Direction {
	return (4 + d + 2) % 4
}

// Navigate returns the Node reached after a move in the direction
func Navigate(n Node, d Direction) Node {
	return AddNode(n, Directions[d])
}
