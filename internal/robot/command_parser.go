package robot

import (
	"errors"
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
