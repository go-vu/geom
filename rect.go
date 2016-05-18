package geom

import (
	"fmt"
	"math"
)

type Rect struct {
	X float64
	Y float64
	W float64
	H float64
}

func MakeRect(p Point, s Size) Rect {
	return Rect{
		X: p.X,
		Y: p.Y,
		W: s.W,
		H: s.H,
	}
}

func (r *Rect) SetOrigin(p Point) {
	r.X = p.X
	r.Y = p.Y
}

func (r *Rect) SetSize(s Size) {
	r.W = s.W
	r.H = s.H
}

func (r *Rect) Grow(k float64) {
	r.X -= k
	r.Y -= k

	k += k

	r.W += k
	r.H += k
}

func (r Rect) Origin() Point {
	return Point{
		X: r.X,
		Y: r.Y,
	}
}

func (r Rect) Center() Point {
	return Point{
		X: r.X + (r.W / 2),
		Y: r.Y + (r.H / 2),
	}
}

func (r Rect) Tip() Point {
	return Point{
		X: r.X + r.W,
		Y: r.Y + r.H,
	}
}

func (r Rect) Size() Size {
	return Size{
		W: r.W,
		H: r.H,
	}
}

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

func (r Rect) Zero() bool {
	return r.Origin().Zero() && r.Size().Zero()
}

func (r Rect) Empty() bool {
	return r.Size().Empty()
}

func (r Rect) Area() float64 {
	return r.Size().Area()
}

func (r Rect) ContainsPoint(p Point) bool {
	return r.X <= p.X && (r.X+r.W) > p.X && r.Y <= p.Y && (r.Y+r.H) > p.Y
}

func (r0 Rect) ContainsRect(r1 Rect) bool {
	return r0.X <= r1.X && (r0.X+r0.W) >= (r1.X+r1.W) && r0.Y <= r1.Y && (r0.Y+r0.H) >= (r1.Y+r1.H)
}

func (r Rect) String() string {
	return fmt.Sprintf("{ %.6g, %.6g, %.6g, %.6g }", r.X, r.Y, r.W, r.H)
}

func (r Rect) Path() Path {
	return AppendRect(MakePath(5), r)
}

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
