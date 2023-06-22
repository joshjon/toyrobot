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

func (s *State) withinBounds(x int, y int) bool {
	return x >= 0 && x <= s.maxX && y >= 0 && y <= s.maxY
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
