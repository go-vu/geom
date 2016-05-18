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

func (self *Rect) SetOrigin(p Point) {
	self.X = p.X
	self.Y = p.Y
}

func (self *Rect) SetSize(s Size) {
	self.W = s.W
	self.H = s.H
}

func (self *Rect) Grow(k float64) {
	self.X -= k
	self.Y -= k

	k += k

	self.W += k
	self.H += k
}

func (self Rect) Origin() Point {
	return Point{
		X: self.X,
		Y: self.Y,
	}
}

func (self Rect) Center() Point {
	return Point{
		X: self.X + (self.W / 2),
		Y: self.Y + (self.H / 2),
	}
}

func (self Rect) Tip() Point {
	return Point{
		X: self.X + self.W,
		Y: self.Y + self.H,
	}
}

func (self Rect) Size() Size {
	return Size{
		W: self.W,
		H: self.H,
	}
}

func (self Rect) Abs() Rect {
	r := self

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

func (self Rect) Intersect(r Rect) Rect {
	r1 := self.Abs()
	r2 := r.Abs()

	x1 := math.Max(r1.X, r2.X)
	y1 := math.Max(r1.Y, r2.Y)

	x2 := math.Min(r1.X+r1.W, r2.X+r2.W)
	y2 := math.Min(r1.Y+r1.H, r2.Y+r2.H)

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

func (self Rect) Merge(r Rect) Rect {
	r1 := self.Abs()
	r2 := r.Abs()

	x1 := math.Min(r1.X, r2.X)
	y1 := math.Min(r1.Y, r2.Y)

	x2 := math.Max(r1.X+r1.W, r2.X+r2.W)
	y2 := math.Max(r1.Y+r1.H, r2.Y+r2.H)

	return Rect{
		X: x1,
		Y: y1,
		W: x2 - x1,
		H: y2 - y1,
	}
}

func (self Rect) Zero() bool {
	return self.Origin().Zero() && self.Size().Zero()
}

func (self Rect) Empty() bool {
	return self.Size().Empty()
}

func (self Rect) Area() float64 {
	return self.Size().Area()
}

func (self Rect) ContainsPoint(p Point) bool {
	return self.X <= p.X && (self.X+self.W) > p.X && self.Y <= p.Y && (self.Y+self.H) > p.Y
}

func (self Rect) ContainsRect(r Rect) bool {
	return self.X <= r.X && (self.X+self.W) >= (r.X+r.W) && self.Y <= r.Y && (self.Y+self.H) >= (r.Y+r.H)
}

func (self Rect) String() string {
	return fmt.Sprintf("{ %.6g, %.6g, %.6g, %.6g }", self.X, self.Y, self.W, self.H)
}

func (self Rect) Path() Path {
	return AppendRect(MakePath(5), self)
}

func MergeRect(slice []Rect, rect Rect) []Rect {
	// First we check if one of the rectangles in the slice contains the one we
	// are trying to add. If that's the case there's nothing we need to do, the
	// rectangle is already contained in the slice.
	for _, r := range slice {
		if r.ContainsRect(rect) {
			return slice
		}
	}

	// Look for rectangles that may intersect with the one we are trying to add,
	// if one is found we merge them and rebuild the slice in case it created
	// new intersections.
	for i, r := range slice {
		if !r.Intersect(rect).Empty() {
			s := append(make([]Rect, 0, len(slice)), r.Merge(rect))

			for j, r := range slice {
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
	return append(slice, rect)
}
