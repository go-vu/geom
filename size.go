package geom

import "fmt"

type Size struct {
	W float64
	H float64
}

func (self Size) Zero() bool {
	return self.W == 0 && self.H == 0
}

func (self Size) Empty() bool {
	return self.W == 0 || self.H == 0
}

func (self Size) Area() float64 {
	return self.W * self.H
}

func (self Size) Ratio() float64 {
	return self.W / self.H
}

func (self Size) String() string {
	return fmt.Sprintf("[%.6g, %.6g]", self.W, self.H)
}
