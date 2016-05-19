package geom

import (
	"reflect"
	"testing"
)

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
		t.Error("calling SetOrigin did not set the rectangle's origin to the expected value:", r)
	}
}

func TestRectSetSize(t *testing.T) {
	r := Rect{}
	r.SetSize(Size{1, 1})

	if r != (Rect{0, 0, 1, 1}) {
		t.Error("calling SetSize did not set the rectangle's size to the expected value:", r)
	}
}

func TestRectOrigin(t *testing.T) {
	r := Rect{1, 1, 0, 0}
	p := r.Origin()

	if p != (Point{1, 1}) {
		t.Error("invalid coordinates returned by the Origin method:", p)
	}
}

func TestRectCenter(t *testing.T) {
	r := Rect{0, 0, 1, 1}
	p := r.Center()

	if p != (Point{0.5, 0.5}) {
		t.Error("invalid coordinates returned by the Center method:", p)
	}
}

func TestRectTip(t *testing.T) {
	r := Rect{0, 0, 1, 1}
	p := r.Tip()

	if p != (Point{1, 1}) {
		t.Error("invalid coordinates returned by the Tip method:", p)
	}
}

func TestRectSize(t *testing.T) {
	r := Rect{0, 0, 1, 1}
	s := r.Size()

	if s != (Size{1, 1}) {
		t.Error("invalid dimensions returned by the Size method:", s)
	}
}

func TestRectAbs(t *testing.T) {
	r := Rect{0, 0, -1, -1}
	a := r.Abs()

	if a != (Rect{-1, -1, 1, 1}) {
		t.Error("invalid rectangle returned by the Abs method:", a)
	}
}

func TestRectIntersectNone(t *testing.T) {
	r1 := Rect{0, 0, 1, 1}
	r2 := Rect{2, 2, 1, 1}
	r3 := r1.Intersect(r2)

	if r3 != (Rect{}) {
		t.Error("invalid non-zero rectangle returned by call to Intersect on non-intersecting rectangles:", r3)
	}
}

func TestRectIntersectSome(t *testing.T) {
	r1 := Rect{0, 0, 1, 1}
	r2 := Rect{0.5, 0.5, 1, 1}
	r3 := r1.Intersect(r2)

	if r3 != (Rect{0.5, 0.5, 0.5, 0.5}) {
		t.Error("invalid non-zero rectangle returned by call to Intersect on partially intersecting rectangles:", r3)
	}
}

func TestRectIntersectFull(t *testing.T) {
	r1 := Rect{0, 0, 1, 1}
	r2 := Rect{0.25, 0.25, 0.5, 0.5}
	r3 := r1.Intersect(r2)

	if r3 != (Rect{0.25, 0.25, 0.5, 0.5}) {
		t.Error("invalid non-zero rectangle returned by call to Intersect on fully intersecting rectangles:", r3)
	}
}

func TestRectMergeIntersectNone(t *testing.T) {
	r1 := Rect{0, 0, 1, 1}
	r2 := Rect{2, 2, 1, 1}
	r3 := r1.Merge(r2)

	if r3 != (Rect{0, 0, 3, 3}) {
		t.Error("invalid non-zero rectangle returned by call to Merge on non-intersecting rectangles:", r3)
	}
}

func TestRectMergeIntersectSome(t *testing.T) {
	r1 := Rect{0, 0, 1, 1}
	r2 := Rect{0.5, 0.5, 1, 1}
	r3 := r1.Merge(r2)

	if r3 != (Rect{0, 0, 1.5, 1.5}) {
		t.Error("invalid non-zero rectangle returned by call to Merge on partially intersecting rectangles:", r3)
	}
}

func TestRectMergeIntersectFull(t *testing.T) {
	r1 := Rect{0, 0, 1, 1}
	r2 := Rect{0.25, 0.25, 0.5, 0.5}
	r3 := r1.Merge(r2)

	if r3 != (Rect{0, 0, 1, 1}) {
		t.Error("invalid non-zero rectangle returned by call to Merge on fully intersecting rectangles:", r3)
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

func TestRectString(t *testing.T) {
	r := Rect{0.5, 0.25, 1, 2}
	s := r.String()

	if s != "{ 0.5, 0.25, 1, 2 }" {
		t.Errorf("invalid string representation of a rectangle: %#v %s", r, s)
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

func TestCenterRect(t *testing.T) {
	r1 := Rect{0, 0, 10, 10}
	r2 := Rect{0, 0, 2, 2}
	r3 := CenterRect(r1, r2)

	if r3 != (Rect{4, 4, 2, 2}) {
		t.Error("invalid centered rectangle:", r3)
	}
}

func TestMergeRect(t *testing.T) {
	tests := []struct {
		key  string
		rect Rect
		in   []Rect
		out  []Rect
	}{
		{
			key:  "empty slice",
			rect: Rect{0, 0, 1, 1},
			in:   nil,
			out:  []Rect{{0, 0, 1, 1}},
		},
		{
			key:  "non-empty slice, not contained, no intersections",
			rect: Rect{0, 0, 1, 1},
			in:   []Rect{{0, 2, 1, 1}},
			out:  []Rect{{0, 2, 1, 1}, {0, 0, 1, 1}},
		},
		{
			key:  "non-empty slice, contained",
			rect: Rect{0, 0, 1, 1},
			in:   []Rect{{0, 0, 2, 2}},
			out:  []Rect{{0, 0, 2, 2}},
		},
		{
			key:  "non-empty slice, intersects",
			rect: Rect{1, 1, 2, 2},
			in:   []Rect{{0, 0, 2, 2}},
			out:  []Rect{{0, 0, 3, 3}},
		},
		{
			key:  "non-empty slice, re-intersects",
			rect: Rect{1, 1, 2, 2},
			in:   []Rect{{0, 0, 2, 2}, {2.5, 2.5, 1, 1}},
			out:  []Rect{{0, 0, 3.5, 3.5}},
		},
	}

	for _, test := range tests {
		if list := MergeRect(test.in, test.rect); !reflect.DeepEqual(list, test.out) {
			t.Errorf("MergeReect: %s: %#v", test.key, list)
		}
	}
}
