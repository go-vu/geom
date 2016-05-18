package geom

import "fmt"

// THe Size type represents 2D sizes that are composed of a width and height.
type Size struct {
	W float64
	H float64
}

// Zero checks the size value and returns true if it is the zero-value, false
// otherwise.
func (s Size) Zero() bool {
	return s.W == 0 && s.H == 0
}

// Empty checks whether the size of empty, which is true when either the width
// or height are zero.
func (s Size) Empty() bool {
	return s.W == 0 || s.H == 0
}

// Area computes and returns the area of the size (width x height).
func (s Size) Area() float64 {
	return s.W * s.H
}

// Ratio computes and returns the ratio of the size (width / height).
func (s Size) Ratio() float64 {
	return s.W / s.H
}

// The String method returns a human-representation of the size value.
func (s Size) String() string {
	return fmt.Sprintf("[%.6g, %.6g]", s.W, s.H)
}
