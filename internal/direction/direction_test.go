package direction

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromString(t *testing.T) {
	tests := []struct {
		direction string
		want      Direction
		wantErr   error
	}{
		{
			direction: "NORTH",
			want:      North,
		},
		{
			direction: "SOUTH",
			want:      South,
		},
		{
			direction: "EAST",
			want:      East,
		},
		{
			direction: "WEST",
			want:      West,
		},
		{
			direction: "INVALID",
			wantErr:   InvalidDirectionErr,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("direction from %s", tt.direction), func(t *testing.T) {
			got, err := FromString(tt.direction)
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestDirection_Left(t *testing.T) {
	require.Equal(t, West, North.Left())
	require.Equal(t, South, West.Left())
	require.Equal(t, East, South.Left())
	require.Equal(t, North, East.Left())
}

func TestDirection_Right(t *testing.T) {
	require.Equal(t, East, North.Right())
	require.Equal(t, South, East.Right())
	require.Equal(t, West, South.Right())
	require.Equal(t, North, West.Right())
}

func TestDirection_Axis(t *testing.T) {
	tests := []struct {
		direction Direction
		wantX     int
		wantY     int
	}{
		{
			direction: North,
			wantX:     0,
			wantY:     1,
		},
		{
			direction: South,
			wantX:     0,
			wantY:     -1,
		},
		{
			direction: East,
			wantX:     1,
			wantY:     0,
		},
		{
			direction: West,
			wantX:     -1,
			wantY:     0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.direction.String(), func(t *testing.T) {
			gotX, gotY := tt.direction.Axes()
			require.Equal(t, tt.wantX, gotX)
			require.Equal(t, tt.wantY, gotY)
		})
	}
}

func TestDirection_String(t *testing.T) {
	wantStrs := map[Direction]string{
		North:   "NORTH",
		South:   "SOUTH",
		East:    "EAST",
		West:    "WEST",
		beg:     "UNKNOWN",
		end:     "UNKNOWN",
		beg - 1: "UNKNOWN",
		end + 1: "UNKNOWN",
	}

	for dir, want := range wantStrs {
		got := dir.String()
		require.Equal(t, want, got, "expected %s but got %s", want, got)
	}
}
