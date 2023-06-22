package robot

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunSimulation(t *testing.T) {
	var bufWriter bytes.Buffer
	var bufReader bytes.Buffer
	bufReader.Write([]byte("PLACE 0,0,NORTH\n"))
	bufReader.Write([]byte("MOVE\n"))
	bufReader.Write([]byte("REPORT\n"))

	errs := RunSimulation(bytes.NewReader(bufReader.Bytes()), &bufWriter)
	for err := range errs {
		require.NoError(t, err)
	}

	require.Equal(t, "0,1,NORTH\n", bufWriter.String())
}

// TestCommand stores the passed in state on execution. It is used to verify
// that a command can be successfully executed in RunSimulation.
type TestCommand struct {
	got *State
}

func (c *TestCommand) Execute(state *State) error {
	c.got = state
	return nil
}
