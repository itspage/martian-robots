package cli_test

import "github.com/itspage/martian-robots/pkg/cli"

import (
	"reflect"
	"testing"
)

func TestMartianRobots(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name: "sample",
			input: []string{
				"5 3",
				"1 1 E",
				"RFRFRFRF",
				"3 2 N",
				"FRRFLLFFRRFLL",
				"0 3 W",
				"LLFFFLFLFL",
			},
			want: []string{
				"1 1 E",
				"3 3 N LOST",
				"2 3 S",
			},
		},
	}

	for _, tt := range tests {

		cli := cli.CLI{}

		for _, line := range tt.input {
			if err := cli.ReadLine(line); err != nil {
				t.Fatalf("ReadLine error %v", err)
			}
		}
		got, err := cli.Output()
		if err != nil {
			t.Fatalf("Output error %v", err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Output() = %v, want %v", got, tt.want)
		}

	}
}
