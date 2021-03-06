package lapack

import (
	"fmt"
	"testing"

	"github.com/jvlmdr/lin-go/mat"
)

func TestCholFact_Solve(t *testing.T) {
	n := 100
	// Random symmetric positive definite matrix.
	a := randMat(2*n, n)
	a = mat.Mul(mat.T(a), a)
	// Random vector.
	want := randVec(n)
	b := mat.MulVec(a, want)

	// Factorize matrix.
	chol, err := Chol(a)
	if err != nil {
		t.Fatal(err)
	}

	got, err := chol.Solve(b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func ExampleCholFact_Solve() {
	// A = V' V, with V = [1, 1; 2, 1]
	v := mat.NewRows([][]float64{
		{1, 1},
		{2, 1},
	})
	a := mat.Mul(mat.T(v), v)

	// x = [1; 2]
	// b = V' V x
	//   = V' [1, 1; 2, 1] [1; 2]
	//   = [1, 2; 1, 1] [3; 4]
	//   = [11; 7]
	b := []float64{11, 7}

	chol, err := Chol(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	x, err := chol.Solve(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.6g", x)
	// Output:
	// [1 2]
}

func ExampleInvertPosDef() {
	// A = V' V, with V = [1, 1; 2, 1]
	v := mat.NewRows([][]float64{
		{1, 1},
		{2, 1},
	})
	a := mat.Mul(mat.T(v), v)

	b, err := InvertPosDef(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	for j := 0; j < 2; j++ {
		for i := 0; i < 2; i++ {
			if i > 0 {
				fmt.Printf(" ")
			}
			fmt.Printf("%.3g", b.At(i, j))
		}
		fmt.Println()
	}
	// Output:
	// 2 -3
	// -3 5
}
