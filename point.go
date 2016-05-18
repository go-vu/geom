package geom

import "fmt"

// The Point type represents 2D coordinates.
type Point struct {
	X float64
	Y float64
}

// Zero checks if the receiver has the zero-value (both the x and y components
// are zero).
func (p Point) Zero() bool {
	return p.X == 0 && p.Y == 0
}

// The String method returns a human-readable representation of the point value.
func (p Point) String() string {
	return fmt.Sprintf("(%.6g, %.6g)", p.X, p.Y)
}

// WithOrigin translate the receiver to a coordinate system with origin given as
// argument and returns the modified point.
func (p Point) WithOrigin(origin Point) Point {
	return Point{
		X: p.X - origin.X,
		Y: p.Y - origin.Y,
	}
}
