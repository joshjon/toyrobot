package simulation

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/joshjon/toyrobot/internal/direction"
)

var (
	CommandUnknownErr             = errors.New("unknown command")
	CommandEmptyStringErr         = errors.New("command cannot be empty")
	CommandPlaceOptionsErr        = errors.New("command PLACE requires X,Y,F options")
	CommandPlaceOptionXInvalidErr = errors.New("option X must be a valid integer")
	CommandPlaceOptionYInvalidErr = errors.New("option Y must be a valid integer")
	CommandPlaceOptionFInvalidErr = errors.New("option F must be a valid direction (NORTH, SOUTH, EAST, WEST)")
)

const (
	CommandPlaceStr  = "PLACE"
	CommandMoveStr   = "MOVE"
	CommandLeftStr   = "LEFT"
	CommandRightStr  = "RIGHT"
	CommandReportStr = "REPORT"
)

type Reader struct {
	reader io.Reader
}

func NewReader(reader io.Reader) *Reader {
	return &Reader{
		reader: reader,
	}
}

// Run reads in lines from the configured io.Reader, parses them into commands,
// and sends them to the commands channel param. It is intended to be used
// alongside an Executor, which is responsible for executing commands that are
// produced onto a chanel.
func (r *Reader) Run(commands chan<- Command) <-chan error {
	errs := make(chan error, 1)

	go func() {
		defer close(commands)
		defer close(errs)
		scanner := bufio.NewScanner(r.reader)

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

	return errs
}

// CommandFromString parses a string into a Command and returns any validation
// errors that may have occurred in the process.
func CommandFromString(command string) (Command, error) {
	if len(command) == 0 {
		return nil, CommandEmptyStringErr
	}

	parts := strings.Split(command, " ")

	switch strings.ToUpper(parts[0]) {
	case CommandPlaceStr:
		return parseMoveCommand(parts)
	case CommandMoveStr:
		return &CommandMove{}, nil
	case CommandLeftStr:
		return &CommandRotateLeft{}, nil
	case CommandRightStr:
		return &CommandRotateRight{}, nil
	case CommandReportStr:
		return &CommandReport{}, nil
	}
	return nil, CommandUnknownErr
}

func parseMoveCommand(parts []string) (Command, error) {
	if len(parts) != 2 {
		return nil, CommandPlaceOptionsErr
	}

	opts := strings.Split(parts[1], ",")
	if len(opts) != 3 {
		return nil, CommandPlaceOptionsErr
	}

	x, err := strconv.Atoi(opts[0])
	if err != nil {
		return nil, CommandPlaceOptionXInvalidErr
	}

	y, err := strconv.Atoi(opts[1])
	if err != nil {
		return nil, CommandPlaceOptionYInvalidErr
	}

	d, err := direction.FromString(opts[2])
	if err != nil {
		return nil, CommandPlaceOptionFInvalidErr
	}

	return &CommandPlace{
		X:         x,
		Y:         y,
		Direction: d,
	}, nil
}
