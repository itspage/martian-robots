package martian

// Robot represents a robot, it's position, orientation
// and whether it is lost
type Robot struct {
	Position Coordinates
	Orientation Orientation
	Lost        bool
}

// String representation of the status of the Robot
// example: `3 3 N LOST`
func (r *Robot) String() string {
	return ""
}
