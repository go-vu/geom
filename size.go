package geom

import "fmt"

type Size struct {
	W float64
	H float64
}

func (s Size) Zero() bool {
	return s.W == 0 && s.H == 0
}

func (s Size) Empty() bool {
	return s.W == 0 || s.H == 0
}

func (s Size) Area() float64 {
	return s.W * s.H
}

func (s Size) Ratio() float64 {
	return s.W / s.H
}

func (s Size) String() string {
	return fmt.Sprintf("[%.6g, %.6g]", s.W, s.H)
}
