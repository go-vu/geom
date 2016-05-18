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

func (m Margin) TopLeft(p Point) Point {
	return Point{
		X: p.X + m.Left,
		Y: p.Y + m.Top,
	}
}

func (m Margin) BottomRight(p Point) Point {
	return Point{
		X: p.X - m.Right,
		Y: p.Y - m.Bottom,
	}
}

func (m Margin) GrowRect(r Rect) Rect {
	return Rect{
		X: r.X - m.Left,
		Y: r.Y - m.Top,
		W: r.W + m.Width(),
		H: r.H + m.Height(),
	}
}

func (m Margin) GrowSize(s Size) Size {
	return Size{
		W: s.W + m.Width(),
		H: s.H + m.Height(),
	}
}

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

func (m Margin) ShrinkSize(s Size) Size {
	return Size{
		W: math.Max(0, s.W-m.Width()),
		H: math.Max(0, s.H-m.Height()),
	}
}

func (m Margin) Width() float64 {
	return m.Left + m.Right
}

func (m Margin) Height() float64 {
	return m.Top + m.Bottom
}

func (m Margin) Size() Size {
	return Size{
		W: m.Width(),
		H: m.Height(),
	}
}

func (m Margin) String() string {
	return fmt.Sprintf("margins { top = %g, bottom = %g, left = %g, right = %g }", m.Top, m.Bottom, m.Left, m.Right)
}
