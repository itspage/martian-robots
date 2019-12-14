package cli

import (
	"fmt"

	"github.com/itspage/martian-robots/pkg/martian"
)

const (
	commandLeft    = "L"
	commandRight   = "R"
	commandForward = "F"
)

// CLI exposes the martian packages as a command line interface
// a CLI initialised with CLI{} is ready to use
type CLI struct {
	grid         *martian.Grid
	robots       []*martian.Robot
	currentRobot *martian.Robot
}

// ReadLine takes a line of input and processes it
func (c *CLI) ReadLine(line string) error {
	// If no grid - assume first line and create
	if c.grid == nil {
		width, height := 0, 0
		if _, err := fmt.Sscanf(line, "%d %d", &width, &height); err != nil {
			return fmt.Errorf("syntax error: %v", err)
		}
		var err error
		c.grid, err = martian.NewGrid(width, height)
		return err
	}

	// If no current robot - assume create
	if c.currentRobot == nil {
		x, y := 0, 0
		o := ""
		if _, err := fmt.Sscanf(line, "%d %d %s", &x, &y, &o); err != nil {
			return fmt.Errorf("syntax error: %v", err)
		}
		var err error
		c.currentRobot, err = martian.NewRobot(x, y, martian.Orientation(o))
		return err
	}

	// otherwise assume instructions
	if len(line) >= 100 {
		return ErrMaxInstructionLength
	}

	for _, i := range line {
		// parse command
		var command martian.Command
		switch string(i) {
		case commandLeft:
			command = martian.CommandFunc(martian.CommandLeft)
		case commandRight:
			command = martian.CommandFunc(martian.CommandRight)
		case commandForward:
			command = martian.CommandFunc(martian.CommandForward)
		default:
			return ErrInvalidCommand
		}

		// apply command
		if err := command.Apply(c.grid, c.currentRobot); err != nil {
			return fmt.Errorf("error carrying out instruction: %v", err)
		}

	}
	// successfully carried out instructions on this robot
	c.robots = append(c.robots, c.currentRobot)
	c.currentRobot = nil

	return nil
}

// Output returns the current (string representation) state of all robots
func (c *CLI) Output() ([]string, error) {
	lines := []string{}
	for _, r := range c.robots {
		lines = append(lines, r.String())
	}
	return lines, nil
}
