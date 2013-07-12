package zmat

// Describes a read-only matrix.
type Const interface {
	Size() Size
	At(i, j int) complex128
}

// Describes a fixed-size matrix whose elements can be modified.
type Mutable interface {
	Const
	Set(i, j int, x complex128)
}
