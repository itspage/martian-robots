package martian

import (
	"testing"
)

func TestRobotString(t *testing.T) {
	tests := []struct {
		name  string
		robot Robot
		want  string
	}{
		{
			name: "found robot",
			robot: Robot{
				position: Coordinates{
					x: 2,
					y: 3,
				},
				orientation: OrientationSouth,
			},
			want: "2 3 S",
		},
		{
			name: "lost robot",
			robot: Robot{
				position: Coordinates{
					x: 3,
					y: 3,
				},
				orientation: OrientationNorth,
				lost:        true,
			},
			want: "3 3 N LOST",
		},
	}

	for _, tt := range tests {
		if got := tt.robot.String(); got != tt.want {
			t.Fatalf("String() = got %s, want %s", got, tt.want)
		}
	}
}
