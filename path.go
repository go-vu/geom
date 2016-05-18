package geom

// PathElementType is an enumeration representing the different kinds of path
// elements supported by 2D paths.
type PathElementType int

const (
	// MoveTo is used for path elements that move the path position to a new
	// location.
	MoveTo PathElementType = iota

	// LineTo is used for path elements that move the path position to a new
	// location while drawin a straight line.
	LineTo

	// QuadCurveTo is used for path elements that draw a quadriatic cruve.
	QuadCurveTo

	// CubicCurveTo is used for path elements that draw a cubic curve.
	CubicCurveTo

	// ClosePath is used for path elements that terminate drawing of a sub-path.
	ClosePath
)

// A PathElement represents a single step in a 2D path, it's composed of a type
// that must be one of the constant values of PathElementType, and a list of 2D
// points that are interpreted differently based on the element's type.
type PathElement struct {
	// The type of the path element, see PathElementType for details on what
	// values this field can take.
	Type PathElementType

	// The array of points that are arguments to the path operation represented
	// by this element. Based on the elemen type, one or more of the points are
	// actually meaningful to set.
	Points [3]Point
}

// A Path is a drawable representation of an arbitrary shape, it's copmosed of a
// list of elements describing each step of the whole drawing operation.
type Path struct {
	// The list of elements in the path. A program shouldn't have to manipulate
	// this field and should be considered read-only.
	Elements []PathElement
}

// The Shape interface represents values that can be converted to 2D paths.
type Shape interface {
	// Returns a 2D path representing the abstract shape that satisifed the
	// interface.
	//
	// The returned path shouldn't be retained by the shape or any other part
	// of the program, the caller of that method is considered the owner of the
	// returned value and is free to modify it.
	Path() Path
}

// MakePath returns a Path value initialized with a slice of elements with the
// capacity given as argument.
func MakePath(cap int) Path {
	return Path{
		Elements: make([]PathElement, 0, cap),
	}
}

// AppendPath concatenates the Path given as second arugment to the one given as
// first argument. It behaves in a similar way to the builtin `append` function,
// the Path value given as first argument may or may not be modified by the append
// operation.
func AppendPath(path Path, other Path) Path {
	return Path{Elements: append(path.Elements, other.Elements...)}
}

// AppendRect efficiently append a rectangle to a Path and returns the modified
// value.
//
// Calling this function is equivalent to calling:
//
//	path = AppendPath(path, rect.Path())
//
// but the implementation is optimized to avoid unnecessary memory allocations.
func AppendRect(path Path, rect Rect) Path {
	x0 := rect.X
	y0 := rect.Y
	x1 := rect.X + rect.W
	y1 := rect.Y + rect.H
	return AppendPolygon(path, Point{x0, y0}, Point{x1, y0}, Point{x1, y1}, Point{x0, y1})
}

// AppendPolygon appends an arbitrary polygon to a path and returns the modified
// value, the polygon is made of the points given as arguments to the function.
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

// MoveTo appends a path element that moves the current path position to the
// point given as argument.
func (p *Path) MoveTo(pt Point) {
	p.append(PathElement{
		Type:   MoveTo,
		Points: [...]Point{pt, {}, {}},
	})
}

// LineTo appends a path element that draws a line from the current path
// position to the point given as argument.
func (p *Path) LineTo(pt Point) {
	p.ensureStartWithMoveTo()
	p.append(PathElement{
		Type:   LineTo,
		Points: [...]Point{pt, {}, {}},
	})
}

// QuadCurveTo appends a path element that draws a quadriatic curve from the
// current path position to `pt` and centered in `c`.
func (p *Path) QuadCurveTo(c Point, pt Point) {
	p.ensureStartWithMoveTo()
	p.append(PathElement{
		Type:   QuadCurveTo,
		Points: [...]Point{c, pt, {}},
	})
}

// CubicCurveTo appends a path element that draws a cubic curve from the current
// path position to `pt` with centers in `c1` and `c2`.
func (p *Path) CubicCurveTo(c1 Point, c2 Point, pt Point) {
	p.ensureStartWithMoveTo()
	p.append(PathElement{
		Type:   CubicCurveTo,
		Points: [...]Point{c1, c2, pt},
	})
}

// Close appends a path element that closes the current shape by drawing a line
// between the current path position and the last move-to element added to the
// path.
func (p *Path) Close() {
	p.append(PathElement{
		Type: ClosePath,
	})
}

// Clear erases every element in the path.
func (p *Path) Clear() {
	p.Elements = p.Elements[:0]
}

// Copy creates a copy of the path and returns it, the returned value and the
// receiver do not share the same slice of elements and can be safely modified
// independently.
func (p *Path) Copy() Path {
	if p.Empty() {
		return Path{}
	}

	p1 := Path{
		Elements: make([]PathElement, len(p.Elements)),
	}

	copy(p1.Elements, p.Elements)
	return p1
}

// LastPoint returns the 2D coordinates of the current path position.
func (p *Path) LastPoint() Point {
	return p.lastPointAt(len(p.Elements) - 1)
}

func (p *Path) lastPointAt(n int) Point {
	if n < 0 {
		return Point{}
	}

	switch e := p.Elements[n-1]; e.Type {
	case MoveTo, LineTo:
		return e.Points[0]

	case QuadCurveTo:
		return e.Points[1]

	case CubicCurveTo:
		return e.Points[2]

	default:
		return p.lastPointAt(n - 1)
	}
}

// Empty checks if the path is empty, which means it has no element.
func (p *Path) Empty() bool {
	return len(p.Elements) == 0
}

// Path satisfies the Shape interface by returning a copy of the path it is
// called on.
func (p Path) Path() Path {
	return p.Copy()
}

func (p *Path) append(e PathElement) {
	p.Elements = append(p.Elements, e)
}

func (p *Path) ensureStartWithMoveTo() {
	if n := len(p.Elements); n == 0 {
		// Empty paths must start with a MoveTo element.
		p.MoveTo(Point{})

	} else if p.Elements[n-1].Type == ClosePath {
		// When a subpath is closed the path that contains more elements must
		// be restarted with a MoveTo as well.
		p.MoveTo(p.LastPoint())
	}
}
