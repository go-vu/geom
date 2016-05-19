package geom

import "testing"

func TestMakeMargin(t *testing.T) {
	if m := MakeMargin(0.5); m != (Margin{
		Top:    0.5,
		Bottom: 0.5,
		Left:   0.5,
		Right:  0.5,
	}) {
		t.Error("MakeMargin returned an invalid value:", m)
	}
}

func TestMarginTopLeftZero(t *testing.T) {
	m := Margin{}
	p := m.TopLeft(Point{
		X: 1,
		Y: 2,
	})

	if p != (Point{X: 1, Y: 2}) {
		t.Error("margin applied invalid transformation to top-left point:", m, p)
	}
}

func TestMarginTopLeftNonZero(t *testing.T) {
	m := MakeMargin(0.5)
	p := m.TopLeft(Point{
		X: 1,
		Y: 2,
	})

	if p != (Point{X: 1.5, Y: 2.5}) {
		t.Error("margin applied invalid transformation to top-left point:", m, p)
	}
}

func TestMarginBottomRightZero(t *testing.T) {
	m := Margin{}
	p := m.BottomRight(Point{
		X: 1,
		Y: 2,
	})

	if p != (Point{X: 1, Y: 2}) {
		t.Error("margin applied invalid transformation to bottom-right point:", m, p)
	}
}

func TestMarginBottomRightNonZero(t *testing.T) {
	m := MakeMargin(0.5)
	p := m.BottomRight(Point{
		X: 1,
		Y: 2,
	})

	if p != (Point{X: 0.5, Y: 1.5}) {
		t.Error("margin applied invalid transformation to bottom-right point:", m, p)
	}
}

func TestMarginGrowRect(t *testing.T) {
	m := MakeMargin(0.5)
	r := m.GrowRect(Rect{W: 2, H: 2})

	if r != (Rect{
		X: -0.5,
		Y: -0.5,
		W: 3,
		H: 3,
	}) {
		t.Error("margin growing rectangle produced an invalid value:", m, r)
	}
}

func TestMarginShrinkRect(t *testing.T) {
	m := MakeMargin(0.5)
	r := m.ShrinkRect(Rect{W: 2, H: 2})

	if r != (Rect{
		X: 0.5,
		Y: 0.5,
		W: 1,
		H: 1,
	}) {
		t.Error("margin shrinking rectangle produced an invalid value:", m, r)
	}
}

func TestMarginShrinkRectEmpty(t *testing.T) {
	m := MakeMargin(2)
	r := m.ShrinkRect(Rect{W: 2, H: 2})

	if r != (Rect{X: 1, Y: 1}) {
		t.Error("margin shrinking rectangle produced an invalid value:", m, r)
	}

}

func TestMarginWidth(t *testing.T) {
	m := MakeMargin(0.25)
	w := m.Width()

	if w != 0.5 {
		t.Error("invalid margin width:", m, w)
	}
}

func TestMarginHeight(t *testing.T) {
	m := MakeMargin(0.25)
	h := m.Height()

	if h != 0.5 {
		t.Error("invalid margin height:", m, h)
	}
}

func TestMarginSize(t *testing.T) {
	m := MakeMargin(0.25)
	s := m.Size()

	if s != (Size{W: 0.5, H: 0.5}) {
		t.Error("invalid margin size:", m, s)
	}
}

func TestMarginString(t *testing.T) {
	m := MakeMargin(0.25)
	s := m.String()

	if s != "margin { top = 0.25, bottom = 0.25, left = 0.25, right = 0.25 }" {
		t.Errorf("invalid string representation of a margin value: %#v %s", m, s)
	}
}
