package robo

import "fmt"

// Node represents a coordinate on a 2D plane
type Node struct {
	X int64
	Y int64
}

func (n Node) String() string {
	return fmt.Sprintf("(%d,%d)", n.X, n.Y)
}

// AddNode finds the sum of two Nodes
func AddNode(a Node, b Node) Node {
	return Node{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}
