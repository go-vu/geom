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

func (self *Path) MoveTo(p Point) {
	self.append(PathElement{
		Type:   MoveTo,
		Points: [...]Point{p, {}, {}},
	})
}

func (self *Path) LineTo(p Point) {
	self.ensureStartWithMoveTo()
	self.append(PathElement{
		Type:   LineTo,
		Points: [...]Point{p, {}, {}},
	})
}

func (self *Path) QuadCurveTo(c Point, p Point) {
	self.ensureStartWithMoveTo()
	self.append(PathElement{
		Type:   QuadCurveTo,
		Points: [...]Point{c, p, {}},
	})
}

func (self *Path) CubicCurveTo(c1 Point, c2 Point, p Point) {
	self.ensureStartWithMoveTo()
	self.append(PathElement{
		Type:   CubicCurveTo,
		Points: [...]Point{c1, c2, p},
	})
}

func (self *Path) Close() {
	self.append(PathElement{
		Type: ClosePath,
	})
}

func (self *Path) Clear() {
	self.Elements = self.Elements[:0]
}

func (self *Path) Copy() Path {
	if self.Empty() {
		return Path{}
	}

	p := Path{
		Elements: make([]PathElement, len(self.Elements)),
	}

	copy(p.Elements, self.Elements)
	return p
}

func (self *Path) LastPoint() Point {
	return self.lastPointAt(len(self.Elements) - 1)
}

func (self *Path) lastPointAt(n int) Point {
	if n < 0 {
		return Point{}
	}

	switch e := self.Elements[n-1]; e.Type {
	case MoveTo, LineTo:
		return e.Points[0]

	case QuadCurveTo:
		return e.Points[1]

	case CubicCurveTo:
		return e.Points[2]

	default:
		return self.lastPointAt(n - 1)
	}
}

func (self *Path) Empty() bool {
	return len(self.Elements) == 0
}

func (self Path) Path() Path {
	return self.Copy()
}

func (self *Path) append(e PathElement) {
	self.Elements = append(self.Elements, e)
}

func (self *Path) ensureStartWithMoveTo() {
	if n := len(self.Elements); n == 0 {
		// Empty paths must start with a MoveTo element.
		self.MoveTo(Point{})

	} else if self.Elements[n-1].Type == ClosePath {
		// When a subpath is closed the path that contains more elements must
		// be restarted with a MoveTo as well.
		self.MoveTo(self.LastPoint())
	}
}
