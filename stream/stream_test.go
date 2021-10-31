package stream

import "testing"

func TestStream_Some(t *testing.T) {
	arr := []int{1, 2, 3}
	if !New(arr).
		Some(func(item *int) bool { return *item%2 == 0 }) {
		t.Errorf("Expected true")
	}

	if New(arr).
		Filter(func(item *int) bool { return *item%2 != 0 }).
		Some(func(item *int) bool { return *item%2 == 0 }) {
		t.Errorf("Expected false")
	}
}
