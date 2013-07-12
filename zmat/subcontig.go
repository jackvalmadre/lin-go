package zmat

import "github.com/jackvalmadre/go-vec/zvec"

// Submatrix within a contiguous matrix, column-major order.
type SubContiguous struct {
	Rows   int
	Cols   int
	Stride int
	// The (i, j)-th element is at Elements[j*Stride+i].
	Elements []complex128
}

func (A SubContiguous) Size() Size                 { return Size{A.Rows, A.Cols} }
func (A SubContiguous) At(i, j int) complex128     { return A.Elements[j*A.Stride+i] }
func (A SubContiguous) Set(i, j int, x complex128) { A.Elements[j*A.Stride+i] = x }

func (A SubContiguous) Submatrix(r Rect) SubContiguous {
	// Extract from first to last elements.
	i0, j0 := r.Min.I, r.Min.J
	i1, j1 := r.Max.I, r.Max.J
	a := j0*A.Stride + i0
	b := (j1-1)*A.Stride + (i1 - 1) + 1
	return SubContiguous{r.Rows(), r.Cols(), A.Stride, A.Elements[a:b]}
}

// Returns a wrapper for accessing elements as a vector.
func (A SubContiguous) Vec() zvec.Mutable { return MutableVec(A) }

// Returns MutableT(A).
func (A SubContiguous) T() Mutable { return MutableT(A) }

// Returns MutableColumn(A).
func (A SubContiguous) Col(j int) zvec.Mutable { return MutableCol(A, j) }

// Returns MutableRow(A).
func (A SubContiguous) Row(i int) zvec.Mutable { return MutableRow(A, i) }