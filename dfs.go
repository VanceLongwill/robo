package robo

import "context"

// DFSer is used to perform and hold the state of a Depth First Search with an optional Robot which will follow the path generated
type DFSer struct {
	r       Robot
	visited VisitMap
	scanner Scanner
}

// recursiveDFS performs a recursive Depth-first search while moving the robot around the area accordingly. If at an time the robot is unable to move to a Node, the search is aborted and an error returned. The search is also aborted (no error) when the context is cancelled.
func (d *DFSer) recursiveDFS(ctx context.Context, n Node, dir Direction) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		hasRobot := d.r != nil
		// if a robot has been provided
		if hasRobot {
			// move the robot to the new Node
			if err := d.r.Move(dir); err != nil {
				return err
			}
		}
		// mark it as visited
		d.visited.Add(n)
		// iterate of adjacent Nodes
		for _, adj := range d.scanner.Scan(n) {
			if !d.visited.Contains(adj.Node) {
				// recursively call the DFS algorithm with the unvisited Node
				if err := d.recursiveDFS(ctx, adj.Node, adj.Direction); err != nil {
					return err
				}
			}
		}
		if hasRobot {
			// having got the areas around the current location, walk robot back to where it was from
			if err := d.r.Move(oppositeDirection(dir)); err != nil {
				return err
			}
		}
	}
	return nil
}

// Run initiates a DFS search
func (d *DFSer) Run(ctx context.Context) error {
	var pos Node
	var dir Direction
	if d.r != nil {
		pos = d.r.Position()
		dir = d.r.Direction()
	}
	return d.recursiveDFS(ctx, pos, dir)
}

// Visited exposes the current visited nodes
func (d *DFSer) Visited() VisitMap {
	return d.visited
}

// NewDFS sets up a new Depth First Search with sensible defaults
func NewDFS(options ...Option) *DFSer {
	d := &DFSer{}
	d.visited = DefaultVisitMap{}
	d.scanner = DefaultScanner

	for _, opt := range options {
		opt(d)
	}

	return d
}

// Option is used to override defaults
type Option func(*DFSer)

// WithRobot specifies a custom robot
func WithRobot(r Robot) Option {
	return func(d *DFSer) {
		d.r = r
	}
}

// WithVisitMap specifies a custom implementation of a visit map
func WithVisitMap(v VisitMap) Option {
	return func(d *DFSer) {
		d.visited = v
	}
}

// WithScanner specifies a custom scanner (the default scanner just returns adjacent nodes in all directions)
func WithScanner(s Scanner) Option {
	return func(d *DFSer) {
		d.scanner = s
	}
}
