package clap

import (
	"fmt"
	"testing"

	"github.com/jvlmdr/lin-go/cmat"
)

// Minimum-residual solution to over-constrained system by QR decomposition.
func TestQRFact_Solve_overdetermined(t *testing.T) {
	m, n := 150, 100
	a, b, want, err := overDetProb(m, n)
	if err != nil {
		t.Fatal(err)
	}

	qr, err := QR(a)
	if err != nil {
		t.Fatal(err)
	}
	got, err := qr.Solve(false, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

// Minimum-norm solution to under-constrained system by QR decomposition.
func TestQRFact_Solve_underdetermined(t *testing.T) {
	m, n := 100, 150
	a, b, want, err := underDetProb(m, n)
	if err != nil {
		t.Fatal(err)
	}

	// Take QR factorization of conjugate transpose.
	qr, err := QR(cmat.H(a))
	if err != nil {
		t.Fatal(err)
	}
	got, err := qr.Solve(true, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func ExampleQRFact_Solve() {
	a := cmat.NewRows([][]complex128{
		{4, 2},
		{1, 1},
		{2, 0},
	})
	b_over := []complex128{6, 7, 4}
	b_under := []complex128{39, 19}

	qr, err := QR(a)
	if err != nil {
		fmt.Println(err)
		return
	}

	x_over, err := qr.Solve(false, b_over)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(formatSlice(x_over, 'f', 3))

	x_under, err := qr.Solve(true, b_under)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(formatSlice(x_under, 'f', 3))
	// Output:
	// [(1.000+0.000i) (2.000+0.000i)]
	// [(8.000+0.000i) (3.000+0.000i) (2.000+0.000i)]
}
