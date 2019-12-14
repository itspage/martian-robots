package martian

type Coordinates struct {
	X int
	Y int
}

type Grid struct {
	width      int
	height     int
	lostCoords []Coordinates
}

func NewGrid(width, height int) (*Grid, error) {
	if width > 50 {
		return nil, ErrMaxGridWidth
	}
	if width < 1 {
		return nil, ErrMinGridWidth
	}
	if height > 50 {
		return nil, ErrMaxGridHeight
	}
	if height < 1 {
		return nil, ErrMinGridHeight
	}
	return &Grid{
		width:  width,
		height: height,
	}, nil
}
