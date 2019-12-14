package martian_test

import (
	"testing"

	"github.com/itspage/martian-robots/pkg/martian"
)

func TestRobotString(t *testing.T) {
	tests := []struct {
		name  string
		robot martian.Robot
		want  string
	}{
		{
			name: "found robot",
			robot: martian.Robot{
				Position: martian.Coordinates{
					X: 2,
					Y: 3,
				},
				Orientation: martian.OrientationSouth,
			},
			want: "2 3 S",
		},
		{
			name: "lost robot",
			robot: martian.Robot{
				Position: martian.Coordinates{
					X: 3,
					Y: 3,
				},
				Orientation: martian.OrientationNorth,
				Lost:        true,
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
