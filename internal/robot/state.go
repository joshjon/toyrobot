package robot

import (
	"fmt"
	"io"

	"github.com/joshjon/toyrobot/internal/direction"
)

type State struct {
	maxX      int
	maxY      int
	posX      int
	posY      int
	direction direction.Direction
	placed    bool
	out       io.Writer
}

func (s *State) Report() error {
	_, err := fmt.Fprintf(s.out, "%d,%d,%s\n",
		s.posX, s.posY, s.direction.String(),
	)
	return err
}

func (s *State) withinBounds(x int, y int) bool {
	return x >= 0 && x <= s.maxX && y >= 0 && y <= s.maxY
}
