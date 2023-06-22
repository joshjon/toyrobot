package robot

import (
	"bufio"
	"fmt"
	"io"
)

const maxX, maxY = 5, 5

func RunSimulation(in io.Reader, out io.Writer) <-chan error {
	commands, readErrs := readCommands(in)
	errs := make(chan error, 1)

	go func() {
		defer close(errs)
		state := State{
			maxX: maxX,
			maxY: maxY,
			out:  out,
		}
		for {
			select {
			case command, ok := <-commands:
				if !ok {
					return
				}
				if err := command.Execute(&state); err != nil {
					errs <- fmt.Errorf("unexpected error: %w", err)
				}
			case err := <-readErrs:
				if err != nil {
					errs <- err
				}
			}
		}
	}()

	return errs
}

func readCommands(reader io.Reader) (<-chan Command, <-chan error) {
	commands := make(chan Command, 50)
	errs := make(chan error, 1)

	go func() {
		defer close(commands)
		defer close(errs)
		scanner := bufio.NewScanner(reader)

		for scanner.Scan() {
			if err := scanner.Err(); err != nil {
				errs <- fmt.Errorf("unexpected error: %w", err)
				continue
			}

			line := scanner.Text()
			command, err := CommandFromString(line)
			if err != nil {
				errs <- fmt.Errorf("bad input: %w", err)
				continue
			}
			commands <- command
		}
	}()

	return commands, errs
}
