package robot

import (
	"fmt"

	"github.com/joshjon/toyrobot/internal/direction"
)

type CommandPlace struct {
	X         int
	Y         int
	Direction direction.Direction
}

func (c *CommandPlace) Execute(state *State) {
	if state.withinBounds(c.X, c.Y) {
		state.posX = c.X
		state.posY = c.Y
		state.direction = c.Direction
		state.placed = true
	}
}

type CommandMove struct{}

func (c *CommandMove) Execute(state *State) {
	x, y := state.direction.Axes()
	if state.placed && state.withinBounds(state.posX+x, state.posY+y) {
		state.posX += x
		state.posY += y
	}
}

type CommandRotateLeft struct{}

func (c *CommandRotateLeft) Execute(state *State) {
	if state.placed {
		state.direction = state.direction.Left()
	}
}

type CommandRotateRight struct{}

func (c *CommandRotateRight) Execute(state *State) {
	if state.placed {
		state.direction = state.direction.Right()
	}
}

type CommandReport struct{}

func (c *CommandReport) Execute(state *State) {
	fmt.Printf("%d,%d,%s\n", state.posX, state.posY, state.direction.String())
}
