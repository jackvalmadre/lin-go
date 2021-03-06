package cmat

import "testing"

func TestCopy(t *testing.T) {
	a := NewRows([][]complex128{
		{1, 2, 3},
		{4, 5, 6},
	})
	got := NewRows([][]complex128{
		{-1, 0, 1},
		{-1, 2, -1},
	})
	Copy(got, a)
	want := NewRows([][]complex128{
		{1, 2, 3},
		{4, 5, 6},
	})
	testMatEq(t, want, got)
}

func TestSub(t *testing.T) {
	a := NewRows([][]complex128{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	r := Rect{Pos{1, 0}, Pos{3, 2}}
	got := Sub(a, r)
	want := NewRows([][]complex128{
		{4, 5},
		{7, 8},
	})
	testMatEq(t, want, got)
}

func TestCopy_ref(t *testing.T) {
	got := NewRows([][]complex128{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	r := Rect{Pos{1, 0}, Pos{3, 2}}
	a := NewRows([][]complex128{
		{-1, -2},
		{-3, -4},
	})
	Copy(Ref{got, r}, a)
	want := NewRows([][]complex128{
		{1, 2, 3},
		{-1, -2, 6},
		{-3, -4, 9},
	})
	testMatEq(t, want, got)
}
