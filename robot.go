package robo

// Robot represents a robot capable of moving in 4 cardinal directions and keeping track of its relative position
type Robot interface {
	Move(Direction) error
	Position() Node
	Direction() Direction
	DistanceTravelled() int64
}

// DefaultRobot is the default implementation of the Robot interface
type DefaultRobot struct {
	pos               Node
	direction         Direction
	distanceTravelled int64
}

// Move moves the robot 1 step in the specified direction, returning an error if it is unable to do so
func (r *DefaultRobot) Move(d Direction) error {
	r.pos = Navigate(r.pos, d)
	r.distanceTravelled++
	return nil
}

// Position returns the relative position of the robot
func (r *DefaultRobot) Position() Node {
	return r.pos
}

// Direction returns the current direction of the robot
func (r *DefaultRobot) Direction() Direction {
	return r.direction
}

// DistanceTravelled returns the total distance travelled by the robot so far
func (r *DefaultRobot) DistanceTravelled() int64 {
	return r.distanceTravelled
}

// NewDefaultRobot creates a new robot which can move up, down, left, or right
func NewDefaultRobot() *DefaultRobot {
	return &DefaultRobot{
		pos:       Node{X: 0, Y: 0},
		direction: Up,
	}
}
