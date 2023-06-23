package simulation

import "io"

const (
	maxX, maxY = 5, 5
	bufferSize = 50
)

type CommandReader interface {
	Run(commands chan<- Command) <-chan error
}

type CommandExecutor interface {
	Run(commands <-chan Command, state *State) <-chan error
}

type Simulator struct {
	reader   CommandReader
	executor CommandExecutor
	state    *State
}

func NewSimulator(reader *Reader, executor *Executor, writer io.Writer) *Simulator {
	return &Simulator{
		reader:   reader,
		executor: executor,
		state: &State{
			maxX:   maxX,
			maxY:   maxY,
			writer: writer,
		},
	}
}

// Run runs a new robot simulation. It reads in commands from the configured
// Reader and executes them via the Executor. The reading and execution logic
// is processed in separate goroutines, with channels used for async
// communication. Reporting is also done via the configured io.Writer.
//
// The behaviour of this implementation is similar to a pub/sub pattern, where
// events (commands) are published onto a queue (buffered channel) that is
// subscribed to by a worker which handles the events.
func (s *Simulator) Run() <-chan error {
	commands := make(chan Command, bufferSize)
	readErrs := s.reader.Run(commands)
	execErrs := s.executor.Run(commands, s.state)
	simulatorErrs := make(chan error, bufferSize)

	go func() {
		defer close(simulatorErrs)
		for {
			select {
			case readErr := <-readErrs:
				simulatorErrs <- readErr
			case execErr, ok := <-execErrs:
				if !ok {
					return
				}
				simulatorErrs <- execErr
			}
		}
	}()

	return simulatorErrs
}
