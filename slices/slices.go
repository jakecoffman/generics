package slices

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
