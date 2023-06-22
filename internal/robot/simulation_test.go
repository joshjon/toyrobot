package robot

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunSimulation(t *testing.T) {
	tc := &TestCommand{}
	commands := make(chan Command, 1)

	go func() {
		commands <- tc
		close(commands)
	}()

	RunSimulation(commands)
	require.NotNil(t, tc.got)
}

// TestCommand stores the passed in state on execution. It is used to verify
// that a command can be successfully executed in RunSimulation.
type TestCommand struct {
	got *State
}

func (c *TestCommand) Execute(state *State) {
	c.got = state
}
