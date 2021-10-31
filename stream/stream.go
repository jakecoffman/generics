package stream

type Stream[T any] struct {
	arr []T
}

func New[T any](arr []T) *Stream[T] {
	return &Stream[T]{arr: arr}
}

func (s *Stream[T]) Filter(filter func(item *T) bool) *Stream[T] {
	newArr := []T{}
	for i := range s.arr {
		if filter(&s.arr[i]) {
			newArr = append(newArr, s.arr[i])
		}
	}
	return New(newArr)
}

func (s *Stream[T]) Some(find func(item *T) bool) bool {
	for i := range s.arr {
		if find(&s.arr[i]) {
			return true
		}
	}
	return false
}

func (s *Stream[T]) ToSlice() []T {
	return s.arr
}
