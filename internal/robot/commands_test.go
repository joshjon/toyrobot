package robot

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/joshjon/toyrobot/internal/direction"
)

func TestCommandPlace_Execute(t *testing.T) {
	tests := []struct {
		name    string
		command *CommandPlace
		initial State
		want    State
	}{
		{
			name: "place command within bounds",
			command: &CommandPlace{
				X:         3,
				Y:         3,
				Direction: direction.North,
			},
			initial: defaultState(),
			want:    placedState(3, 3, direction.North),
		},
		{
			name: "place command outside of x min bounds",
			command: &CommandPlace{
				X:         -1,
				Y:         3,
				Direction: direction.North,
			},
			initial: defaultState(),
			want:    defaultState(),
		},
		{
			name: "place command outside of x max bounds",
			command: &CommandPlace{
				X:         maxX + 1,
				Y:         3,
				Direction: direction.North,
			},
			initial: defaultState(),
			want:    defaultState(),
		},
		{
			name: "place command outside of y min bounds",
			command: &CommandPlace{
				X:         0,
				Y:         -1,
				Direction: direction.North,
			},
			initial: defaultState(),
			want:    defaultState(),
		},
		{
			name: "place command outside of y max bounds",
			command: &CommandPlace{
				X:         0,
				Y:         maxY + 1,
				Direction: direction.North,
			},
			initial: defaultState(),
			want:    defaultState(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.initial
			tt.command.Execute(&got)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestCommandMove_Execute(t *testing.T) {
	tests := []struct {
		name    string
		command Command
		initial State
		want    State
	}{
		{
			name:    "move X within bounds",
			command: &CommandMove{},
			initial: placedState(3, 3, direction.West),
			want:    placedState(2, 3, direction.West),
		},
		{
			name:    "move Y within bounds",
			command: &CommandMove{},
			initial: placedState(3, 3, direction.North),
			want:    placedState(3, 4, direction.North),
		},
		{
			name:    "move hit min X bounds",
			command: &CommandMove{},
			initial: placedState(0, 3, direction.West),
			want:    placedState(0, 3, direction.West),
		},
		{
			name:    "move hit max X bounds",
			command: &CommandMove{},
			initial: placedState(maxX, 3, direction.East),
			want:    placedState(maxX, 3, direction.East),
		},
		{
			name:    "move hit min Y bounds",
			command: &CommandMove{},
			initial: placedState(3, 0, direction.South),
			want:    placedState(3, 0, direction.South),
		},
		{
			name:    "move hit max Y bounds",
			command: &CommandMove{},
			initial: placedState(3, maxY, direction.North),
			want:    placedState(3, maxY, direction.North),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.initial
			tt.command.Execute(&got)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestCommandRotateLeft_Execute(t *testing.T) {
	initial := placedState(0, 0, direction.North)
	want := placedState(0, 0, direction.West)
	got := initial

	command := &CommandRotateLeft{}
	command.Execute(&got)
	require.Equal(t, want, got)
}

func TestCommandRotateRight_Execute(t *testing.T) {
	initial := placedState(0, 0, direction.North)
	want := placedState(0, 0, direction.East)
	got := initial

	command := &CommandRotateRight{}
	command.Execute(&got)
	require.Equal(t, want, got)
}

func TestCommandReport_Execute(t *testing.T) {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = orig }()

	state := placedState(3, 3, direction.East)
	command := CommandReport{}
	command.Execute(&state)
	w.Close()
	out, _ := io.ReadAll(r)
	require.Equal(t, fmt.Sprintf("%d,%d,%s\n", state.posX, state.posY, state.direction.String()), string(out))
}

func defaultState() State {
	return State{
		maxX: maxX,
		maxY: maxY,
	}
}

func placedState(posX int, posY int, direction direction.Direction) State {
	return State{
		maxX:      maxX,
		maxY:      maxY,
		posX:      posX,
		posY:      posY,
		direction: direction,
		placed:    true,
	}
}
