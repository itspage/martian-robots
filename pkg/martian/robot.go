package martian

import "fmt"

// Robot represents a robot, it's position, orientation
// and whether it is lost
type Robot struct {
	position    Coordinates
	orientation Orientation
	lost        bool
}

func NewRobot(x, y int, o Orientation) (*Robot, error) {
	switch o {
	case OrientationNorth, OrientationEast, OrientationSouth, OrientationWest:
	default:
		return nil, ErrInvalidOrientation
	}
	return &Robot{
		position: Coordinates{
			x: x,
			y: y,
		},
		orientation: o,
	}, nil
}

// String representation of the status of the Robot
// example: `3 3 N LOST`
func (r *Robot) String() string {
	positionStr := fmt.Sprintf("%d %d %s", r.position.x, r.position.y, r.orientation)
	if r.lost {
		return fmt.Sprintf("%s LOST", positionStr)
	}
	return positionStr
}
