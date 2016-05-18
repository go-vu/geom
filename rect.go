package geom

import (
	"fmt"
	"math"
)

// The Rect type represents a 2D rectangle made of the coordinates of an origin
// point, and width and height dimensions.
type Rect struct {
	X float64
	Y float64
	W float64
	H float64
}

// MakeRect constructs a Rect value from an origin and dimensions provided as a
// Point and Size value..
func MakeRect(p Point, s Size) Rect {
	return Rect{
		X: p.X,
		Y: p.Y,
		W: s.W,
		H: s.H,
	}
}

// SetOrigin changes the origin of the rectangle instance it is called on to the
// Point passed as argument.
func (r *Rect) SetOrigin(p Point) {
	r.X = p.X
	r.Y = p.Y
}

// SetSize changes the dimensions of the rectangle instance it is called on to
// the Size passed as argument.
func (r *Rect) SetSize(s Size) {
	r.W = s.W
	r.H = s.H
}

// The Grow method modifies the origin and dimensions of the rectangle it is
// called on by the value passed as argument.
func (r *Rect) Grow(k float64) {
	r.X -= k
	r.Y -= k

	k += k

	r.W += k
	r.H += k
}

// Origin returns the current origin of the rectangle as a Point value.
func (r Rect) Origin() Point {
	return Point{
		X: r.X,
		Y: r.Y,
	}
}

// Center returns the current center of the rectangle as a Point value.
func (r Rect) Center() Point {
	return Point{
		X: r.X + (r.W / 2),
		Y: r.Y + (r.H / 2),
	}
}

// Tip returns the 'bottom-right' point of the rectangle as a Point value.
func (r Rect) Tip() Point {
	return Point{
		X: r.X + r.W,
		Y: r.Y + r.H,
	}
}

// Size returns the dimensions of the rectangle as a Size value.
func (r Rect) Size() Size {
	return Size{
		W: r.W,
		H: r.H,
	}
}

// The Abs method transforms returns a rectangle equivalent to the one it is
// called on where negative components of the dimensions have been converted
// to positive values.
func (r Rect) Abs() Rect {
	if r.W < 0 {
		r.X += r.W
		r.W = -r.W
	}

	if r.H < 0 {
		r.Y += r.H
		r.H = -r.H
	}

	return r
}

// The Intersect method computes the intersection of the rectangle it's called
// on with the one passed as argument, returning the result.
//
// The intersection is the area that is covered by both rectangles, which may be
// a zero-value if the two don't overlap.
func (r0 Rect) Intersect(r1 Rect) Rect {
	r0 = r0.Abs()
	r1 = r1.Abs()

	x1 := math.Max(r0.X, r1.X)
	y1 := math.Max(r0.Y, r1.Y)

	x2 := math.Min(r0.X+r0.W, r1.X+r1.W)
	y2 := math.Min(r0.Y+r0.H, r1.Y+r1.H)

	if x1 > x2 || y1 > y2 {
		return Rect{}
	}

	return Rect{
		X: x1,
		Y: y1,
		W: x2 - x1,
		H: y2 - y1,
	}
}

// The Merge method computes and returns the smallest rectangle that includes
// both the one it is called on and the one passed as argument.
func (r0 Rect) Merge(r1 Rect) Rect {
	r0 = r0.Abs()
	r1 = r1.Abs()

	x1 := math.Min(r0.X, r1.X)
	y1 := math.Min(r0.Y, r1.Y)

	x2 := math.Max(r0.X+r0.W, r1.X+r1.W)
	y2 := math.Max(r0.Y+r0.H, r1.Y+r1.H)

	return Rect{
		X: x1,
		Y: y1,
		W: x2 - x1,
		H: y2 - y1,
	}
}

// The Zero method checks whether the rectangle it is called on is the
// zero-value.
func (r Rect) Zero() bool {
	return r.Origin().Zero() && r.Size().Zero()
}

// The Empty method checks whether the rectangle is empty, which means it has a
// zero area.
func (r Rect) Empty() bool {
	return r.Size().Empty()
}

// The Area method computes and returns the area of the rectangle it is called on
// (width x height).
func (r Rect) Area() float64 {
	return r.Size().Area()
}

// ContainsPoint checks whether the point passed as argument is contained in the
// rectangle it's called on, returning true when that's the case, false otherwise.
func (r Rect) ContainsPoint(p Point) bool {
	return r.X <= p.X && (r.X+r.W) > p.X && r.Y <= p.Y && (r.Y+r.H) > p.Y
}

// ContainsRect checks whether the rectangle passed as argument is contained in
// the one it's called on, returning true when that's the case, false otherwise.
func (r0 Rect) ContainsRect(r1 Rect) bool {
	return r0.X <= r1.X && (r0.X+r0.W) >= (r1.X+r1.W) && r0.Y <= r1.Y && (r0.Y+r0.H) >= (r1.Y+r1.H)
}

// The String method returns a human-readable representation of the rectangle.
func (r Rect) String() string {
	return fmt.Sprintf("{ %.6g, %.6g, %.6g, %.6g }", r.X, r.Y, r.W, r.H)
}

// The Path method satisfies the Shape interface, allowing Rect values to be
// used with programs that manipulate shapes.
func (r Rect) Path() Path {
	return AppendRect(MakePath(5), r)
}

// MergeRect merges a rectangle into an existing list of other rectangles,
// returing the potentially modified slice.
//
// Merging doesn't always append the given rectangle to the list, it won't be
// added if another rectangel in the list already contains the one passed as
// argument for example. A program should consider the list of rectangles passed
// to and returned by the function as the optimal list of rectangle that would
// contain all the merged rectangles. This is particulary useful when keeping
// track of 'dirty' areas in a way that would minimize the 'dirty' surface.
func MergeRect(list []Rect, rect Rect) []Rect {
	// First we check if one of the rectangles in the slice contains the one we
	// are trying to add. If that's the case there's nothing we need to do, the
	// rectangle is already contained in the slice.
	for _, r := range list {
		if r.ContainsRect(rect) {
			return list
		}
	}

	// Look for rectangles that may intersect with the one we are trying to add,
	// if one is found we merge them and rebuild the slice in case it created
	// new intersections.
	for i, r := range list {
		if !r.Intersect(rect).Empty() {
			s := append(make([]Rect, 0, len(list)), r.Merge(rect))

			for j, r := range list {
				if i != j {
					s = MergeRect(s, r)
				}
			}

			return s
		}
	}

	// TOOD: more clever merging policies?
	// -----------------------------------
	//
	// - don't merge if the resulting area is greater than keeping rectangles
	// - look for the merge that creates the smallest area
	// - optimize memory allocation, only recreate the list if needed
	// - sort the list by areas in descending order to improve the fast path
	//
	// While these merges could greatly reduce the areas being redrawn they
	// could also be expansive to compute for potentially little gain due to
	// the way widgets are laid out (parents contain children).
	// It is also possible that generating more redrawing calls may be slower
	// to render than asking the window system to redraw a bigger area once.

	// If none of the merging policies have matched we simply append it to the
	// slice.
	return append(list, rect)
}
