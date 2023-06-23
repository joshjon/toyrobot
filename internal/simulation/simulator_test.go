package simulation

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewSimulator(t *testing.T) {
	var buf bytes.Buffer
	reader := NewReader(&buf)
	executor := NewExecutor()
	simulator := NewSimulator(reader, executor, &buf)
	require.Equal(t, reader, simulator.reader)
	require.Equal(t, executor, simulator.executor)
	require.Equal(t, maxX, simulator.state.maxX)
	require.Equal(t, maxY, simulator.state.maxX)
	require.Equal(t, &buf, simulator.state.writer)
}

func TestSimulator_Run(t *testing.T) {
	mockErrs := make(chan error)
	close(mockErrs)

	r := &mockReader{}
	e := &mockExecutor{}
	r.On("Run", mock.Anything).Return(mockErrs).Once()
	e.On("Run", mock.Anything, mock.Anything).Return(mockErrs).Once()
	simulator := Simulator{
		reader:   r,
		executor: e,
	}
	errs := simulator.Run()
	for err := range errs {
		require.NoError(t, err)
	}
}

func TestSimulator_Run_integration(t *testing.T) {
	var bufWriter bytes.Buffer
	var bufReader bytes.Buffer
	bufReader.Write([]byte("PLACE 1,2,EAST\n"))
	bufReader.Write([]byte("MOVE\n"))
	bufReader.Write([]byte("MOVE\n"))
	bufReader.Write([]byte("LEFT\n"))
	bufReader.Write([]byte("MOVE\n"))
	bufReader.Write([]byte("MOVE\n"))
	bufReader.Write([]byte("RIGHT\n"))
	bufReader.Write([]byte("RIGHT\n"))
	bufReader.Write([]byte("RIGHT\n"))
	bufReader.Write([]byte("MOVE\n"))
	bufReader.Write([]byte("MOVE\n"))
	bufReader.Write([]byte("LEFT\n"))
	bufReader.Write([]byte("MOVE\n"))
	bufReader.Write([]byte("LEFT\n"))
	bufReader.Write([]byte("MOVE\n"))
	bufReader.Write([]byte("REPORT\n"))

	reader := NewReader(&bufReader)
	executor := NewExecutor()
	simulator := NewSimulator(reader, executor, &bufWriter)
	errs := simulator.Run()
	for err := range errs {
		require.NoError(t, err)
	}
	require.Equal(t, "2,3,EAST\n", bufWriter.String())
}

type mockReader struct {
	mock.Mock
}

func (r *mockReader) Run(commands chan<- Command) <-chan error {
	return r.Called(commands).Get(0).(chan error)
}

type mockExecutor struct {
	mock.Mock
}

func (r *mockExecutor) Run(commands <-chan Command, state *State) <-chan error {
	return r.Called(commands, state).Get(0).(chan error)
}
