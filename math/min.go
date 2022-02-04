package math

type comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float64 | ~float32
}

func Min[T comparable](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T comparable](a, b T) T {
	if a > b {
		return a
	}
	return b
}
