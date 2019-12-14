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
			name: "L (N->W)",
			grid: &Grid{
				width:  5,
				height: 5,
			},
			robot: &Robot{
				position:    Coordinates{0, 0},
				orientation: OrientationNorth,
			},
			command: CommandLeft,
			wantRobot: &Robot{
				position:    Coordinates{0, 0},
				orientation: OrientationWest,
			},
			wantGrid: &Grid{
				width:  5,
				height: 5,
			},
		},
		{
			name: "R (N->E)",
			grid: &Grid{
				width:  5,
				height: 5,
			},
			robot: &Robot{
				position:    Coordinates{0, 0},
				orientation: OrientationNorth,
			},
			command: CommandRight,
			wantRobot: &Robot{
				position:    Coordinates{0, 0},
				orientation: OrientationEast,
			},
			wantGrid: &Grid{
				width:  5,
				height: 5,
			},
		},
		// TODO: Test turning from all orientations
		{
			name: "F (OK)",
			grid: &Grid{
				width:  5,
				height: 5,
			},
			robot: &Robot{
				position:    Coordinates{0, 0},
				orientation: OrientationNorth,
			},
			command: CommandForward,
			wantRobot: &Robot{
				position:    Coordinates{0, 1},
				orientation: OrientationNorth,
			},
			wantGrid: &Grid{
				width:  5,
				height: 5,
			},
		},
		{
			name: "F (gets lost)",
			grid: &Grid{
				width:  5,
				height: 5,
			},
			robot: &Robot{
				position:    Coordinates{5, 5},
				orientation: OrientationNorth,
			},
			command: CommandForward,
			wantRobot: &Robot{
				position:    Coordinates{5, 5},
				orientation: OrientationNorth,
				lost:        true,
			},
			wantGrid: &Grid{
				width:  5,
				height: 5,
				lostScents: map[Coordinates]struct{}{
					Coordinates{5, 5}: struct{}{},
				},
			},
		},
		{
			name: "F (avoids getting lost, scent)",
			grid: &Grid{
				width:  5,
				height: 5,
				lostScents: map[Coordinates]struct{}{
					Coordinates{5, 5}: struct{}{},
				},
			},
			robot: &Robot{
				position:    Coordinates{5, 5},
				orientation: OrientationNorth,
			},
			command: CommandForward,
			wantRobot: &Robot{
				position:    Coordinates{5, 5},
				orientation: OrientationNorth,
			},
			wantGrid: &Grid{
				width:  5,
				height: 5,
				lostScents: map[Coordinates]struct{}{
					Coordinates{5, 5}: struct{}{},
				},
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
