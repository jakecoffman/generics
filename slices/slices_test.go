package slices

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	arr = Filter(arr, func(i int) bool {
		return arr[i]%2 == 0
	})
	if !reflect.DeepEqual(arr, []int{2, 4}) {
		t.Error(arr)
	}
}

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	arr2 := Map(arr, func(i int) bool {
		return arr[i]%2 == 0
	})
	if !reflect.DeepEqual(arr2, []bool{false, true, false, true}) {
		t.Error(arr)
	}
}

func TestSome(t *testing.T) {
	arr := []int{1, 3}
	if v := Some(arr, func(i int) bool {
		return arr[i]%2 == 0
	}); v {
		t.Error()
	}
	if v := Some(arr, func(i int) bool {
		return arr[i]%2 != 0
	}); !v {
		t.Error()
	}
}
