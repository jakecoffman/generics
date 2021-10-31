package zip

type Zip[T any, U any] struct {
	a      []T
	b      []U
	cursor int
}

func (z *Zip[T, U]) Next() bool {
	return z.cursor < len(z.a) && z.cursor < len(z.b)
}

func (z *Zip[T, U]) Value() (T, U) {
	a, b := z.a[z.cursor], z.b[z.cursor]
	z.cursor++
	return a, b
}

func New[T any, U any](arr1 []T, arr2 []U) *Zip[T, U] {
	return &Zip[T, U]{a: arr1, b: arr2}
}
