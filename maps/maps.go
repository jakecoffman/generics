package maps

// Keys returns the keys of the map m.
func Keys[T comparable, U any](m map[T]U) []T {
	keys := make([]T, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// Values returns the values of the map m.
func Values[T comparable, U any](m map[T]U) []U {
	var keys []U
	for _, value := range m {
		keys = append(keys, value)
	}
	return keys
}

// Assign assigns all values from the source maps to the target map, and returns the target.
// In other words, it smooshes all the sources into the target.
func Assign[T comparable, U any](target map[T]U, source ...map[T]U) map[T]U {
	for i := range source {
		for k := range source[i] {
			target[k] = source[i][k]
		}
	}
	return target
}
