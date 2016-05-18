package geom

import (
	"fmt"
	"math"
)

type Margin struct {
	Top    float64
	Bottom float64
	Left   float64
	Right  float64
}

func MakeMargin(dist float64) Margin {
	return Margin{
		Top:    dist,
		Bottom: dist,
		Left:   dist,
		Right:  dist,
	}
}

func (self Margin) TopLeft(point Point) Point {
	return Point{
		X: point.X + self.Left,
		Y: point.Y + self.Top,
	}
}

func (self Margin) BottomRight(point Point) Point {
	return Point{
		X: point.X - self.Right,
		Y: point.Y - self.Bottom,
	}
}

func (self Margin) GrowRect(rect Rect) Rect {
	return Rect{
		X: rect.X - self.Left,
		Y: rect.Y - self.Top,
		W: rect.W + self.Width(),
		H: rect.H + self.Height(),
	}
}

func (self Margin) GrowSize(size Size) Size {
	return Size{
		W: size.W + self.Width(),
		H: size.H + self.Height(),
	}
}

func (self Margin) ShrinkRect(rect Rect) Rect {
	r := Rect{
		X: rect.X + self.Left,
		Y: rect.Y + self.Top,
		W: rect.W - self.Width(),
		H: rect.H - self.Height(),
	}

	if r.W < 0 {
		r.X = rect.X + (rect.W / 2)
		r.W = 0
	}

	if r.H < 0 {
		r.Y = rect.Y + (rect.H / 2)
		r.H = 0
	}

	return r
}

func (self Margin) ShrinkSize(size Size) Size {
	return Size{
		W: math.Max(0, size.W-self.Width()),
		H: math.Max(0, size.H-self.Height()),
	}
}

func (self Margin) Width() float64 {
	return self.Left + self.Right
}

func (self Margin) Height() float64 {
	return self.Top + self.Bottom
}

func (self Margin) Size() Size {
	return Size{
		W: self.Width(),
		H: self.Height(),
	}
}

func (self Margin) String() string {
	return fmt.Sprintf("margins { top = %g, bottom = %g, left = %g, right = %g }", self.Top, self.Bottom, self.Left, self.Right)
}
