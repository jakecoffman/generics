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

func Map[T, R any](s *Stream[T], mapper func(*T) R) *Stream[R] {
	newStream := &Stream[R]{}
	for i := range s.arr {
		newStream.arr = append(newStream.arr, mapper(&s.arr[i]))
	}
	return newStream
}
