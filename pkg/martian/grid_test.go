package martian_test

import (
	"testing"

	"github.com/itspage/martian-robots/pkg/martian"
)

func TestNewGrid(t *testing.T) {
	tests := []struct {
		name    string
		width   int
		height  int
		wantErr error
	}{
		{
			name:    "max width",
			width:   51,
			height:  1,
			wantErr: martian.ErrMaxGridWidth,
		},
		{
			name:    "max height",
			width:   1,
			height:  51,
			wantErr: martian.ErrMaxGridHeight,
		},
		{
			name:    "min width",
			width:   0,
			height:  1,
			wantErr: martian.ErrMinGridWidth,
		},
		{
			name:    "min height",
			width:   1,
			height:  0,
			wantErr: martian.ErrMinGridHeight,
		},
	}

	for _, tt := range tests {
		_, gotErr := martian.NewGrid(tt.width, tt.height)
		if tt.wantErr != gotErr {
			t.Errorf("NewGrid() = %v, want %v", gotErr, tt.wantErr)
		}
	}
}
