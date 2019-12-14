package cli_test

import (
	"reflect"
	"testing"

	"github.com/itspage/martian-robots/pkg/cli"
)

func TestMartianRobots(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		want     []string
		wantErrs []error
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
		{
			name: "max instruction length",
			input: []string{
				"5 3",
				"1 1 E",
				"RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR",
			},
			wantErrs: []error{nil, nil, cli.ErrMaxInstructionLength},
			want:     []string{},
		},
		{
			name: "invalid command",
			input: []string{
				"5 3",
				"1 1 E",
				"S",
			},
			wantErrs: []error{nil, nil, cli.ErrInvalidCommand},
			want:     []string{},
		},
	}

	for _, tt := range tests {

		cli := cli.CLI{}

		for i, line := range tt.input {
			if err := cli.ReadLine(line); err != nil {
				if tt.wantErrs != nil && err != tt.wantErrs[i] {
					t.Fatalf("ReadLine() error got = %v, want %v", err, tt.wantErrs[i])
				}
			}
		}
		got, err := cli.Output()
		if err != nil {
			t.Fatalf("Output() error %v", err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Output() = %v, want %v", got, tt.want)
		}

	}
}
