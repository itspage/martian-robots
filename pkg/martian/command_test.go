package martian

import (
	"reflect"
	"testing"
)

var smallGrid *Grid

func TestCommands(t *testing.T) {
	tests := []struct {
		name      string
		grid      *Grid
		robot     *Robot
		command   CommandFunc
		wantRobot *Robot
		wantGrid  *Grid
	}{
		{
			name: "F",
			grid: &Grid{
				width:  5,
				height: 5,
			},
			robot: &Robot{
				Position:    Coordinates{0, 0},
				Orientation: OrientationNorth,
			},
			command: CommandForward,
			wantRobot: &Robot{
				Position:    Coordinates{0, 1},
				Orientation: OrientationNorth,
			},
			wantGrid: &Grid{
				width:  5,
				height: 5,
			},
		},
	}

	for _, tt := range tests {
		if err := tt.command(tt.grid, tt.robot); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(tt.robot, tt.wantRobot) {
			t.Errorf("command() robot = %v, want %v", tt.robot, tt.wantRobot)
		}
		if !reflect.DeepEqual(tt.grid, tt.wantGrid) {
			t.Errorf("command() grid = %v, want %v", tt.grid, tt.wantGrid)
		}
	}
}
