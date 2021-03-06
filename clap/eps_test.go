package clap

import (
	"fmt"
	"testing"

	"github.com/jvlmdr/lin-go/cmat"
)

func TestSolve_overdetermined(t *testing.T) {
	m, n := 150, 100
	a, b, want, err := overDetProb(m, n)
	if err != nil {
		t.Fatal(err)
	}

	got, err := Solve(a, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func TestSolve_underdetermined(t *testing.T) {
	m, n := 100, 150
	a, b, want, err := underDetProb(m, n)
	if err != nil {
		t.Fatal(err)
	}

	got, err := Solve(a, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func ExampleSolve_overdetermined() {
	// Find minimum-error solution to
	//	x     = 3,
	//	    y = 6,
	//	x + y = 3.
	a := cmat.NewRows([][]complex128{
		{1, 0},
		{0, 1},
		{1, 1},
	})
	b := []complex128{3, 6, 3}

	x, err := Solve(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(formatSlice(x, 'f', 3))
	// Output:
	// [(1.000+0.000i) (4.000+0.000i)]
}

func ExampleSolve_underdetermined() {
	// Find minimum-norm solution to
	//	x     + z = 6,
	//	    y + z = 9.
	a := cmat.NewRows([][]complex128{
		{1, 0, 1},
		{0, 1, 1},
	})
	b := []complex128{6, 9}

	x, err := Solve(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(formatSlice(x, 'f', 3))
	// Output:
	// [(1.000+0.000i) (4.000+0.000i) (5.000+0.000i)]
}
