package geom

import (
	"reflect"
	"testing"
)

func TestMakePath(t *testing.T) {
	p := MakePath(10)

	if c := cap(p.Elements); c < 10 {
		t.Error("invalid path capacity:", c)
	}
}

func TestPathMoveTo(t *testing.T) {
	p := Path{}

	p.MoveTo(Point{1, 1})

	if len(p.Elements) != 1 {
		t.Error("invalid path elements:", p.Elements)
		return
	}

	if p.Elements[0] != (PathElement{
		Type:   MoveTo,
		Points: [...]Point{{1, 1}, {}, {}},
	}) {
		t.Error("invalid path elements:", p.Elements)
	}
}

func TestPathLineTo(t *testing.T) {
	p := Path{}

	p.LineTo(Point{1, 1})

	if len(p.Elements) != 2 {
		t.Error("invalid path elements:", p.Elements)
		return
	}

	if p.Elements[0] != (PathElement{
		Type:   MoveTo,
		Points: [...]Point{{}, {}, {}},
	}) {
		t.Error("invalid path elements:", p.Elements)
	}

	if p.Elements[1] != (PathElement{
		Type:   LineTo,
		Points: [...]Point{{1, 1}, {}, {}},
	}) {
		t.Error("invalid path elements:", p.Elements)
	}
}

func TestPathQuadCurveTo(t *testing.T) {
	p := Path{}

	p.QuadCurveTo(Point{}, Point{1, 1})

	if len(p.Elements) != 2 {
		t.Error("invalid path elements:", p.Elements)
		return
	}

	if p.Elements[0] != (PathElement{
		Type:   MoveTo,
		Points: [...]Point{{}, {}, {}},
	}) {
		t.Error("invalid path elements:", p.Elements)
	}

	if p.Elements[1] != (PathElement{
		Type:   QuadCurveTo,
		Points: [...]Point{{}, {1, 1}, {}},
	}) {
		t.Error("invalid path elements:", p.Elements)
	}
}

func TestPathCubicCurveTo(t *testing.T) {
	p := Path{}

	p.CubicCurveTo(Point{}, Point{1, 1}, Point{2, 2})

	if len(p.Elements) != 2 {
		t.Error("invalid path elements:", p.Elements)
		return
	}

	if p.Elements[0] != (PathElement{
		Type:   MoveTo,
		Points: [...]Point{{}, {}, {}},
	}) {
		t.Error("invalid path elements:", p.Elements)
	}

	if p.Elements[1] != (PathElement{
		Type:   CubicCurveTo,
		Points: [...]Point{{}, {1, 1}, {2, 2}},
	}) {
		t.Error("invalid path elements:", p.Elements)
	}
}

func TestPathClear(t *testing.T) {
	p := Path{}

	p.MoveTo(Point{1, 1})

	if p.Empty() {
		t.Error("invalid empty path:", p)
		return
	}

	p.Clear()

	if !p.Empty() {
		t.Error("invalid non-empty path:", p)
		return
	}
}

func TestPathCopy(t *testing.T) {
	p1 := Path{}
	p1.LineTo(Point{1, 1})

	p2 := p1.Copy()

	if !reflect.DeepEqual(p1, p2) {
		t.Error("invalid path copy:", p1, p2)
	}
}
