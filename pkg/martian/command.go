package martian

// Command is the interface containing
// the Apply function, which any Command carried
// out against a Robot must implement
type Command interface {
	Apply(*Grid, *Robot) error
}

// CommandFunc wraps any Command function
// in the interface with Apply
type CommandFunc func(*Grid, *Robot) error

func (c CommandFunc) Apply(g *Grid, r *Robot) error {
	return c(g, r)
}

// CommandLeft takes a Robot and returns
// the Robot with the Left command applied to it
func CommandLeft(g *Grid, r *Robot) error {
	switch r.orientation {
	case OrientationNorth:
		r.orientation = OrientationWest
	case OrientationEast:
		r.orientation = OrientationNorth
	case OrientationSouth:
		r.orientation = OrientationEast
	case OrientationWest:
		r.orientation = OrientationSouth
	}
	return nil
}

// CommandRight takes a Robot and returns
// the Robot with the Right command applied to it
func CommandRight(g *Grid, r *Robot) error {
	switch r.orientation {
	case OrientationNorth:
		r.orientation = OrientationEast
	case OrientationEast:
		r.orientation = OrientationSouth
	case OrientationSouth:
		r.orientation = OrientationWest
	case OrientationWest:
		r.orientation = OrientationNorth
	}
	return nil
}

// CommandForward takes a Robot and returns
// the Robot with the Forward command applied to it
func CommandForward(g *Grid, r *Robot) error {
	// if robot is lost, ignore command
	if r.lost {
		return nil
	}

	newPosition := r.position

	switch r.orientation {
	case OrientationNorth:
		newPosition.y = r.position.y + 1
	case OrientationEast:
		newPosition.x = r.position.x + 1
	case OrientationSouth:
		newPosition.y = r.position.y - 1
	case OrientationWest:
		newPosition.x = r.position.x - 1
	}

	// Check if new position is on grid
	if newPosition.x > g.width || newPosition.y > g.height || newPosition.x < 0 || newPosition.y < 0 {
		// If a scent is already left here, ignore command
		if _, foundScent := g.lostScents[r.position]; foundScent {
			return nil
		}

		// Robot is lost :(
		r.lost = true
		// leave a scent here
		if g.lostScents == nil {
			g.lostScents = make(map[Coordinates]struct{})
		}
		g.lostScents[r.position] = struct{}{}
		return nil
	}

	// Update position
	r.position = newPosition

	return nil
}
