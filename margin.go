package geom

import (
	"fmt"
	"math"
)

// The Margin type represent 2D margins of a rectangle area.
type Margin struct {
	Top    float64
	Bottom float64
	Left   float64
	Right  float64
}

// Makes a new Margin value where every side of a rectangle would be applied the
// margin value passed as argument.
func MakeMargin(m float64) Margin {
	return Margin{
		Top:    m,
		Bottom: m,
		Left:   m,
		Right:  m,
	}
}

// Given a point that would be the top-left corner of a rectangle, applies the
// margin and return the modified coordinates.
func (m Margin) TopLeft(p Point) Point {
	return Point{
		X: p.X + m.Left,
		Y: p.Y + m.Top,
	}
}

// Given a point that would be the bottom-right corner of a rectangle, applies
// the margin and return the modified corrdinates.
func (m Margin) BottomRight(p Point) Point {
	return Point{
		X: p.X - m.Right,
		Y: p.Y - m.Bottom,
	}
}

// Grows the given rectangle by applying the margin and returns the modified
// rectangle.
func (m Margin) GrowRect(r Rect) Rect {
	return Rect{
		X: r.X - m.Left,
		Y: r.Y - m.Top,
		W: r.W + m.Width(),
		H: r.H + m.Height(),
	}
}

// Grows a size value by applying the margin and returns the modified size.
func (m Margin) GrowSize(s Size) Size {
	return Size{
		W: s.W + m.Width(),
		H: s.H + m.Height(),
	}
}

// Shrinks the given rectangle by applying the margin and returns the modified
// rectangle.
func (m Margin) ShrinkRect(r Rect) Rect {
	s := Rect{
		X: r.X + m.Left,
		Y: r.Y + m.Top,
		W: r.W - m.Width(),
		H: r.H - m.Height(),
	}

	if s.W < 0 {
		r.X = r.X + (r.W / 2)
		s.W = 0
	}

	if s.H < 0 {
		r.Y = r.Y + (r.H / 2)
		s.H = 0
	}

	return s
}

// Shrinks the given size by applying the margin and returns the modified size.
func (m Margin) ShrinkSize(s Size) Size {
	return Size{
		W: math.Max(0, s.W-m.Width()),
		H: math.Max(0, s.H-m.Height()),
	}
}

// Returns the sum of the left and right values of the given margin.
func (m Margin) Width() float64 {
	return m.Left + m.Right
}

// Returns the sum of the top and bottom values of the given margin.
func (m Margin) Height() float64 {
	return m.Top + m.Bottom
}

// Returns the combined width and height of the margin as a size value.
func (m Margin) Size() Size {
	return Size{
		W: m.Width(),
		H: m.Height(),
	}
}

// Returns a string representation of the margin value.
func (m Margin) String() string {
	return fmt.Sprintf("margins { top = %g, bottom = %g, left = %g, right = %g }", m.Top, m.Bottom, m.Left, m.Right)
}
