package slices

func Filter[T any](arr []T, f func(i int) bool) []T {
	var result []T
	for i := range arr {
		if f(i) {
			result = append(result, arr[i])
		}
	}
	return result
}

func Map[T, U any](arr []T, f func(i int) U) []U {
	result := make([]U, len(arr))
	for i := range arr {
		result[i] = f(i)
	}
	return result
}

func IndexOf[T any](arr []T, f func(i int) bool) int {
	for i := range arr {
		if f(i) {
			return i
		}
	}
	return -1
}

func Find[T any](arr []T, f func(i int) bool) *T {
	for i := range arr {
		if f(i) {
			return &arr[i]
		}
	}
	return nil
}
