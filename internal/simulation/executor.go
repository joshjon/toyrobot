package simulation

import (
	"fmt"
)

type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

// Run reads in commands from the supplied channel and executes them on the
// State. It is intended to be used alongside a Reader which can produce commands
// onto a channel from any io.Reader implementation.
func (s *Executor) Run(commands <-chan Command, state *State) <-chan error {
	errs := make(chan error, 1)
	go func() {
		defer close(errs)
		for {
			select {
			case command, ok := <-commands:
				if !ok {
					return
				}
				state.Lock()
				if err := command.Execute(state); err != nil {
					errs <- fmt.Errorf("unexpected error: %w", err)
				}
				state.Unlock()
			}
		}
	}()
	return errs
}
