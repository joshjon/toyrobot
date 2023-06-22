package direction

import (
	"errors"
	"strings"
)

var InvalidDirectionErr = errors.New("invalid direction")

type Direction int

const (
	beg Direction = iota
	North
	East
	South
	West
	end
)

func FromString(direction string) (Direction, error) {
	switch strings.ToUpper(direction) {
	case "NORTH":
		return North, nil
	case "EAST":
		return East, nil
	case "SOUTH":
		return South, nil
	case "WEST":
		return West, nil
	}
	return 0, InvalidDirectionErr
}

func (d Direction) Left() Direction {
	dir := d - 1
	if dir == beg {
		return West
	}
	return dir
}

func (d Direction) Right() Direction {
	dir := d + 1
	if dir == end {
		return North
	}
	return dir
}

// Axes maps the direction to its respective x or y-axis.
func (d Direction) Axes() (x int, y int) {
	switch d {
	case North:
		return 0, 1
	case South:
		return 0, -1
	case East:
		return 1, 0
	case West:
		return -1, 0
	}
	return 0, 0
}

func (d Direction) String() string {
	if d < beg || d > end {
		return "UNKNOWN"
	}
	return []string{"UNKNOWN", "NORTH", "EAST", "SOUTH", "WEST", "UNKNOWN"}[d]
}
