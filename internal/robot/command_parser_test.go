package robot

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/joshjon/toyrobot/internal/direction"
)

func TestCommandFromString(t *testing.T) {
	tests := []struct {
		name    string
		command string
		want    Command
		wantErr error
	}{
		{
			name:    "place command",
			command: "PLACE 1,2,NORTH",
			want: &CommandPlace{
				X:         1,
				Y:         2,
				Direction: direction.North,
			},
		},
		{
			name:    "place command invalid x option",
			command: "PLACE X,2,NORTH",
			wantErr: CommandPlaceOptionXInvalidErr,
		},
		{
			name:    "place command invalid y option",
			command: "PLACE 1,Y,NORTH",
			wantErr: CommandPlaceOptionYInvalidErr,
		},
		{
			name:    "place command invalid x option",
			command: "PLACE 1,2,F",
			wantErr: CommandPlaceOptionFInvalidErr,
		},
		{
			name:    "place command invalid x option",
			command: "PLACE",
			wantErr: CommandPlaceOptionsErr,
		},
		{
			name:    "place command invalid x option",
			command: "PLACE 1,2",
			wantErr: CommandPlaceOptionsErr,
		},
		{
			name:    "face command",
			command: "MOVE",
			want:    &CommandMove{},
		},
		{
			name:    "left command",
			command: "LEFT",
			want:    &CommandRotateLeft{},
		},
		{
			name:    "right command",
			command: "RIGHT",
			want:    &CommandRotateRight{},
		},
		{
			name:    "report command",
			command: "REPORT",
			want:    &CommandReport{},
		},
		{
			name:    "command unknown",
			command: "INVALID",
			wantErr: CommandUnknownErr,
		},
		{
			name:    "command must not be empty",
			command: "",
			wantErr: CommandEmptyStringErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CommandFromString(tt.command)
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, got)
		})
	}
}
