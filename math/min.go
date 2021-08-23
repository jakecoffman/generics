package math

type comparable interface {
	~int
}

func Min[T comparable](a ...T) T {
	m := a[0]
	for _, v := range a {
		if m > v {
			m = v
		}
	}
	return m
}
