package simulation

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewExecutor(t *testing.T) {
	executor := NewExecutor()
	require.NotNil(t, executor)
}

func TestExecutor_Run(t *testing.T) {
	state := &State{}
	executor := Executor{}

	wantCommands := 10
	m := &mockCommand{}
	m.On("Execute", state).Return(nil).Times(wantCommands)
	commands := make(chan Command, wantCommands)
	for i := 0; i < wantCommands; i++ {
		commands <- m
	}
	close(commands)

	errs := executor.Run(commands, state)
	for err := range errs {
		require.NoError(t, err)
	}

	m.AssertExpectations(t)
}

type mockCommand struct {
	mock.Mock
}

func (c *mockCommand) Execute(state *State) error {
	return c.Called(state).Error(0)
}
