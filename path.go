package geom

type PathElementType int

const (
	MoveTo PathElementType = iota
	LineTo
	QuadCurveTo
	CubicCurveTo
	ClosePath
)

type PathElement struct {
	Type   PathElementType
	Points [3]Point
}

type Path struct {
	Elements []PathElement
}

type Shape interface {
	Path() Path
}

func MakePath(cap int) Path {
	return Path{
		Elements: make([]PathElement, 0, cap),
	}
}

func AppendPath(path Path, other Path) Path {
	return Path{Elements: append(path.Elements, other.Elements...)}
}

func AppendRect(path Path, rect Rect) Path {
	x0 := rect.X
	y0 := rect.Y
	x1 := rect.X + rect.W
	y1 := rect.Y + rect.H
	return AppendPolygon(path, Point{x0, y0}, Point{x1, y0}, Point{x1, y1}, Point{x0, y1})
}

func AppendPolygon(path Path, points ...Point) Path {
	if len(points) != 0 {
		path.MoveTo(points[0])

		for _, p := range points[1:] {
			path.LineTo(p)
		}

		path.Close()
	}

	return path
}

func (path *Path) MoveTo(p Point) {
	path.append(PathElement{
		Type:   MoveTo,
		Points: [...]Point{p, {}, {}},
	})
}

func (path *Path) LineTo(p Point) {
	path.ensureStartWithMoveTo()
	path.append(PathElement{
		Type:   LineTo,
		Points: [...]Point{p, {}, {}},
	})
}

func (path *Path) QuadCurveTo(c Point, p Point) {
	path.ensureStartWithMoveTo()
	path.append(PathElement{
		Type:   QuadCurveTo,
		Points: [...]Point{c, p, {}},
	})
}

func (path *Path) CubicCurveTo(c1 Point, c2 Point, p Point) {
	path.ensureStartWithMoveTo()
	path.append(PathElement{
		Type:   CubicCurveTo,
		Points: [...]Point{c1, c2, p},
	})
}

func (path *Path) Close() {
	path.append(PathElement{
		Type: ClosePath,
	})
}

func (path *Path) Clear() {
	path.Elements = path.Elements[:0]
}

func (path *Path) Copy() Path {
	if path.Empty() {
		return Path{}
	}

	p := Path{
		Elements: make([]PathElement, len(path.Elements)),
	}

	copy(p.Elements, path.Elements)
	return p
}

func (path *Path) LastPoint() Point {
	return path.lastPointAt(len(path.Elements) - 1)
}

func (path *Path) lastPointAt(n int) Point {
	if n < 0 {
		return Point{}
	}

	switch e := path.Elements[n-1]; e.Type {
	case MoveTo, LineTo:
		return e.Points[0]

	case QuadCurveTo:
		return e.Points[1]

	case CubicCurveTo:
		return e.Points[2]

	default:
		return path.lastPointAt(n - 1)
	}
}

func (path *Path) Empty() bool {
	return len(path.Elements) == 0
}

func (path Path) Path() Path {
	return path.Copy()
}

func (path *Path) append(e PathElement) {
	path.Elements = append(path.Elements, e)
}

func (path *Path) ensureStartWithMoveTo() {
	if n := len(path.Elements); n == 0 {
		// Empty paths must start with a MoveTo element.
		path.MoveTo(Point{})

	} else if path.Elements[n-1].Type == ClosePath {
		// When a subpath is closed the path that contains more elements must
		// be restarted with a MoveTo as well.
		path.MoveTo(path.LastPoint())
	}
}
