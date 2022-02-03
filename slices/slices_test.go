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

func TestReduce(t *testing.T) {
	var v []int
	add := func(a, b int) int {
		return a + b
	}
	if r := Reduce(v, add); r != 0 {
		t.Error(r)
	}
	v = []int{1}
	if r := Reduce(v, add); r != 1 {
		t.Error(r)
	}
	v = []int{1, 9, 8, 2, 7, 3, 6, 4, 5, 5}
	if r := Reduce(v, add); r != 50 {
		t.Error(r)
	}
}

func TestUnique(t *testing.T) {
	v := []int{1, 1, 2, 3, 3, 3, 3, 4}
	if r := Unique(v); !reflect.DeepEqual(r, []int{1, 2, 3, 4}) {
		t.Error(r)
	}
	v = nil
	if r := Unique(v); !reflect.DeepEqual(r, []int{}) {
		t.Error(r)
	}
}
