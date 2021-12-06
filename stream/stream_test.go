package stream

import (
	"reflect"
	"strconv"
	"testing"
)

func TestStream_Some(t *testing.T) {
	arr := []int{1, 12, 3}
	stream := New(arr)
	if !stream.Some(func(item *int) bool { return *item%2 == 0 }) {
		t.Errorf("Expected true")
	}

	stream = New(arr)
	stream = stream.Filter(func(item *int) bool { return *item%2 != 0 })
	if stream.Some(func(item *int) bool { return *item%2 == 0 }) {
		t.Errorf("Expected false")
	}

	stream2 := Map(New(arr), func(t *int) string {
		return strconv.Itoa(*t)
	})
	stream2 = stream2.Filter(func(item *string) bool {
		return len(*item) == 1
	})
	if !reflect.DeepEqual([]string{"1", "3"}, stream2.ToSlice()) {
		t.Error(stream2.ToSlice())
	}
}
