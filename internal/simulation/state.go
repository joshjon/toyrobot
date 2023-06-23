package simulation

import (
	"fmt"
	"io"
	"sync"

	"github.com/joshjon/toyrobot/internal/direction"
)

// State represents the overall robot state.
type State struct {
	sync.Mutex
	maxX      int
	maxY      int
	posX      int
	posY      int
	direction direction.Direction
	placed    bool
	writer    io.Writer
}

// report reports the current robot state to the configured writer.
func (s *State) report() error {
	_, err := fmt.Fprintf(s.writer, "%d,%d,%s\n",
		s.posX, s.posY, s.direction.String(),
	)
	return err
}

// withinBounds checks if the provided x,y coordinates are within the max bounds
// of the simulator.
func (s *State) withinBounds(x int, y int) bool {
	return x >= 0 && x <= s.maxX && y >= 0 && y <= s.maxY
}
