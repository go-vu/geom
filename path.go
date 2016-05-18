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
// but the implementation is more optimized and avoid unnecessary memory
// allocations.
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
func (path *Path) MoveTo(p Point) {
	path.append(PathElement{
		Type:   MoveTo,
		Points: [...]Point{p, {}, {}},
	})
}

// LineTo appends a path element that draws a line from the current path
// position to the point given as argument.
func (path *Path) LineTo(p Point) {
	path.ensureStartWithMoveTo()
	path.append(PathElement{
		Type:   LineTo,
		Points: [...]Point{p, {}, {}},
	})
}

// QuadCurveTo appends a path element that draws a quadriatic curve from the
// current path position to `p` and centered in `c`.
func (path *Path) QuadCurveTo(c Point, p Point) {
	path.ensureStartWithMoveTo()
	path.append(PathElement{
		Type:   QuadCurveTo,
		Points: [...]Point{c, p, {}},
	})
}

// CubicCurveTo appends a path element that draws a cubic curve from the current
// path position to `p` with centers in `c1` and `c2`.
func (path *Path) CubicCurveTo(c1 Point, c2 Point, p Point) {
	path.ensureStartWithMoveTo()
	path.append(PathElement{
		Type:   CubicCurveTo,
		Points: [...]Point{c1, c2, p},
	})
}

// Close appends a path element that closes the current shape by drawing a line
// between the current path position and the last move-to element added to the
// path.
func (path *Path) Close() {
	path.append(PathElement{
		Type: ClosePath,
	})
}

// Clear erases every element in the path.
func (path *Path) Clear() {
	path.Elements = path.Elements[:0]
}

// Copy creates a copy of the path and returns it, the returned value and the
// receiver do not share the same slice of elements and can be safely modified
// independently.
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

// LastPoint returns the 2D coordinates of the current path position.
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

// Empty checks if the path is empty, which means it has no element.
func (path *Path) Empty() bool {
	return len(path.Elements) == 0
}

// Path satisfies the Shape interface by returning a copy of the path it is
// called on.
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
