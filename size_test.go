package geom

import "testing"

func TestSizeZeroTrue(t *testing.T) {
	s := Size{}

	if !s.Zero() {
		t.Error("Size zero value was not detected by the Zero method")
	}
}

func TestSizeZeroFalse(t *testing.T) {
	s := Size{W: 1}

	if s.Zero() {
		t.Error("Size non-zero value was not detected by the Zero method")
	}
}

func TestSizeEmptyTrue(t *testing.T) {
	s := Size{W: 1}

	if !s.Empty() {
		t.Error("Size empty value was not detected by the Empty method")
	}
}

func TestSizeEmptyFalse(t *testing.T) {
	s := Size{W: 1, H: 1}

	if s.Empty() {
		t.Error("Size non-empty value was not detected by the Empty method")
	}
}

func TestSizeArea(t *testing.T) {
	s := Size{W: 2, H: 3}
	a := s.Area()

	if a != 6 {
		t.Error("invalid value returned by the Area method:", s, a)
	}
}

func TestSizeRatio(t *testing.T) {
	s := Size{W: 2, H: 4}
	r := s.Ratio()

	if r != 0.5 {
		t.Error("invalid value returned by the Ratio method:", s, r)
	}
}

func TestSizeString(t *testing.T) {
	s := Size{W: 2, H: 3}
	a := s.String()

	if a != "[2, 3]" {
		t.Errorf("invalid string representation of Size value: %#v %s", s, a)
	}
}
