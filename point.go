package geom

import "fmt"

type Point struct {
	X float64
	Y float64
}

func (self Point) Zero() bool {
	return self.X == 0 && self.Y == 0
}

func (self Point) String() string {
	return fmt.Sprintf("(%.6g, %.6g)", self.X, self.Y)
}

func (self Point) WithOrigin(origin Point) Point {
	return Point{
		X: self.X - origin.X,
		Y: self.Y - origin.Y,
	}
}
