package zmat

import "github.com/jackvalmadre/lin-go/zvec"

// Describes a dense matrix with contiguous storage in column-major order.
//
// Being contiguous enables reshaping.
// Being contiguous and column-major enables column slicing and appending.
type Contiguous struct {
	Rows int
	Cols int
	// The (i, j)-th element is at Elements[j*Rows+i].
	Elements []complex128
}

// Makes a new rows x cols contiguous matrix.
func MakeContiguous(rows, cols int) Contiguous {
	return Contiguous{rows, cols, make([]complex128, rows*cols)}
}

// Copies B into a new contiguous matrix.
func MakeContiguousCopy(B Const) Contiguous {
	rows, cols := RowsCols(B)
	A := MakeContiguous(rows, cols)
	Copy(A, B)
	return A
}

func (A Contiguous) Size() Size                 { return Size{A.Rows, A.Cols} }
func (A Contiguous) At(i, j int) complex128     { return A.Elements[j*A.Rows+i] }
func (A Contiguous) Set(i, j int, x complex128) { A.Elements[j*A.Rows+i] = x }

func (A Contiguous) ColMajorArray() []complex128 { return A.Elements }
func (A Contiguous) Stride() int                 { return A.Rows }

// Transpose without copying.
func (A Contiguous) T() ContiguousRowMajor { return ContiguousRowMajor(A) }

// Returns a vectorization which accesses the array directly.
func (A Contiguous) Vec() zvec.Mutable { return zvec.Slice(A.Elements) }

// Modifies the rows and columns of a contiguous matrix.
// The number of elements must remain constant.
//
// The returned matrix references the same data.
func (A Contiguous) Reshape(s Size) Contiguous {
	if s.Area() != A.Size().Area() {
		panic("Number of elements must match to resize")
	}
	return Contiguous{s.Rows, s.Cols, A.Elements}
}

// Slices the columns.
//
// The returned matrix references the same data.
func (A Contiguous) Slice(j0, j1 int) Contiguous {
	return Contiguous{A.Rows, j1 - j0, A.Elements[j0*A.Rows : j1*A.Rows]}
}

// Appends a column.
//
// The returned matrix may reference the same data.
func (A Contiguous) AppendVector(x zvec.Const) Contiguous {
	if A.Rows != x.Len() {
		panic("Dimension of vector does not match matrix")
	}
	elements := zvec.Append(A.Elements, x)
	return Contiguous{A.Rows, A.Cols + 1, elements}
}

// Appends a matrix horizontally. The number of rows must match.
//
// The returned matrix may reference the same data.
func (A Contiguous) AppendMatrix(B Const) Contiguous {
	if A.Rows != Rows(B) {
		panic("Dimension of matrices does not match")
	}
	elements := zvec.Append(A.Elements, Vec(B))
	return Contiguous{A.Rows, A.Cols + Cols(B), elements}
}

// Appends a column-major matrix horizontally. The number of rows must match.
//
// The returned matrix may reference the same data.
func (A Contiguous) AppendContiguous(B Contiguous) Contiguous {
	if A.Rows != B.Rows {
		panic("Dimension of matrices does not match")
	}
	elements := append(A.Elements, B.Elements...)
	return Contiguous{A.Rows, A.Cols + B.Cols, elements}
}

// Selects a submatrix within the contiguous matrix.
func (A Contiguous) Submat(r Rect) ContiguousSubmat {
	// Extract from first to last element.
	a := r.Min.I + r.Min.J*A.Rows
	b := (r.Max.I - 1) + (r.Max.J-1)*A.Rows + 1
	return ContiguousSubmat{r.Rows(), r.Cols(), A.Rows, A.Elements[a:b]}
}

// Returns a mutable column as a slice vector.
func (A Contiguous) Col(j int) zvec.Slice {
	return ContiguousCol(A, j)
}

// Returns MutableRow(A, i).
func (A Contiguous) Row(i int) zvec.Mutable { return MutableRow(A, i) }
