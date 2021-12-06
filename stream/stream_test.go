package stream

import (
	"reflect"
	"strconv"
	"testing"
)

func TestStream(t *testing.T) {
	arr := []int{1, 12, 3}

	t.Run("Some works", func(t *testing.T) {
		stream := New(arr)
		if !stream.Some(func(item *int) bool { return *item%2 == 0 }) {
			t.Errorf("Expected true")
		}
	})

	t.Run("Filtering", func(t *testing.T) {
		stream := New(arr)
		stream = stream.Filter(func(item *int) bool { return *item%2 != 0 })
		if stream.Some(func(item *int) bool { return *item%2 == 0 }) {
			t.Errorf("Expected false")
		}
	})

	t.Run("Mapping", func(t *testing.T) {
		stream := Map(New(arr), func(t *int) string {
			return strconv.Itoa(*t)
		})
		stream = stream.Filter(func(item *string) bool {
			return len(*item) == 1
		})
		if !reflect.DeepEqual([]string{"1", "3"}, stream.ToSlice()) {
			t.Error(stream.ToSlice())
		}
	})
}
