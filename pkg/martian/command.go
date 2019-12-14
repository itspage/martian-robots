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
func CommandLeft(*Grid, *Robot) error {
	return nil
}

// CommandRight takes a Robot and returns
// the Robot with the Right command applied to it
func CommandRight(*Grid, *Robot) error {
	return nil
}

// CommandForward takes a Robot and returns
// the Robot with the Forward command applied to it
func CommandForward(*Grid, *Robot) error {
	return nil
}
