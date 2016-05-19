package geom

import "fmt"

// The Margin type represent 2D margins of a rectangle area.
type Margin struct {
	Top    float64
	Bottom float64
	Left   float64
	Right  float64
}

// MakeMargin takes a numeric value as argument that defines the top, left,
// right, and bottom components of the returned margin value.
func MakeMargin(m float64) Margin {
	return Margin{
		Top:    m,
		Bottom: m,
		Left:   m,
		Right:  m,
	}
}

// TopLeft takes a point as argument that would be the top-left corner of a
// rectangle, applies the margin and return the modified coordinates.
func (m Margin) TopLeft(p Point) Point {
	return Point{
		X: p.X + m.Left,
		Y: p.Y + m.Top,
	}
}

// BottomRight takes a point as argument that would be the bottom-right corner
// of a rectangle, applies the margin and return the modified corrdinates.
func (m Margin) BottomRight(p Point) Point {
	return Point{
		X: p.X - m.Right,
		Y: p.Y - m.Bottom,
	}
}

// GrowRect grows the given rectangle by applying the margin and returns the
// modified rectangle.
func (m Margin) GrowRect(r Rect) Rect {
	return Rect{
		X: r.X - m.Left,
		Y: r.Y - m.Top,
		W: r.W + m.Width(),
		H: r.H + m.Height(),
	}
}

// ShrinkRect shrinks the given rectangle by applying the margin and returns the
// modified rectangle.
func (m Margin) ShrinkRect(r Rect) Rect {
	s := Rect{
		X: r.X + m.Left,
		Y: r.Y + m.Top,
		W: r.W - m.Width(),
		H: r.H - m.Height(),
	}

	if s.W < 0 {
		s.X = r.X + (r.W / 2)
		s.W = 0
	}

	if s.H < 0 {
		s.Y = r.Y + (r.H / 2)
		s.H = 0
	}

	return s
}

// Width returns the sum of the left and right values of the given margin.
func (m Margin) Width() float64 {
	return m.Left + m.Right
}

// Height returns the sum of the top and bottom values of the given margin.
func (m Margin) Height() float64 {
	return m.Top + m.Bottom
}

// Size returns the combined width and height of the margin as a size value.
func (m Margin) Size() Size {
	return Size{
		W: m.Width(),
		H: m.Height(),
	}
}

// The String method returns a string representation of the margin value.
func (m Margin) String() string {
	return fmt.Sprintf("margin { top = %g, bottom = %g, left = %g, right = %g }", m.Top, m.Bottom, m.Left, m.Right)
}
