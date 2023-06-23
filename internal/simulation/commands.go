package simulation

import (
	"github.com/joshjon/toyrobot/internal/direction"
)

// Command is a command that can be executed on a State in order to apply a
// mutation or retrieve data.
type Command interface {
	Execute(state *State) error
}

type CommandPlace struct {
	X         int
	Y         int
	Direction direction.Direction
}

func (c *CommandPlace) Execute(state *State) error {
	if state.withinBounds(c.X, c.Y) {
		state.posX = c.X
		state.posY = c.Y
		state.direction = c.Direction
		state.placed = true
	}
	return nil
}

type CommandMove struct{}

func (c *CommandMove) Execute(state *State) error {
	x, y := state.direction.Axes()
	if state.placed && state.withinBounds(state.posX+x, state.posY+y) {
		state.posX += x
		state.posY += y
	}
	return nil
}

type CommandRotateLeft struct{}

func (c *CommandRotateLeft) Execute(state *State) error {
	if state.placed {
		state.direction = state.direction.Left()
	}
	return nil
}

type CommandRotateRight struct{}

func (c *CommandRotateRight) Execute(state *State) error {
	if state.placed {
		state.direction = state.direction.Right()
	}
	return nil
}

type CommandReport struct{}

func (c *CommandReport) Execute(state *State) error {
	return state.report()
}
