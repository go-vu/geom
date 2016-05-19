package geom

import "testing"

func TestPointZeroTrue(t *testing.T) {
	p := Point{}

	if !p.Zero() {
		t.Error("Point zero value was not detected by the Zero method")
	}
}

func TestPointZeroFalse(t *testing.T) {
	p := Point{X: 1}

	if p.Zero() {
		t.Error("Point non-zero value was not detected by the Zero method")
	}
}

func TestPointString(t *testing.T) {
	p := Point{X: 1, Y: 2}
	s := p.String()

	if s != "(1, 2)" {
		t.Errorf("invalid string representation of a point: %#v %s", p, s)
	}
}

func TestPointWithOrigin(t *testing.T) {
	p := Point{X: 3, Y: 2}
	o := Point{X: 1, Y: -1}
	q := p.WithOrigin(o)

	if q != (Point{X: 2, Y: 3}) {
		t.Error("invalid point coordinates when changing origin:", p, o, q)
	}
}
