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

func TestPathCopyEmpty(t *testing.T) {
	p1 := Path{}
	p2 := p1.Copy()

	if !reflect.DeepEqual(p1, p2) {
		t.Error("copy of empty path did not produce another empty path:", p2)
	}
}

func TestPathPath(t *testing.T) {
	p1 := Path{}
	p1.LineTo(Point{1, 1})

	p2 := p1.Path()

	if !reflect.DeepEqual(p1, p2) {
		t.Error("invalid path copy:", p1, p2)
	}
}

func TestPathLastPointEmpty(t *testing.T) {
	p1 := Path{}
	pt := p1.LastPoint()

	if pt != (Point{}) {
		t.Error("last point of an empty path is not the zero-value point:", pt)
	}
}

func TestPathLastPointMoveTo(t *testing.T) {
	p1 := Path{}
	p1.MoveTo(Point{1, 1})
	pt := p1.LastPoint()

	if pt != (Point{1, 1}) {
		t.Error("invalid last point returned after move-to operation:", p1, pt)
	}
}

func TestPathLastPointLineTo(t *testing.T) {
	p1 := Path{}
	p1.LineTo(Point{1, 1})
	pt := p1.LastPoint()

	if pt != (Point{1, 1}) {
		t.Error("invalid last point returned after line-to operation:", p1, pt)
	}
}

func TestPathLastPointQuadCurveTo(t *testing.T) {
	p1 := Path{}
	p1.QuadCurveTo(Point{0.5, 0.5}, Point{1, 1})
	pt := p1.LastPoint()

	if pt != (Point{1, 1}) {
		t.Error("invalid last point returned after quad-curve-to operation:", p1, pt)
	}
}

func TestPathLastPointCubicCurveTo(t *testing.T) {
	p1 := Path{}
	p1.CubicCurveTo(Point{0.25, 0.25}, Point{0.75, 0.75}, Point{1, 1})
	pt := p1.LastPoint()

	if pt != (Point{1, 1}) {
		t.Error("invalid last point returned after cubic-curve-to operation:", p1, pt)
	}
}

func TestPathEnsureMoveToAfterClosePath(t *testing.T) {
	p := Path{}
	p.LineTo(Point{1, 1})
	p.Close()
	p.LineTo(Point{1, 0})

	if !reflect.DeepEqual(p, Path{
		Elements: []PathElement{
			{
				Type:   MoveTo,
				Points: [...]Point{{}, {}, {}},
			},
			{
				Type:   LineTo,
				Points: [...]Point{{1, 1}, {}, {}},
			},
			{
				Type:   ClosePath,
				Points: [...]Point{{}, {}, {}},
			},
			{
				Type:   MoveTo,
				Points: [...]Point{{1, 1}, {}, {}},
			},
			{
				Type:   LineTo,
				Points: [...]Point{{1, 0}, {}, {}},
			},
		},
	}) {
		t.Errorf("invalid path built by [line-to, close-path, line-to] operations: %#v", p)
	}
}

func TestAppendPath(t *testing.T) {
	p1 := Path{}
	p1.LineTo(Point{1, 1})

	p2 := Path{}
	p2.LineTo(Point{2, 2})

	p1 = AppendPath(p1, p2)

	if !reflect.DeepEqual(p1, Path{
		Elements: []PathElement{
			{
				Type:   MoveTo,
				Points: [...]Point{{}, {}, {}},
			},
			{
				Type:   LineTo,
				Points: [...]Point{{1, 1}, {}, {}},
			},
			{
				Type:   MoveTo,
				Points: [...]Point{{}, {}, {}},
			},
			{
				Type:   LineTo,
				Points: [...]Point{{2, 2}, {}, {}},
			},
		},
	}) {
		t.Errorf("invalid path build by appending a path to another: %#v", p1)
	}
}
