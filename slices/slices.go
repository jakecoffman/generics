package slices

import "sync"

// Filter removes elements to which f is false.
func Filter[T any](arr []T, f func(i int) bool) []T {
	var result []T
	for i := range arr {
		if f(i) {
			result = append(result, arr[i])
		}
	}
	return result
}

// Map applies f to every element of arr.
func Map[T, U any](arr []T, f func(i int) U) []U {
	result := make([]U, len(arr))
	for i := range arr {
		result[i] = f(i)
	}
	return result
}

// MapWithPool is Map but runs with a goroutine pool of workers.
func MapWithPool[T, U any](arr []T, poolSize int, f func(i int) U) []U {
	inputs := make(chan int)
	result := make([]U, len(arr))

	wg := sync.WaitGroup{}
	wg.Add(poolSize)
	for w := 1; w <= poolSize; w++ {
		go func() {
			for i := range inputs {
				result[i] = f(i)
			}
			wg.Done()
		}()
	}

	for i := range arr {
		inputs <- i
	}
	close(inputs)
	wg.Wait()

	return result
}

// IndexOf returns the index for which f is true.
func IndexOf[T any](arr []T, f func(i int) bool) int {
	for i := range arr {
		if f(i) {
			return i
		}
	}
	return -1
}

// Find returns a pointer to the element where f is true.
func Find[T any](arr []T, f func(i int) bool) *T {
	for i := range arr {
		if f(i) {
			return &arr[i]
		}
	}
	return nil
}

// Some returns true if f is true for any element in the array.
func Some[T any](arr []T, f func(i int) bool) bool {
	for i := range arr {
		if f(i) {
			return true
		}
	}
	return false
}

// Reduce accumulates a value from the array from left to right and returns the result.
func Reduce[T any](arr []T, f func(a, b T) T) T {
	if len(arr) == 0 {
		var zero T
		return zero
	}
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = f(result, arr[i])
	}
	return result
}

// Unique returns a new array with duplicates removed.
func Unique[T comparable](arr []T) []T {
	result := []T{}
	m := map[T]struct{}{}
	for i := range arr {
		if _, ok := m[arr[i]]; !ok {
			result = append(result, arr[i])
			m[arr[i]] = struct{}{}
		}
	}
	return result
}

// ForEach calls f on each element of arr.
func ForEach[T any](arr []T, f func(e *T)) {
	for i := range arr {
		f(&arr[i])
	}
}

// Pop removes the last item from the array and returns the pointer to it.
// An array pointer must be passed so the array length can be modified.
// A pointer is returned to avoid expensive copies.
func Pop[T any](arr *[]T) *T {
	length := len(*arr)
	if length == 0 {
		return nil
	}
	r := &(*arr)[length-1]
	*arr = (*arr)[:length-1]
	return r
}

// Push adds elements to the end of arr and returns the new length of the array.
func Push[T any](arr *[]T, elements ...T) int {
	*arr = append(*arr, elements...)
	return len(*arr)
}

// Unshift adds elements to the front of arr and returns the new length of the array.
func Unshift[T any](arr *[]T, elements ...T) int {
	*arr = append(elements, *arr...)
	return len(*arr)
}

// Shift removes an element from the front of arr and returns the new length of the array.
func Shift[T any](arr *[]T) int {
	if len(*arr) == 0 {
		return 0
	}
	*arr = (*arr)[1:]
	return len(*arr)
}

// Splice changes the contents of an array by removing or replacing existing
// elements and/or adding new elements in place.
func Splice[T any](arr *[]T, start, deleteCount int, items ...T) []T {
	rv := (*arr)[start : start+deleteCount]
	*arr = append((*arr)[:start], append(items, (*arr)[start+deleteCount:]...)...)
	return rv
}

// Reverse reverses the array. Unlike JavaScript this does not mutate the array passed.
func Reverse[T any](arr []T) []T {
	result := make([]T, len(arr))
	copy(result, arr)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
