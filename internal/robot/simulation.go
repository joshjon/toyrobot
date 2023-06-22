package robot

import (
	"github.com/joshjon/toyrobot/internal/direction"
)

const maxX, maxY = 5, 5

type Command interface {
	Execute(state *State)
}

type State struct {
	maxX      int
	maxY      int
	posX      int
	posY      int
	direction direction.Direction
	placed    bool
}

func RunSimulation(commands <-chan Command) {
	state := State{
		maxX: maxX,
		maxY: maxY,
	}

	for command := range commands {
		command.Execute(&state)
	}
}
