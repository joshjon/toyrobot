package simulation

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/joshjon/toyrobot/internal/direction"
)

func TestState_Report(t *testing.T) {
	var buf bytes.Buffer
	state := placedState(3, 3, direction.East)
	state.writer = &buf
	err := state.report()
	require.NoError(t, err)
	got := buf.String()
	require.Equal(t, fmt.Sprintf("%d,%d,%s\n", state.posX, state.posY, state.direction.String()), got)
}
