package geom

import "testing"

func TestMakeRect(t *testing.T) {
	test := func(p Point, s Size, r Rect) {
		if x := MakeRect(p, s); x != r {
			t.Error("MakeRect:", x, "!=", r)
		}
	}

	test(Point{}, Size{}, Rect{})
	test(Point{1, 2}, Size{3, 4}, Rect{1, 2, 3, 4})
}

func TestRectSetOrigin(t *testing.T) {
	r := Rect{}
	r.SetOrigin(Point{1, 1})

	if r != (Rect{1, 1, 0, 0}) {
		t.Error(r)
	}
}

func TestRectSetSize(t *testing.T) {
	r := Rect{}
	r.SetSize(Size{1, 1})

	if r != (Rect{0, 0, 1, 1}) {
		t.Error(r)
	}
}

func TestRectOrigin(t *testing.T) {
	r := Rect{1, 1, 0, 0}

	if r.Origin() != (Point{1, 1}) {
		t.Error(r)
	}
}

func TestRectSize(t *testing.T) {
	r := Rect{0, 0, 1, 1}

	if r.Size() != (Size{1, 1}) {
		t.Error(r)
	}
}

func TestRectZero(t *testing.T) {
	if !(Rect{}).Zero() {
		t.Error("zero Rect is non-zero")
	}

	if (Rect{1, 0, 0, 0}).Zero() {
		t.Error("non-zero Rect is zero")
	}

	if (Rect{0, 1, 0, 0}).Zero() {
		t.Error("non-zero Rect is zero")
	}

	if (Rect{0, 0, 1, 0}).Zero() {
		t.Error("non-zero Rect is zero")
	}

	if (Rect{0, 0, 0, 1}).Zero() {
		t.Error("non-zero Rect is zero")
	}
}

func TestRectEmpty(t *testing.T) {
	if !(Rect{}).Empty() {
		t.Error("empty Rect is non-empty")
	}

	if !(Rect{1, 0, 0, 0}).Empty() {
		t.Error("empty Rect is non-empty")
	}

	if !(Rect{0, 1, 0, 0}).Empty() {
		t.Error("empty Rect is non-empty")
	}

	if !(Rect{0, 0, 1, 0}).Empty() {
		t.Error("empty Rect is non-empty")
	}

	if !(Rect{0, 0, 0, 1}).Empty() {
		t.Error("empty Rect is non-empty")
	}

	if (Rect{0, 0, 1, 1}).Empty() {
		t.Error("non-empty Rect is empty")
	}
}

func TestRectArea(t *testing.T) {
	if (Rect{}).Area() != 0 {
		t.Error("empty Rect area is not zero")
	}

	if (Rect{0, 0, 1, 0}).Area() != 0 {
		t.Error("empty Rect area is not zero")
	}

	if (Rect{0, 0, 0, 1}).Area() != 0 {
		t.Error("empty Rect area is not zero")
	}

	if (Rect{0, 0, 1, 1}).Area() != 1 {
		t.Error("non-empty Rect area is invalid")
	}
}

func TestRectContainsPoint(t *testing.T) {
	if (Rect{}).ContainsPoint(Point{}) {
		t.Error("empty Rect can't contain points")
	}

	if !(Rect{0, 0, 1, 1}).ContainsPoint(Point{}) {
		t.Error("non-empty Rect should contain (0, 0)")
	}
}

func TestRectContainsRect(t *testing.T) {
	if !(Rect{}).ContainsRect(Rect{}) {
		t.Error("empty Rect can't contain another Rect")
	}

	if !(Rect{0, 0, 1, 1}).ContainsRect(Rect{0.25, 0.25, 0.5, 0.5}) {
		t.Error("non-empty Rect should contain another Rect")
	}
}

func TestRectPath(t *testing.T) {
	r := Rect{0, 0, 1, 1}
	p := r.Path()

	if len(p.Elements) != 5 {
		t.Error("invalid rectangle path:", p)
		return
	}

	if p.Elements[0] != (PathElement{
		Type:   MoveTo,
		Points: [...]Point{{}, {}, {}},
	}) {
		t.Error("invalid rectangle path element:", p.Elements[0])
	}

	if p.Elements[1] != (PathElement{
		Type:   LineTo,
		Points: [...]Point{{1, 0}, {}, {}},
	}) {
		t.Error("invalid rectangle path element:", p.Elements[1])
	}

	if p.Elements[2] != (PathElement{
		Type:   LineTo,
		Points: [...]Point{{1, 1}, {}, {}},
	}) {
		t.Error("invalid rectangle path element:", p.Elements[2])
	}

	if p.Elements[3] != (PathElement{
		Type:   LineTo,
		Points: [...]Point{{0, 1}, {}, {}},
	}) {
		t.Error("invalid rectangle path element:", p.Elements[3])
	}

	if p.Elements[4] != (PathElement{
		Type:   ClosePath,
		Points: [...]Point{{}, {}, {}},
	}) {
		t.Error("invalid rectangle path element:", p.Elements[4])
	}
}
